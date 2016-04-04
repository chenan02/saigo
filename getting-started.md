# Getting Started


  1. Install [GVM](https://github.com/moovweb/gvm)

  ```bash
  bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
  ```

  Place this line in `~/.zshrc` or `~/.bashrc` to source the GVM directory
  ```bash
  [[ -s "$HOME/.gvm/scripts/gvm" ]] && source "$HOME/.gvm/scripts/gvm"
  ```

  Reopen your terminal or run `source $HOME/.gvm/scripts/gvm`

  2. Install Go

  ```bash
  gvm install go1.4
  ```

  Read about the Go command from the Go documentation: https://golang.org/doc/articles/go_command.html

  3. Set your $GOPATH accordingly. `mygo` can be whatever name you wish to choose for your workspace

  ```bash
  $ mkdir $HOME/mygo
  $ export GOPATH=$HOME/mygo
  ```

  4. Run your first program

  ```bash
  $ mkdir -p $GOPATH/src/github.com/user
  $ mkdir $GOPATH/src/github.com/user/hello
  $ cd $GOPATH/src/github.com/user/hello
  ```

  Next, create a file named hello.go inside the hello directory, containing the following Go code.

  ```go
  package main

  import "fmt"

  func main() {
        fmt.Printf("Hello, world.\n")
  }
  ```

  Now you can build and install that program with the go tool:

  ```bash
  $ go install
  ```

  This command builds the hello command, producing an executable binary. It then installs that binary to the workspace's bin directory as hello

  ```bash
  $ $GOPATH/bin/hello
  Hello, world.
  ```

  Once you have added $GOPATH/bin to your PATH, just type the binary name:
  ```bash
  $ hello
  Hello, world.
  ```

# Exercises

The exercises in this repo assume that the learner has traversed a basic [tutorial](http://www.tutorialspoint.com/go/index.htm) on Go. If you find the first exercise overwhelming, you should probably try working through a tutorial and returning.

# Recommended Resources

There's no need to read through all of these resources but keep them handy when you need a reminder.

  1. [How to Write Go Code](https://golang.org/doc/code.html): This document demonstrates the development of a simple Go package and introduces the go tool, the standard way to fetch, build, and install Go packages and commands.
  2. [Effective Go](https://golang.org/doc/effective_go.html) : All the basic data types, control structures, style guide explained through examples.
  3. [A Tour of Go](https://tour.golang.org/welcome/1): An interactive tutorial for playing with Go
  4. [Go Playground](https://play.golang.org/) : A useful resource to write code in the browser
