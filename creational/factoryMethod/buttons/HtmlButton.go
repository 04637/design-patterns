package buttons

import "fmt"

// 具体产品 Html按钮
type HtmlButton struct {
}

func (p *HtmlButton) Render() {
	fmt.Println("render HtmlButton")
	p.OnClick()
}

func (p *HtmlButton) OnClick() {
	fmt.Println("Click! Button says - 'This is HtmlButton'")
}
