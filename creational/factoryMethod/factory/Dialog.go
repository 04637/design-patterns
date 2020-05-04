package factory

import "design-patterns/creational/factoryMethod/buttons"

// 基础创建者
type IDialog interface {
	CreateButton() buttons.Button
}

type Dialog struct {

}

func (p *Dialog) RenderWindow(dialog IDialog) {
	var okButton buttons.Button
	okButton = dialog.CreateButton()
	okButton.Render()
}
