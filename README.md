Go Lang WebAssembly Framework

Diagram:
https://lucid.app/lucidchart/a09db756-5a66-4feb-a40f-d36833948040/edit?viewport_loc=-10%2C-10%2C2389%2C1233%2C0_0&invitationId=inv_cbbca1a4-48c0-4f4b-8a31-23588f8c35bf

compiling go to wasm

    Setting Go environment:
        go env -w GOOS=js GOARCH=wasm
        GOARCH=amd64, GOOS=windows

    Compiling:
        go build -o main.wasm
