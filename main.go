package main

import (
	"fmt"
	"github.com/nickmcgrath/gowaf/example/components"
	"github.com/nickmcgrath/gowaf/nf"
	"strings"
	"syscall/js"
	"time"
)

func testMyComponent() {
	component2 := components.MyOtherComponent{NodeWrap: nf.NewNode("div"), Paragraph: "my component"}
	component2.Render()

	component3 := components.MyOtherComponent2{NodeWrap: nf.NewNode("div"), Paragraph: "my component 2"}
	component3.Render()

	component := components.MyComponent{
		Title:             "my component title",
		Data:              "my data",
		Integer:           5,
		NodeWrap:          nf.NewNodeFromParentId("div", "bodyId"),
		MyOtherComponent:  &component2,
		MyOtherComponent2: &component3}

	component.SetUp()

	component2.Paragraph = "new paragraph"
	time.Sleep(2 * time.Second)
	component2.Render()
	time.Sleep(2 * time.Second)
	component.Title = "new title"
	component.Data = "new data"
	component.Update()
}
func testBasic() {
	dom := js.Global().Get("document")
	newElement := dom.Call("createElement", "div")
	body := dom.Call("getElementById", "bodyId")
	body.Call("appendChild", newElement)

	str := "hello"
	integer := 123
	str2 := "World"
	//y := fmt.Sprint(str,integer)
	x := make([]interface{}, 0)
	x = append(x, str)
	x = append(x, &integer)
	x = append(x, &str2)

	builder := strings.Builder{}
	for _, v := range x {
		fmt.Printf("%T\n", v)
		switch v.(type) {
		case string:
			builder.WriteString(fmt.Sprint(v))
		case *string:
			builder.WriteString(fmt.Sprintf("%p", v))
		}
	}
	newElement.Set("innerText", builder.String())

}

func main() {
	testMyComponent()
	//testBasic()
}
