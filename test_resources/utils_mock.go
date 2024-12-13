package testresources

import (
	"kcommit/src"
)

type UtilsMock struct {
	ValidateInputReturnCalled int
	HandleErrorCalledWith     string
	ExitWithErrorCalledWith   string
}

func (u *UtilsMock) CommitTypesToListItems(commitTypes []src.CommitType) []src.ListItem {
	var listItems []src.ListItem
	for _, commitType := range commitTypes {
		listItems = append(listItems, src.ListItem{
			T: commitType.Type,
			D: commitType.Description,
		})
	}
	return listItems
}

func (u *UtilsMock) ValidateInput(v string) {
	u.ValidateInputReturnCalled += 1
}

func (u *UtilsMock) HandleError(err error, message string) {
	u.HandleErrorCalledWith = message
}

func (u *UtilsMock) ExitWithError(message string) {
	u.ExitWithErrorCalledWith = message
}
