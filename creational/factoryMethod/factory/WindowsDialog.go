package factory

import "design-patterns/creational/factoryMethod/buttons"

type WindowsDialog struct {

}

func (p *WindowsDialog) CreateButton() buttons.Button {
	return new(buttons.WindowsButton)
}
