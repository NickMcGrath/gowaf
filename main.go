package main

import (
	"fmt"
	"strings"
	"syscall/js"
	"time"
)

type nodeWrapper struct {
	node       *js.Value
	nodeString string
}

func (n *nodeWrapper) getNode() *js.Value {
	return n.node
}
func (n *nodeWrapper) setNode(node *js.Value) {
	n.node = node
}
func (n *nodeWrapper) newNode(nodeType string, parentId string) {
	dom := js.Global().Get("document")
	newElement := dom.Call("createElement", nodeType)
	body := dom.Call("getElementById", parentId)
	body.Call("appendChild", newElement)
	n.node = &newElement
}
func (n *nodeWrapper) setInnerHTML(innerHTML ...string) {
	n.nodeString = strings.Join(innerHTML, "")
	n.node.Set("innerHTML", strings.Join(innerHTML, ""))
}
func (n *nodeWrapper) String() string {
	return n.nodeString
}

//NodePackFactory gives NodePack the functions needed
type NodePackFactory struct {
}

func (nf *NodePackFactory) newNode(nodeType string, parentId string) *nodeWrapper {
	var nw nodeWrapper
	nw.newNode(nodeType, parentId)
	return &nw
}

//MyComponent is an example of a users class.
//They can do what ever they want and when ever they want to render they call nodePack.SetNode()
type myComponent struct {
	nodeWrap         *nodeWrapper
	title            string
	data             string
	myOtherComponent myOtherComponent
}

func (myc *myComponent) Render() {
	fmt.Println("<h1>", myc.title, "</h2>",
		"<p>", myc.data, "</p>", myc.myOtherComponent.nodeWrap.String())

	myc.nodeWrap.setInnerHTML("<h1>", myc.title, "</h2>",
		"<p>", myc.data, "</p>", myc.myOtherComponent.nodeWrap.String())
}

type myOtherComponent struct {
	nodeWrap  *nodeWrapper
	paragraph string
}

func (myc *myOtherComponent) Render() {
	myc.nodeWrap.setInnerHTML("<h1>", myc.paragraph, "</h2>")
}

func testMyComponent() {
	var nf NodePackFactory
	component2 := myOtherComponent{nodeWrap: nf.newNode("div", "bodyId"), paragraph: "my paragraph"}
	component2.Render()
	component2.paragraph = "yolo"

	component := myComponent{
		title:            "my component title",
		data:             "my data",
		nodeWrap:         nf.newNode("div", "bodyId"),
		myOtherComponent: component2}

	component.Render()
	component.title = "new title"
	component.data = "new data"
	time.Sleep(2 * time.Second)
	component2.Render()
	component.Render()
}

func main() {
	testMyComponent()
}
