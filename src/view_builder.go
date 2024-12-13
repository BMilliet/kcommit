package src

type ViewBuilderInterface interface {
	NewListView(title string, op []ListItem, height int) string
	NewTextFieldView(title, placeHolder string) string
}

type ViewBuilder struct {
}

func NewViewBuilder() *ViewBuilder {
	return &ViewBuilder{}
}

func (b *ViewBuilder) NewListView(title string, op []ListItem, height int) string {
	endValue := ""
	ListView(title, op, height, &endValue)
	return endValue
}

func (b *ViewBuilder) NewTextFieldView(title, placeHolder string) string {
	endValue := ""
	TextFieldView(title, placeHolder, &endValue)
	return endValue
}
