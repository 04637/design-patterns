package buttons

import "fmt"

type WindowsButton struct {

}

func (p *WindowsButton) Render() {
	fmt.Println("render WindowsButton")
	p.OnClick()
}

func (p *WindowsButton) OnClick() {
	fmt.Println("Click! Button says - 'This is WindowsButton'")
}
