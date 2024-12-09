package main

import (
	"kcommit/src"
	"testing"
	"time"
)

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
	referenceTime := time.Date(2024, 12, 1, 12, 0, 0, 0, time.UTC)
	lessThanOneMonthAgo := referenceTime.AddDate(0, 0, -20)
	oneMonthAgo := referenceTime.AddDate(0, -1, 0)
	twoMonthsAgo := referenceTime.AddDate(0, -2, 0)

	history := src.HistoryModel{
		Projects: map[string]map[string]src.BranchDetail{
			"ProjectA": {
				"Branch1": {Scope: "feature", UpdatedAt: twoMonthsAgo},
				"Branch2": {Scope: "bugfix", UpdatedAt: referenceTime},
			},
			"ProjectB": {
				"Branch1": {Scope: "hotfix", UpdatedAt: oneMonthAgo},
				"Branch2": {Scope: "ui", UpdatedAt: lessThanOneMonthAgo},
			},
			"ProjectC": {
				"Branch1": {Scope: "docs", UpdatedAt: lessThanOneMonthAgo},
			},
		},
	}

	history.CleanOldBranches()

	_, err := history.FindBranchData("ProjectA", "Branch2")
	if err != nil {
		t.Errorf("FindBranchData() returned an unexpected error")
		return
	}

	_, err = history.FindBranchData("ProjectB", "Branch2")
	if err != nil {
		t.Errorf("FindBranchData() returned an unexpected error")
		return
	}

	_, err = history.FindBranchData("ProjectC", "Branch1")
	if err != nil {
		t.Errorf("FindBranchData() returned an unexpected error")
		return
	}
}
