# learning-go
Learning GoLang

# Introduction
- Compiled lang.
- Go tool can run file directly, without VM like Java has JVM
- Executables are different for OS, we can develeop binary for different OS
- Mindset is different than other lang like it doesn't has try/catch because it don't need it
- It misses a lot of features compared to other programming langs, again it does because it doesn't need them
- Object Oriented: Yes and No
- Designed with simplicity: What you see on the screen is the code: no overloading/overridding
- Lexer does a lot of work

# [go help command](https://www.youtube.com/watch?v=QEZlivtFOZk&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=4)
- `go help`: to list all command
- `go help run`: to get docs for `run` command
- `go help gopath`: to get docs for `gopath` command
- `go env GOPATH`: to get the `GOPATH`

# Build for different OS
- `go env`: list of all GO envs
- In Mac OS: `GOOS="darwin"`
- So, to build window exe: `GOOS="windows" go build` or `GOOS="linux" go build` for linux

# [Memory Management](https://www.youtube.com/watch?v=G1SP9uDJD0g&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=11)
- Memory allocation and deallocation happens automatically
- `new`
    - Allocate memory but not INIT
    - you'll get a memory address
    - zeroed storage
- `make`
    - Allocate memory and INIT
    - you'll get a memory address
    - non-zeroed storage
- GC happens automatically anything out of scope or `nil`

# [Slices]
- Slices in Go are a flexible and efficient way to represent arrays, and they are often used in place of arrays because of their dynamic size and added features