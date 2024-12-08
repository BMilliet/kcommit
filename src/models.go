package src

import (
	"encoding/json"
	"fmt"
)

type HistoryModel struct {
	Projects map[string]map[string]BranchDetail `json:"projects"`
}

type BranchDetail struct {
	Scope string `json:"scope"`
}

func (h *HistoryModel) hasProject(projectName string) bool {
	_, projectExists := h.Projects[projectName]
	return projectExists
}

func (h *HistoryModel) HasBranch(projectName string, branchName string) bool {
	projectBranches, projectExists := h.Projects[projectName]
	if !projectExists {
		return false
	}
	_, branchExists := projectBranches[branchName]
	return branchExists
}

func (h *HistoryModel) FindBranchData(projectName string, branchName string) (*BranchDetail, error) {
	if !h.hasProject(projectName) {
		return nil, fmt.Errorf("project '%s' not found", projectName)
	}

	if !h.HasBranch(projectName, branchName) {
		return nil, fmt.Errorf("branch '%s' not found in project '%s'", branchName, projectName)
	}

	branch := h.Projects[projectName][branchName]
	return &branch, nil
}

func (h *HistoryModel) addProject(projectName string) {
	if !h.hasProject(projectName) {
		h.Projects[projectName] = make(map[string]BranchDetail)
	}
}

func (h *HistoryModel) AddBranch(projectName, branchName string) error {
	if !h.hasProject(projectName) {
		h.addProject(projectName)
	}

	if h.HasBranch(projectName, branchName) {
		return fmt.Errorf("branch '%s' already exists in project '%s'", branchName, projectName)
	}

	h.Projects[projectName][branchName] = BranchDetail{Scope: ""}
	return nil
}

func (h *HistoryModel) ToJson() (string, error) {
	data := map[string][]ProjectModel{
		"projects": h.ToProjectModel(),
	}

	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error converting ProjectModel to JSON: %v", err)
	}
	return string(jsonBytes), nil
}

func (h *HistoryModel) ToProjectModel() []ProjectModel {
	var projects []ProjectModel

	for projectName, branches := range h.Projects {
		project := ProjectModel{
			Name:     projectName,
			Branches: []BranchModel{},
		}

		for branchName, branchDetail := range branches {
			project.Branches = append(project.Branches, BranchModel{
				Name:  branchName,
				Scope: branchDetail.Scope,
			})
		}

		projects = append(projects, project)
	}

	return projects
}

type HistoryDTO struct {
	Projects []ProjectModel `json:projects`
}

type ProjectModel struct {
	Name     string        `json:"name"`
	Branches []BranchModel `json:"branches"`
}

type BranchModel struct {
	Name  string `json:"name"`
	Scope string `json:"scope"`
}

type CommitType struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type CommitRules struct {
	CommitTypes []CommitType `json:"commitTypes"`
}

func DefaultRules() *CommitRules {
	l := []CommitType{
		{
			Type:        "feat",
			Description: "Adds a new feature to the project.",
		},
		{
			Type:        "fix",
			Description: "Fixes a bug in the code.",
		},
		{
			Type:        "chore",
			Description: "Auxiliary tasks, such as dependency updates or configuration changes.",
		},
		{
			Type:        "style",
			Description: "Changes that do not affect functionality (e.g., formatting, whitespace).",
		},
		{
			Type:        "refactor",
			Description: "Refactors code without changing existing functionality.",
		},
		{
			Type:        "test",
			Description: "Adds or updates automated tests.",
		},
		{
			Type:        "build",
			Description: "Changes related to the build system or external dependencies.",
		},
		{
			Type:        "revert",
			Description: "Reverts a previous commit.",
		},
		{
			Type:        "perf",
			Description: "Performance improvements in the code.",
		},
		{
			Type:        "ci",
			Description: "Changes to the continuous integration configuration.",
		},
		{
			Type:        "docs",
			Description: "Updates documentation only, without changing the code.",
		},
	}

	return &CommitRules{
		CommitTypes: l,
	}
}

func CreateHistoryModelFromDTO(dto *HistoryDTO) HistoryModel {
	history := HistoryModel{
		Projects: make(map[string]map[string]BranchDetail),
	}

	for _, project := range dto.Projects {
		projectBranches := make(map[string]BranchDetail)

		for _, branch := range project.Branches {
			projectBranches[branch.Name] = BranchDetail{
				Scope: branch.Scope,
			}
		}

		history.Projects[project.Name] = projectBranches
	}

	return history
}
