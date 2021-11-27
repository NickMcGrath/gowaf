package nf

import "github.com/nickmcgrath/gowaf/wraps"

//NodePackFactory gives NodePack the functions needed

func NewNodeFromParentId(nodeType string, parentId string) *wraps.NodeWrapper {
	var nw wraps.NodeWrapper
	nw.NewNodeOfParentId(nodeType, parentId)
	return &nw
}

func NewNodeOfParentWrap(nodeType string, wrapper wraps.NodeWrapper) *wraps.NodeWrapper {
	var nw wraps.NodeWrapper
	nw.NewNodeOfParentWrap(nodeType, wrapper)
	return &nw
}

func NewNodeOfChildWrap(nodeType string, wrapper wraps.NodeWrapper) *wraps.NodeWrapper {
	var nw wraps.NodeWrapper
	nw.NewNodeOfChildWrap(nodeType, wrapper)
	return &nw
}

func NewNode(nodeType string) *wraps.NodeWrapper {
	var nw wraps.NodeWrapper
	nw.NewNode(nodeType)
	return &nw
}
