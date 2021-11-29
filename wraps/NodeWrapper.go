package wraps

import (
	"strings"
	"syscall/js"
)

// values we can get from a Node: https://www.w3schools.com/jsref/dom_obj_all.asp

type NodeWrapper struct {
	Name           string
	Node           *js.Value
	RenderCallBack func()
}

func (n *NodeWrapper) Render() {
	n.RenderCallBack()
	elements := js.Global().Get("document").Call("getElementsByTagName", n.Name)
	numElements := elements.Length()
	for i := 0; i < numElements; i++ {
		element := elements.Index(i)
		element.Set("innerHTML", "")
		element.Call("appendChild", *n.Node)
	}
}

func (n *NodeWrapper) GetNode() *js.Value {
	return n.Node
}
func (n *NodeWrapper) SetNode(Node *js.Value) {
	n.Node = Node
}

func BlankDiv() *js.Value {
	newElement := js.Global().Get("document").Call("createElement", "div")
	return &newElement
}

func (n *NodeWrapper) NewNode(element string) {
	dom := js.Global().Get("document")
	newElement := dom.Call("createElement", element)
	n.Node = &newElement
}

func (n *NodeWrapper) NewNodeOfParentId(element string, parentId string) {
	dom := js.Global().Get("document")
	newElement := dom.Call("createElement", element)
	body := dom.Call("getElementById", parentId)
	body.Call("appendChild", newElement)
	n.Node = &newElement
}

func (n *NodeWrapper) NewNodeOfParentWrap(element string, parent NodeWrapper) {
	dom := js.Global().Get("document")
	newElement := dom.Call("createElement", element)
	parent.Node.Call("appendChild", parent.Node)
	n.Node = &newElement
}

func (n *NodeWrapper) NewNodeOfChildWrap(element string, child NodeWrapper) {
	dom := js.Global().Get("document")
	newElement := dom.Call("createElement", element)
	newElement.Call("appendChild", child.Node)
	n.Node = &newElement
}

func (n *NodeWrapper) SetInnerHTML(innerHTML ...string) {
	n.Node.Set("innerHTML", strings.Join(innerHTML, ""))
}

//func (n *NodeWrapper) Compose(elements ...interface{}) {
//	for element := range elements {
//		switch reflect.TypeOf(element).Kind() {
//		case reflect.TypeOf("").Kind():
//			fmt.Println("string", element)
//		}
//	}
//}

func (n *NodeWrapper) String() string {
	return n.Node.Get("innerHTML").String()
}
