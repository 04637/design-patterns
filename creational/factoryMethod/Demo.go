package main

import "design-patterns/creational/factoryMethod/factory"

// 客户端代码
func main() {
	dialog := configure("windows")
	runBusinessLogic(dialog)
}

func configure(plat string) factory.IDialog{
	switch plat {
	case "html":
		return new(factory.HtmlDialog)
	case "windows":
		return new(factory.WindowsDialog)
	default:
		return nil
	}
}

func runBusinessLogic(dialog factory.IDialog) {
	var d factory.Dialog
	d.RenderWindow(dialog)
}
