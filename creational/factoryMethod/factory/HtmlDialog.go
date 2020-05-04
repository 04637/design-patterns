package factory

import "design-patterns/creational/factoryMethod/buttons"

// 具体创建者
type HtmlDialog struct {

}

func (p *HtmlDialog) CreateButton() buttons.Button {
	return new(buttons.HtmlButton)
}