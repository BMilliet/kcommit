package src

import (
	"encoding/json"
	"fmt"
	"time"
)

type HistoryModel struct {
	Projects map[string]map[string]BranchDetail `json:"projects"`
}

type BranchDetail struct {
	Scope     string    `json:"scope"`
	UpdatedAt time.Time `json:"updated_at"`
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
		return nil, fmt.Errorf("FindBranchData -> %s", projectName)
	}

	if !h.HasBranch(projectName, branchName) {
		return nil, fmt.Errorf("FindBranchData -> %s %s", branchName, projectName)
	}

	branch := h.Projects[projectName][branchName]
	return &branch, nil
}

func (h *HistoryModel) SetBranch(projectName string, branchName string, scope string) {
	h.Projects[projectName][branchName] = BranchDetail{Scope: scope, UpdatedAt: time.Now()}
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
		return fmt.Errorf("AddBranch -> %s %s", branchName, projectName)
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
		return "", fmt.Errorf("ToJson -> %v", err)
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
				Name:      branchName,
				Scope:     branchDetail.Scope,
				UpdatedAt: branchDetail.UpdatedAt,
			})
		}

		projects = append(projects, project)
	}

	return projects
}

type HistoryDTO struct {
	Projects []ProjectModel `json:"projects"`
}

type ProjectModel struct {
	Name     string        `json:"name"`
	Branches []BranchModel `json:"branches"`
}

type BranchModel struct {
	Name      string    `json:"name"`
	Scope     string    `json:"scope"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommitType struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type CommitRules struct {
	CommitTypes []CommitType `json:"commitTypes"`
}

func CreateHistoryModelFromDTO(dto *HistoryDTO) HistoryModel {
	history := HistoryModel{
		Projects: make(map[string]map[string]BranchDetail),
	}

	for _, project := range dto.Projects {
		projectBranches := make(map[string]BranchDetail)

		for _, branch := range project.Branches {
			projectBranches[branch.Name] = BranchDetail{
				Scope:     branch.Scope,
				UpdatedAt: branch.UpdatedAt,
			}
		}

		history.Projects[project.Name] = projectBranches
	}

	return history
}

// TODO: this logic needs to be adjusted. Couting just the year means cleaning the cache completely on day 1.
func (h *HistoryModel) CleanOldBranches() {
	currentTime := time.Now()
	currentYear := currentTime.Year()

	oneMonthAgo := currentTime.AddDate(0, -1, 0)

	for project, branches := range h.Projects {
		for branch, details := range branches {
			if details.UpdatedAt.Before(oneMonthAgo) || details.UpdatedAt.Year() < currentYear {
				delete(branches, branch)
			}
		}

		if len(branches) == 0 {
			delete(h.Projects, project)
		}
	}
}
