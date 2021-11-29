package main

import (
	"github.com/nickmcgrath/gowaf/example/components"
)

func testMyComponent() {
	components.CreateMyComponent("test", "ing").Component.Render()
}

func main() {
	testMyComponent()
}
