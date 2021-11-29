package components

import (
	"github.com/nickmcgrath/gowaf/wraps"
)

type Component struct {
	NodeWrap   *wraps.NodeWrapper
	childNodes []*wraps.NodeWrapper
}

func (c *Component) Render() {
	c.NodeWrap.Render()
	for _, node := range c.childNodes {
		node.Render()
	}
}

type BigBox struct {
	Component *Component
	Title     string
	Data      string
}

type ChildComponent struct {
	Component *Component
	Title     string
}

func CreateMyComponent(title string, data string) *BigBox {
	comp := &Component{nil, nil}
	box := &BigBox{nil, title, data}
	newElement := wraps.BlankDiv()
	nodeWrap := wraps.NodeWrapper{"BigBox", newElement, func() {
		newElement.Set("innerHTML", title+data)
	}}
	comp.NodeWrap = &nodeWrap
	box.Component = comp
	return box
}

// func CreateMyComponent(titl string, dat string, children ...*wraps.NodeWrapper) *MyComponent {
// 	node := nf.NewNode("div")
// 	component := MyComponent{nil, titl, dat}
// 	node.Compose = func() {
// 		node.SetInnerHTML("<loltag>", component.Title, "</loltag>")
// 	}
// 	renderFunc := func() {
// 		node.Compose()
// 		body := js.Global().Get("document").Call("getElementById", "bodyId")
// 		body.Call("appendChild", *node.GetNode())
// 		for _, child := range children {
// 			child.Compose()
// 			fmt.Println(child.String())
// 			fmt.Println(node.String())
// 			(*node.GetNode()).Call("appendChild", child.GetNode())
// 		}
// 	}
// 	props := ComponentProperties{node, children, renderFunc}

// 	component.Properties = &props
// 	return &component
// }

// func CreateChildComponent(title string) *ChildComponent {
// 	node := nf.NewNode("div")
// 	component := ChildComponent{nil, title}
// 	node.Compose = func() {
// 		node.SetInnerHTML("<h2>", component.Title, "</h2>")
// 	}
// 	props := ComponentProperties{node, nil, func() {
// 		node.SetInnerHTML("<h2>", component.Title, "</h2>")
// 		//body := js.Global().Get("document").Call("getElementById", "bodyId")
// 		//body.Call("appendChild", *node.GetNode())
// 	}}

// 	component.Properties = &props
// 	return &component
// }
