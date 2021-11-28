package main

import (
	"github.com/nickmcgrath/gowaf/example/components"
	"github.com/nickmcgrath/gowaf/nf"
	"time"
)

func testMyComponent() {
	component2 := components.MyOtherComponent{NodeWrap: nf.NewNode("div"), Paragraph: "my component"}
	component2.Render()

	component3 := components.MyOtherComponent2{NodeWrap: nf.NewNode("div"), Paragraph: "my component 2"}
	component3.Render()

	component := components.MyComponent{
		Title:             "my component title",
		Data:              "my data",
		NodeWrap:          nf.NewNodeFromParentId("div", "bodyId"),
		MyOtherComponent:  component2,
		MyOtherComponent2: component3}

	component.Render()

	component2.Paragraph = "new paragraph"
	time.Sleep(2 * time.Second)
	component2.Render()
	time.Sleep(2 * time.Second)
	component.Title = "new title"
	component.Data = "new data"
	component.Render()
}

func main() {
	testMyComponent()
}
