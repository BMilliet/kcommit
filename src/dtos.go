package src

import (
	"time"
)

type HistoryDTO struct {
	Projects []ProjectDTO `json:"projects"`
}

type ProjectDTO struct {
	Name     string      `json:"name"`
	Branches []BranchDTO `json:"branches"`
}

type BranchDTO struct {
	Name      string    `json:"name"`
	Scope     string    `json:"scope"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommitTypeDTO struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type CommitRulesDTO struct {
	CommitTypeDTOs []CommitTypeDTO `json:"commitTypes"`
}

func (dto *HistoryDTO) ToModel() History {
	history := History{
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
