# golang-learning

My first project in golang, which I am going to use for studying.
This repo started by following https://golang.org/doc/tutorial/create-module

Run the code:

- cd hello
- go build
- ./hello.exe

or just

- go run hello.go

Run tests:

- cd greetings
- go test -v

Compile and install the application

- cd hello
- go list -f '{{.Target}}' // To discover the install path
- go install
- hello

Study notes:

- Tour of go: https://tour.golang.org/list
- Why types look the way they do? https://blog.golang.org/declaration-syntax
