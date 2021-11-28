package main

import (
	"github.com/nickmcgrath/gowaf/example/components"
	"time"
)

func testMyComponent() {
	child := components.CreateChildComponent("child")
	parent := components.CreateMyComponent("parent", "dat", child.Properties.NodeWrap)
	parent.Properties.Render()
	child.Title = "tits"
	time.Sleep(2 * time.Second)
	child.Properties.Render()
	//parent.Properties.Render()
	//child.Properties.Render()

}

func main() {
	testMyComponent()
}

