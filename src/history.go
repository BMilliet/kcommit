package src

import (
	"encoding/json"
	"fmt"
	"time"
)

type History struct {
	Projects map[string]map[string]BranchDetail `json:"projects"`
}

type BranchDetail struct {
	Scope     string    `json:"scope"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (h *History) hasProject(projectName string) bool {
	_, projectExists := h.Projects[projectName]
	return projectExists
}

func (h *History) HasBranch(projectName string, branchName string) bool {
	projectBranches, projectExists := h.Projects[projectName]
	if !projectExists {
		return false
	}
	_, branchExists := projectBranches[branchName]
	return branchExists
}

func (h *History) FindBranchData(projectName string, branchName string) (*BranchDetail, error) {
	if !h.hasProject(projectName) {
		return nil, fmt.Errorf("FindBranchData -> %s", projectName)
	}

	if !h.HasBranch(projectName, branchName) {
		return nil, fmt.Errorf("FindBranchData -> %s %s", branchName, projectName)
	}

	branch := h.Projects[projectName][branchName]
	return &branch, nil
}

func (h *History) SetBranch(projectName string, branchName string, scope string) {
	h.Projects[projectName][branchName] = BranchDetail{Scope: scope, UpdatedAt: time.Now()}
}

func (h *History) addProject(projectName string) {
	if !h.hasProject(projectName) {
		h.Projects[projectName] = make(map[string]BranchDetail)
	}
}

func (h *History) AddBranch(projectName, branchName string) error {
	if !h.hasProject(projectName) {
		h.addProject(projectName)
	}

	if h.HasBranch(projectName, branchName) {
		return fmt.Errorf("AddBranch -> %s %s", branchName, projectName)
	}

	h.Projects[projectName][branchName] = BranchDetail{Scope: ""}
	return nil
}

func (h *History) ToJson() (string, error) {
	data := map[string][]ProjectDTO{
		"projects": h.ToProjectDTO(),
	}

	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("ToJson -> %v", err)
	}
	return string(jsonBytes), nil
}

func (h *History) ToProjectDTO() []ProjectDTO {
	var projects []ProjectDTO

	for projectName, branches := range h.Projects {
		project := ProjectDTO{
			Name:     projectName,
			Branches: []BranchDTO{},
		}

		for branchName, branchDetail := range branches {
			project.Branches = append(project.Branches, BranchDTO{
				Name:      branchName,
				Scope:     branchDetail.Scope,
				UpdatedAt: branchDetail.UpdatedAt,
			})
		}

		projects = append(projects, project)
	}

	return projects
}

func (h *History) CleanOldBranches(currentTime time.Time) {
	oneMonthAgo := currentTime.AddDate(0, -1, 0)

	for project, branches := range h.Projects {
		for branch, details := range branches {
			if details.UpdatedAt.Before(oneMonthAgo) {
				delete(branches, branch)
			}
		}

		if len(branches) == 0 {
			delete(h.Projects, project)
		}
	}
}
