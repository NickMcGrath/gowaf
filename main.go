package main

import (
	"syscall/js"
	"time"
)

//NodePack must be added to all components
type NodePack struct {
	getNode func() *js.Value
	setNode func(innerHTML string)
}

//NodePackFactory gives NodePack the functions needed
type NodePackFactory struct {
	node *js.Value
}

func (nf *NodePackFactory) newNode(nodeType string) {
	dom := js.Global().Get("document")
	newElement := dom.Call("createElement", nodeType)
	body := dom.Call("getElementById", "bodyId")
	body.Call("appendChild", newElement)
	nf.node = &newElement
}
func (nf *NodePackFactory) getNode() *js.Value {
	return nf.node
}

func (nf *NodePackFactory) setNode(innerHTML string) {
	nf.node.Set("innerHTML", innerHTML)
}

//MyComponent is an example of a users class.
//They can do what ever they want and when ever they want to render they call nodePack.SetNode()
type myComponent struct {
	nodePack         NodePack
	title            string
	data             string
	myOtherComponent myOtherComponent
}

func (myc myComponent) Render() {
	myc.nodePack.setNode("<h1>" + myc.title + "</h2>" +
		"<p>" + myc.data + "</p>" + myc.myOtherComponent.nodePack.getNode().String())
}

type myOtherComponent struct {
	nodePack  NodePack
	paragraph string
}

func (myc myOtherComponent) Render() {
	myc.nodePack.setNode("<h1>" + myc.paragraph + "</h2>")
}

func testMyComponent() {
	nf := NodePackFactory{}
	nf.newNode("div")
	pack := NodePack{getNode: nf.getNode, setNode: nf.setNode}
	nf2 := NodePackFactory{}
	nf2.newNode("div")
	pack2 := NodePack{getNode: nf2.getNode, setNode: nf2.setNode}
	component2 := myOtherComponent{nodePack: pack2, paragraph: "my paragraph"}
	component2.Render()
	component2.paragraph = "yolo"

	component := myComponent{nodePack: pack,
		title:            "my component title",
		data:             "my data",
		myOtherComponent: component2}
	component.Render()
	component.title = "new title"
	component.data = "new data"
	component2.Render()
	time.Sleep(2 * time.Second)
	component.Render()
}

func main() {

	testMyComponent()

}
