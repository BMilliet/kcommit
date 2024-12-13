package src

type ViewBuilderInterface interface {
	NewListView(title string, op []ListItem, height int) string
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
