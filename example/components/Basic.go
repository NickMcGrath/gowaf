package components

import (
	"fmt"
	"github.com/nickmcgrath/gowaf/wraps"
	"syscall/js"
)

//MyComponent is an example of a users class.
//They can do what ever they want and when ever they want to render they call nodePack.SetNode()
type Component struct {
	NodeWrap *wraps.NodeWrapper
	//Title      string
	//Data       string
	childNodes []*wraps.NodeWrapper
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
	myc.NodeWrap.Compose()
	body := js.Global().Get("document").Call("ElementById", "bodyId")
	body.Call("appendChild", myc.NodeWrap.GetNode())
	//myc.NodeWrap.SetInnerHTML("<h1>", myc.Paragraph, "</h2>")
}

//Child
type MyOtherComponent2 struct {
	NodeWrap  *wraps.NodeWrapper
	Paragraph string
}

func (myc *MyOtherComponent2) Render() {
	myc.NodeWrap.SetInnerHTML("<h1>", myc.Paragraph, "</h2>")
}
