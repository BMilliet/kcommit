package src

import (
	"encoding/json"
	"fmt"
	"os"
)

type UtilsInterface interface {
	CommitTypeDTOsToListItems(commitTypes []CommitTypeDTO) []ListItem
	ValidateInput(v string)
	HandleError(err error, message string)
	ExitWithError(message string)
}

type Utils struct{}

func NewUtils() *Utils {
	return &Utils{}
}

func (u *Utils) CommitTypeDTOsToListItems(commitTypes []CommitTypeDTO) []ListItem {
	var listItems []ListItem
	for _, commitType := range commitTypes {
		listItems = append(listItems, ListItem{
			T: commitType.Type,
			D: commitType.Description,
		})
	}
	return listItems
}

func (u *Utils) ValidateInput(v string) {
	if v == ExitSignal {
		os.Exit(0)
	}
}

func (u *Utils) HandleError(err error, message string) {
	if err != nil {
		msg := fmt.Sprintf(message+" -> ", err.Error())
		u.ExitWithError(msg)
	}
}

func (u *Utils) ExitWithError(message string) {
	s := DefaultStyles()
	println((s.Text(message, s.ErrorColor)))
	os.Exit(1)
}

func ParseJSONContent[T any](jsonString string) (*T, error) {
	var targetStruct T
	err := json.Unmarshal([]byte(jsonString), &targetStruct)
	if err != nil {
		return nil, fmt.Errorf("ParseJSONContent -> %v", err)
	}
	return &targetStruct, nil
}
