package components

import (
	"fmt"
	"github.com/nickmcgrath/gowaf/wraps"
)

//MyComponent is an example of a users class.
//They can do what ever they want and when ever they want to render they call nodePack.SetNode()
type MyComponent struct {
	NodeWrap          *wraps.NodeWrapper
	Title             string
	Data              string
	MyOtherComponent  MyOtherComponent
	MyOtherComponent2 MyOtherComponent2
}

func (myc *MyComponent) Render() {
	fmt.Println("<h1>", myc.Title, "</h2>",
		"<p>", myc.Data, "</p>", myc.MyOtherComponent.NodeWrap)

	myc.NodeWrap.SetInnerHTML("<h1>", myc.Title, "</h2>",
		"<p>", myc.Data, "</p>", myc.MyOtherComponent.NodeWrap.String(), myc.MyOtherComponent2.NodeWrap.String())
}

//Child
type MyOtherComponent struct {
	NodeWrap  *wraps.NodeWrapper
	Paragraph string
}

func (myc *MyOtherComponent) Render() {
	myc.NodeWrap.SetInnerHTML("<h1>", myc.Paragraph, "</h2>")
}

//Child
type MyOtherComponent2 struct {
	NodeWrap  *wraps.NodeWrapper
	Paragraph string
}

func (myc *MyOtherComponent2) Render() {
	myc.NodeWrap.SetInnerHTML("<h1>", myc.Paragraph, "</h2>")
}
