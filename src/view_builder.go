package src

type ViewBuilderInterface interface {
	NewListView(title string, op []ListItem, height int) ListItem
	NewTextFieldView(title, placeHolder string) string
}

type ViewBuilder struct{}

func NewViewBuilder() *ViewBuilder {
	return &ViewBuilder{}
}

func (b *ViewBuilder) NewListView(title string, op []ListItem, height int) ListItem {
	endValue := ListItem{}
	ListView(title, op, height, &endValue)
	return endValue
}

func (b *ViewBuilder) NewTextFieldView(title, placeHolder string) string {
	endValue := ""
	TextFieldView(title, placeHolder, &endValue)
	return endValue
}
