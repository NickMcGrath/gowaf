#Go Lang WebAssembly Framework
![Gopher image](https://raw.githubusercontent.com/egonelbre/gophers/10cc13c5e29555ec23f689dc985c157a8d4692ab/vector/projects/surfing-js.svg)
*Image learning javascript*

##Diagram:
    https://lucid.app/lucidchart/a09db756-5a66-4feb-a40f-d36833948040/edit?viewport_loc=-10%2C-10%2C2389%2C1233%2C0_0&invitationId=inv_cbbca1a4-48c0-4f4b-8a31-23588f8c35bf

##Environment Setup:
    https://www.jetbrains.com/help/go/webassembly-project.html

##compiling go to wasm:
###Setting Go environment:
go env -w GOOS=js GOARCH=wasm
###Compiling:
go build -o main.wasm
