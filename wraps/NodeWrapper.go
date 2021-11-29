package wraps

import (
	"fmt"
	"strings"
	"syscall/js"
)

// values we can get from a node: https://www.w3schools.com/jsref/dom_obj_all.asp

//type NodeTree struct {
//	Head     *js.Value
//	Children []NodeTree
//	Values   []*interface{}
//}

type NodeWrapper struct {
	node     *js.Value
	template string
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
func (n *NodeWrapper) SetChildNode(child *NodeWrapper) {
	n.node.Call("appendChild", child.GetNode())
}

func (n *NodeWrapper) SetInnerHTML(innerHTML ...interface{}) {
	builder := strings.Builder{}
	for _, v := range innerHTML {
		builder.WriteString(fmt.Sprint(v))
	}
	n.node.Set("innerHTML", builder.String())
}
func (n *NodeWrapper) SetTemplateAndValues(innerTemplate string, values ...interface{}) {
	n.template = innerTemplate
	fmt.Println(fmt.Sprintf(innerTemplate, values...))
	n.node.Set("innerHTML", fmt.Sprintf(innerTemplate, values...))
}
func (n *NodeWrapper) UpdateValues(values ...interface{}) {
	fmt.Println(fmt.Sprintf(n.template, values...))
	n.node.Get("firstElementChild").Set("innerHTML", fmt.Sprintf(n.template, values...))
}

func (n *NodeWrapper) String() string {
	return n.node.Get("innerHTML").String()
}
