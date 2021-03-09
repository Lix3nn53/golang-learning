# golang-learning

My first project in golang, which I am going to use for studying.
This repo started by following https://golang.org/doc/tutorial/create-module

### Run the code:

- cd hello
- go build
- ./hello.exe

or just

- go run hello.go

### Run tests:

- cd greetings
- go test -v

### Compile and install the application:

- cd hello
- go list -f '{{.Target}}' // To discover the install path
- go install
- hello

## Study go tour

Study codes are in `tour` folder

- Tour of go: https://tour.golang.org/list
- Why types look the way they do? https://blog.golang.org/declaration-syntax

### Flow control statements: for, if, else, switch and defer

- defers: https://tour.golang.org/flowcontrol/13

### More types: structs, slices, and maps.

- Unlike C, Go has no pointer arithmetic.
- A struct is a collection of fields.
- Struct fields can be accessed through a struct pointer. To access the field X of a struct when we have the struct pointer p we could write (\*p).X. However that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
- An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.
- a[low : high]
- [3]bool{true, true, false} => array, []bool{true, true, false} => slice
- Appending to a slice: If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.
- Functions are values too. They can be passed around just like other values.
- Go functions may be closures. A closure is a function value that references variables from outside its body.

### Methods and interfaces

- A method is a function with a special receiver argument.
- Remember: a method is just a function with a receiver argument.
- You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
- You can declare methods with pointer receivers: https://tour.golang.org/methods/4
- Functions with a pointer argument must take a pointer while methods with pointer receivers take either a value or a pointer as the receiver when they are called: https://tour.golang.org/methods/6
- Choosing a value or pointer receiver: https://tour.golang.org/methods/8
- Interfaces: https://tour.golang.org/methods/9
- Interfaces are implemented implicitly: https://tour.golang.org/methods/10
- The empty interface: https://tour.golang.org/methods/14
- Type assertions: https://tour.golang.org/methods/15
- Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
- Readers: https://tour.golang.org/methods/21
- Images exercise: https://tour.golang.org/methods/25

### Concurrency

- A goroutine is a lightweight thread managed by the Go runtime.
- `go f(x, y, z)` The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.
- Channels are a typed conduit through which you can send and receive values with the channel operator, <-. https://tour.golang.org/concurrency/2
- Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel: https://tour.golang.org/concurrency/3
- A sender can close a channel to indicate that no more values will be sent. https://tour.golang.org/concurrency/4
- The select statement lets a goroutine wait on multiple communication operations. A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready. https://tour.golang.org/concurrency/5
- The default case in a select is run if no other case is ready.
- Exercise: Equivalent Binary Trees
- sync.Mutex: https://tour.golang.org/concurrency/9

## Study Google I/O 2012

https://youtu.be/f6kdp27TYZs

- Study codes are in `patterns` folder

## Study Codewalk: Share Memory By Communicating

https://golang.org/doc/codewalk/sharemem/

- Study codes are in `codewalk` folder