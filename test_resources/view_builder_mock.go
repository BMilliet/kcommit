package testresources

import (
	"kcommit/src"
)

type ViewBuilderMock struct {
	NewListViewReturnValue      string
	NewListViewCalled           int
	NewTextFieldViewReturnValue string
	NewTextFieldViewCalled      int
}

func (b *ViewBuilderMock) NewListView(title string, op []src.ListItem, height int) src.ListItem {
	b.NewListViewCalled += 1
	return src.ListItem{
		T: b.NewListViewReturnValue,
	}
}

func (b *ViewBuilderMock) NewTextFieldView(title, placeHolder string) string {
	b.NewTextFieldViewCalled += 1
	return b.NewTextFieldViewReturnValue
}
