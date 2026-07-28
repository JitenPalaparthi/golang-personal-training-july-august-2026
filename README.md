
## Go commands

```bash 
# To create a module
    go mod init demo
# To run go application

go run hello.go

# To files bcz new function is create in greet.go
go run hello.go greet.go 

go run . 

# where the binary is created
go run --work .

# build go application,debug build

go build .

# debug build with name 

# dwarf and symbol information

go build -o debug_hello_demo . 

# release build

go build -ldflags="-w -s" -o release_hello_demo .

# to see symbols and linker stuff 

go tool nm debug_hello_demo

# cross compilatio or cross targer builds
# compilation targets
aix/ppc64
android/386
android/amd64
android/arm
android/arm64
darwin/amd64
darwin/arm64
dragonfly/amd64
freebsd/386
freebsd/amd64
freebsd/arm
freebsd/arm64
illumos/amd64
ios/amd64
ios/arm64
js/wasm
linux/386
linux/amd64
linux/arm
linux/arm64
linux/loong64
linux/mips
linux/mips64
linux/mips64le
linux/mipsle
linux/ppc64
linux/ppc64le
linux/riscv64
linux/s390x
netbsd/386
netbsd/amd64
netbsd/arm
netbsd/arm64
openbsd/386
openbsd/amd64
openbsd/arm
openbsd/arm64
openbsd/ppc64
openbsd/riscv64
plan9/386
plan9/amd64
plan9/arm
solaris/amd64
wasip1/wasm
windows/386
windows/amd64
windows/arm64

# GOOS and GOARCH

#cross compilation for linux
GOOS=linux GOARCH=amd64 go build -o linux-amd64-hello .
```