# GO
Let's go

### Download

wget https://dl.google.com/go/go1.13.7.linux-armv6l.tar.gz -O go.tar.gz

### Extract (RPI)

sudo tar -C /usr/local -xzf go.tar.gz

### Terminal setup

nano ~/.bashrc
export GOPATH=$HOME/go
export PATH=/usr/local/go/bin:$PATH:$GOPATH/bin
source ~/.bashrc

### Hello World

nano hello-world.go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}

### Run Program

go run hello-world.go
