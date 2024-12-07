package src

import (
	"encoding/json"
	"fmt"
)

func ParseJSONContent[T any](jsonString string) (*T, error) {
	var targetStruct T
	err := json.Unmarshal([]byte(jsonString), &targetStruct)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}
	return &targetStruct, nil
}

func CommitTypesToListItems(commitTypes []CommitType) []ListItem {
	var listItems []ListItem
	for _, commitType := range commitTypes {
		listItems = append(listItems, ListItem{
			Title: commitType.Type,
			Desc:  commitType.Description,
		})
	}
	return listItems
}
