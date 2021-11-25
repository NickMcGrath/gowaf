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
	nodePack NodePack
	title    string
	data     string
}

func (myc myComponent) Render() {
	myc.nodePack.setNode("<h1>" + myc.title + "</h2>" +
		"<p>" + myc.data + "</p>")
}

func testMyComponent() {
	nf := NodePackFactory{}
	nf.newNode("div")
	pack := NodePack{getNode: nf.getNode, setNode: nf.setNode}
	component := myComponent{nodePack: pack,
		title: "my component title",
		data:  "my data"}
	component.Render()
	component.title = "new title"
	component.data = "new data"
	time.Sleep(2 * time.Second)
	component.Render()
}

func main() {

	testMyComponent()

}
