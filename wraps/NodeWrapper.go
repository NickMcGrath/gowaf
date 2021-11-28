package wraps

import (
	"strings"
	"syscall/js"
)

// values we can get from a node: https://www.w3schools.com/jsref/dom_obj_all.asp

type NodeWrapper struct {
	node *js.Value
}

func (n *NodeWrapper) GetNode() *js.Value {
	return n.node
}
func (n *NodeWrapper) SetNode(node *js.Value) {
	n.node = node
}

func (n *NodeWrapper) NewNode(element string) {
	dom := js.Global().Get("document")
	newElement := dom.Call("createElement", element)
	n.node = &newElement
}

func (n *NodeWrapper) NewNodeOfParentId(element string, parentId string) {
	dom := js.Global().Get("document")
	newElement := dom.Call("createElement", element)
	body := dom.Call("getElementById", parentId)
	body.Call("appendChild", newElement)
	n.node = &newElement
}

func (n *NodeWrapper) NewNodeOfParentWrap(element string, parent NodeWrapper) {
	dom := js.Global().Get("document")
	newElement := dom.Call("createElement", element)
	parent.node.Call("appendChild", parent.node)
	n.node = &newElement
}

func (n *NodeWrapper) NewNodeOfChildWrap(element string, child NodeWrapper) {
	dom := js.Global().Get("document")
	newElement := dom.Call("createElement", element)
	newElement.Call("appendChild", child.node)
	n.node = &newElement
}

func (n *NodeWrapper) SetInnerHTML(innerHTML ...string) {
	n.node.Set("innerHTML", strings.Join(innerHTML, ""))
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
	return n.node.Get("innerHTML").String()
}
