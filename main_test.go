package main

import (
	"kcommit/src"
	testresources "kcommit/test_resources"
	"testing"
	"time"
)

// --- Test models ---

func TestHistoryModel(t *testing.T) {
	dto := src.HistoryDTO{
		Projects: []src.ProjectModel{

			{
				Name: "Project 1",
				Branches: []src.BranchModel{
					{
						Name:  "proj_1_branch_1",
						Scope: "11",
					},
					{
						Name:  "proj_1_branch_2",
						Scope: "12",
					},
				},
			},

			{
				Name: "Project 2",
				Branches: []src.BranchModel{
					{
						Name:  "proj_2_branch_1",
						Scope: "21",
					},
					{
						Name:  "proj_2_branch_2",
						Scope: "22",
					},
				},
			},
		},
	}

	historyModel := src.CreateHistoryModelFromDTO(&dto)

	if len(historyModel.Projects) != 2 {
		t.Errorf("expected 2 projects, got %d", len(historyModel.Projects))
	}

	tests := []struct {
		projectName   string
		branchName    string
		expectedScope string
	}{
		{"Project 1", "proj_1_branch_1", "11"},
		{"Project 1", "proj_1_branch_2", "12"},
		{"Project 2", "proj_2_branch_1", "21"},
		{"Project 2", "proj_2_branch_2", "22"},
	}

	for _, test := range tests {
		branchData, err := historyModel.FindBranchData(test.projectName, test.branchName)
		if err != nil {
			t.Errorf("error finding branch %s in project %s: %v", test.branchName, test.projectName, err)
			continue
		}

		if branchData.Scope != test.expectedScope {
			t.Errorf("unexpected scope for branch %s in project %s: expected %s, got %s",
				test.branchName, test.projectName, test.expectedScope, branchData.Scope)
		}
	}

	newProjectName := "Project 3"
	newBranch1Name := "proj_3_branch_1"
	err := historyModel.AddBranch(newProjectName, newBranch1Name)
	if err != nil {
		t.Errorf("unexpected error when adding branch '%s' to project '%s': %v", newBranch1Name, newProjectName, err)
	}

	newBranch2Name := "proj_3_branch_2"
	err = historyModel.AddBranch(newProjectName, newBranch2Name)
	if err != nil {
		t.Errorf("unexpected error when adding branch '%s' to project '%s': %v", newBranch2Name, newProjectName, err)
	}

	if len(historyModel.Projects) != 3 {
		t.Errorf("expected 3 projects after adding new project, got %d", len(historyModel.Projects))
	}

	historyModel.SetBranch(newProjectName, newBranch1Name, "42")
	historyModel.SetBranch(newProjectName, newBranch2Name, "99")

	tests = []struct {
		projectName   string
		branchName    string
		expectedScope string
	}{
		{"Project 1", "proj_1_branch_1", "11"},
		{"Project 1", "proj_1_branch_2", "12"},
		{"Project 2", "proj_2_branch_1", "21"},
		{"Project 2", "proj_2_branch_2", "22"},
		{newProjectName, newBranch1Name, "42"},
		{newProjectName, newBranch2Name, "99"},
	}

	for _, test := range tests {
		if !historyModel.HasBranch(test.projectName, test.branchName) {
			t.Errorf("expected branch '%s' in project '%s' to exist, but it does not", test.branchName, test.projectName)
			continue
		}

		branchData, err := historyModel.FindBranchData(test.projectName, test.branchName)
		if err != nil {
			t.Errorf("error finding branch '%s' in project '%s': %v", test.branchName, test.projectName, err)
			continue
		}

		if branchData.Scope != test.expectedScope {
			t.Errorf("unexpected scope for branch '%s' in project '%s': expected '%s', got '%s'",
				test.branchName, test.projectName, test.expectedScope, branchData.Scope)
		}
	}
}

func TestCleanOldBranches(t *testing.T) {
	referenceTime := time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)
	history := setupHistoryModel(referenceTime)

	history.CleanOldBranches(referenceTime)

	assertBranchExists(t, history, "ProjectA", "Branch2", true)
	assertBranchExists(t, history, "ProjectB", "Branch2", true)
	assertBranchExists(t, history, "ProjectC", "Branch1", true)
	assertBranchExists(t, history, "ProjectD", "Branch1", false)
	assertBranchExists(t, history, "ProjectA", "Branch1", false)
}

func TestCleanOldBranchesYear(t *testing.T) {
	referenceTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	history := setupHistoryModel(referenceTime)

	history.CleanOldBranches(referenceTime)

	assertBranchExists(t, history, "ProjectA", "Branch2", true)
	assertBranchExists(t, history, "ProjectB", "Branch2", true)
	assertBranchExists(t, history, "ProjectC", "Branch1", true)
	assertBranchExists(t, history, "ProjectD", "Branch1", false)
	assertBranchExists(t, history, "ProjectA", "Branch1", false)
}

// --- Test mocks ---

func TestFileManagerMock(t *testing.T) {
	mock := testresources.FileManagerMock{
		CheckIfPathExistsReturns: map[string]interface{}{
			"valid/path": true,
			"false/path": false,
		},
		ReadFileContentReturns: map[string]interface{}{
			"path": "content",
		},
		GetHistoryContentReturns: "history",
		BasicSetupReturnValue:    nil,
	}

	_, err := mock.CheckIfPathExists("valid/path")
	if err != nil {
		t.Errorf("TestFileManagerMock CheckIfPathExists failed")
		return
	}

	_, err = mock.CheckIfPathExists("false/path")
	if err != nil {
		t.Errorf("TestFileManagerMock CheckIfPathExists failed")
		return
	}

	_, err = mock.CheckIfPathExists("invalid/path")
	if err == nil {
		t.Errorf("TestFileManagerMock CheckIfPathExists failed")
		return
	}

	calledWith := []string{"valid/path", "false/path", "invalid/path"}
	if !containsSame(calledWith, mock.CheckIfPathExistsCalledWith) {
		t.Errorf("TestFileManagerMock CheckIfPathExistsCalledWith failed")
		return
	}

	if !(mock.CheckIfPathExistsCalled == 3) {
		t.Errorf("TestFileManagerMock CheckIfPathExistsCalled failed")
		return
	}

	content, err := mock.ReadFileContent("path")
	if err != nil {
		t.Errorf("TestFileManagerMock ReadFileContent failed")
		return
	}

	if content != "content" {
		t.Errorf("TestFileManagerMock ReadFileContent failed")
		return
	}

	if !(mock.ReadFileContentCalled == 1) {
		t.Errorf("TestFileManagerMock ReadFileContent failed")
		return
	}

	content, err = mock.GetHistoryContent()
	if err != nil {
		t.Errorf("TestFileManagerMock ReadFileContent failed")
		return
	}

	if content != "history" {
		t.Errorf("TestFileManagerMock GetHistoryContent failed")
		return
	}

	mock.WriteHistoryContent("123")

	if mock.WriteHistoryContentWrittenContent != "123" {
		t.Errorf("TestFileManagerMock WriteHistoryContent failed")
		return
	}

	mock.BasicSetup()

	if !(mock.BasicSetupCalled == 1) {
		t.Errorf("TestFileManagerMock BasicSetup failed")
		return
	}
}

func TestGitMock(t *testing.T) {
	mock := testresources.GitMock{
		GetCurrentBranchReturnValue: "main",
		GitCommitReturnValue:        "commit",
		IsGitRepositoryReturnValue:  true,
	}

	branch, err := mock.GetCurrentBranch()
	if err != nil {
		t.Errorf("TestGitMock GetCurrentBranchReturnValue failed")
		return
	}

	if branch != "main" {
		t.Errorf("TestGitMock GetCurrentBranchReturnValue failed")
		return
	}

	if !(mock.GetCurrentBranchCalled == 1) {
		t.Errorf("TestGitMock GetCurrentBranchReturnValue failed")
		return
	}

	commit, err := mock.GitCommit("commit")
	if err != nil {
		t.Errorf("TestGitMock GitCommit failed")
		return
	}

	if commit != "commit" {
		t.Errorf("TestGitMock GitCommit failed")
		return
	}

	if mock.GitCommitReturnValue != "commit" {
		t.Errorf("TestGitMock GitCommit failed")
		return
	}

	if !(mock.GitCommitCalled == 1) {
		t.Errorf("TestGitMock GitCommit failed")
		return
	}

	if !mock.IsGitRepository() {
		t.Errorf("TestGitMock IsGitRepository failed")
		return
	}
}

func TestViewBuilderMock(t *testing.T) {
	mock := testresources.ViewBuilderMock{
		NewListViewReturnValue:      "newList",
		NewTextFieldViewReturnValue: "newTextField",
	}

	l := []src.ListItem{}

	resp := mock.NewListView("", l, 0)

	if resp != "newList" {
		t.Errorf("ViewBuilderMock NewListView failed")
		return
	}

	if mock.NewListViewCalled != 1 {
		t.Errorf("ViewBuilderMock NewListView failed")
		return
	}

	resp = mock.NewTextFieldView("", "")

	if resp != "newTextField" {
		t.Errorf("ViewBuilderMock NewTextFieldView failed")
		return
	}

	if mock.NewTextFieldViewCalled != 1 {
		t.Errorf("ViewBuilderMock NewTextFieldView failed")
		return
	}
}

func TestRunnerHappyPath(t *testing.T) {
	fileManager := testresources.FileManagerMock{}

	utils := testresources.UtilsMock{}

	git := testresources.GitMock{
		IsGitRepositoryReturnValue: true,
	}

	viewBuilder := testresources.ViewBuilderMock{}

	r := src.NewRunner(&fileManager, &git, &utils, &viewBuilder)
	r.Start()

	if !(git.IsGitRepositoryCalled == 1) {
		t.Errorf("Runner TestRunnerHappyPath failed")
		return
	}
}

// --- helpers ---

func containsSame(list1, list2 []string) bool {
	if len(list1) != len(list2) {
		return false
	}

	c := make(map[string]int)
	for _, e := range list1 {
		c[e]++
	}

	for _, e := range list2 {
		if c[e] == 0 {
			return false
		}
		c[e]--
	}

	return true
}

func setupHistoryModel(referenceTime time.Time) src.HistoryModel {
	lessThanOneMonthAgo := referenceTime.AddDate(0, 0, -20)
	oneMonthAgo := referenceTime.AddDate(0, -1, 0)
	oneMonthAndOneDayAgo := referenceTime.AddDate(0, -1, -1)
	twoMonthsAgo := referenceTime.AddDate(0, -2, 0)

	return src.HistoryModel{
		Projects: map[string]map[string]src.BranchDetail{
			"ProjectA": {
				"Branch1": {Scope: "feature", UpdatedAt: twoMonthsAgo},
				"Branch2": {Scope: "bugfix", UpdatedAt: lessThanOneMonthAgo},
			},
			"ProjectB": {
				"Branch1": {Scope: "hotfix", UpdatedAt: oneMonthAgo},
				"Branch2": {Scope: "ui", UpdatedAt: lessThanOneMonthAgo},
			},
			"ProjectC": {
				"Branch1": {Scope: "docs", UpdatedAt: lessThanOneMonthAgo},
			},
			"ProjectD": {
				"Branch1": {Scope: "tests", UpdatedAt: oneMonthAndOneDayAgo},
			},
		},
	}
}

func assertBranchExists(t *testing.T, history src.HistoryModel, project, branch string, shouldExist bool) {
	_, err := history.FindBranchData(project, branch)
	if shouldExist && err != nil {
		t.Errorf("Branch %s/%s should exist.", project, branch)
	} else if !shouldExist && err == nil {
		t.Errorf("Branch %s/%s should not exist.", project, branch)
	}
}