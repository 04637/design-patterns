package buttons

// 通用产品接口 按钮
type Button interface {
	Render()
	OnClick()
}
