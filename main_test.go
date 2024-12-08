package main

import (
	"kcommit/src"
	"testing"
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

	historyModel.SetBranchScope(newProjectName, newBranch1Name, "42")
	historyModel.SetBranchScope(newProjectName, newBranch2Name, "99")

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
		// Verifica se a branch existe
		if !historyModel.HasBranch(test.projectName, test.branchName) {
			t.Errorf("expected branch '%s' in project '%s' to exist, but it does not", test.branchName, test.projectName)
			continue
		}

		// Verifica os detalhes da branch
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
