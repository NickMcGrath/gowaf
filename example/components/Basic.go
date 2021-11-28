package components

import (
	"fmt"
	"github.com/nickmcgrath/gowaf/nf"
	"github.com/nickmcgrath/gowaf/wraps"
	"syscall/js"
)

type ComponentProperties struct {
	NodeWrap *wraps.NodeWrapper
	//Title      string
	//Data       string
	childNodes []*wraps.NodeWrapper
	Render     func()
}

type MyComponent struct {
	Properties *ComponentProperties
	Title      string
	Data       string
}

type ChildComponent struct {
	Properties *ComponentProperties
	Title      string
}

func CreateMyComponent(titl string, dat string, children ...*wraps.NodeWrapper) *MyComponent {
	node := nf.NewNode("div")
	component := MyComponent{nil, titl, dat}
	node.Compose = func() {
		node.SetInnerHTML("<loltag>", component.Title, "</loltag>")
	}
	renderFunc := func() {
		node.Compose()
		body := js.Global().Get("document").Call("getElementById", "bodyId")
		body.Call("appendChild", *node.GetNode())
		for _, child := range children {
			child.Compose()
			fmt.Println(child.String())
			fmt.Println(node.String())
			(*node.GetNode()).Call("appendChild", child.GetNode())
		}
	}
	props := ComponentProperties{node, children, renderFunc}

	component.Properties = &props
	return &component
}

func CreateChildComponent(title string) *ChildComponent {
	node := nf.NewNode("div")
	component := ChildComponent{nil, title}
	node.Compose = func() {
		node.SetInnerHTML("<h2>", component.Title, "</h2>")
	}
	props := ComponentProperties{node, nil, func() {
		node.SetInnerHTML("<h2>", component.Title, "</h2>")
		//body := js.Global().Get("document").Call("getElementById", "bodyId")
		//body.Call("appendChild", *node.GetNode())
	}}

	component.Properties = &props
	return &component
}
