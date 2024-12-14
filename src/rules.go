package src

func DefaultRules() *CommitRulesDTO {
	l := []CommitTypeDTO{
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

	return &CommitRulesDTO{
		CommitTypeDTOs: l,
	}
}
