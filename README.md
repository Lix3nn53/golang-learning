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

## Study notes

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
- ```
  func WordCount(s string) map[string]int {
  fields := strings.Fields(s)
  fmt.Printf("%#v\n", fields)
  m := make(map[string]int)

  for i, s := range fields {
  fmt.Println(i, s)
  m[s] = m[s] + 1
  }

  return m
  }
  ```

- Functions are values too. They can be passed around just like other values.
- Go functions may be closures. A closure is a function value that references variables from outside its body.
- ```
  func adder() func(int) int {
    sum := 0
    return func(x int) int {
      sum += x
      return sum
    }
  }

  func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 3; i++ {
      fmt.Println(
        pos(i),
        neg(-2*i),
      )
    }
  }
  ```

- ```
  // fibonacci is a function that returns
  // a function that returns an int.
  func fibonacci() func() int {
    num1, num2 := 0, 0

    return func() int {
      if num2 == 0 {
        num2 = 1
        return 0
      }

          num1, num2 = num2, num1 + num2
          return num1
    }
  }
  func main() {
    f := fibonacci()
    for i := 0; i < 15; i++ {
    fmt.Println(f())
    }
  }
  ```

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
- ```
  package main

  import "fmt"

  type IPAddr [4]byte

  // TODO: Add a "String() string" method to IPAddr.
  func (p IPAddr) String() string {
    return fmt.Sprintf("%v.%v.%v.%v", p[0], p[1], p[2], p[3])
  }

  func main() {
    hosts := map[string]IPAddr{
      "loopback":  {127, 0, 0, 1},
      "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
      fmt.Printf("%v: %v\n", name, ip)
    }
  }
  ```

- Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
- ```
  package main

  import (
    "fmt"
  )

  type ErrNegativeSqrt struct {
    Value float64
  }

  func (e *ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", e.Value)
  }

  func Sqrt(x float64) (float64, error) {
    if (x < 0) {
      return 0, &ErrNegativeSqrt {
        x,
      }
    }

    z := 1.0

    diff := 1.0
    for diff > 0.0001 {
      a := z
      z -= (z*z - x) / (2*z)
      diff = a - z
      if (diff < 0) {
        diff *= -1
      }
    }

    return z, nil
    // return 0, nil
  }

  func main() {
    fmt.Println(Sqrt(4))
    fmt.Println(Sqrt(-2))
  }
  ```

- Readers: https://tour.golang.org/methods/21
- ```
  package main

  import (
    "io"
    "os"
    "strings"
  )

  type rot13Reader struct {
    r io.Reader
  }

  func rot13(x byte) byte {
      switch {
      case x >= 65 && x <= 77:
          fallthrough
      case x >= 97 && x <= 109:
          x = x + 13
      case x >= 78 && x <= 90:
          fallthrough
      case x >= 110 && x <= 122:
          x = x - 13
      }
      return x
  }

  func (r13 *rot13Reader) Read(b []byte) (int, error) {
      n, err := r13.r.Read(b)
      for i := 0; i <= n; i++ {
          b[i] = rot13(b[i])
      }
      return n, err
  }

  func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
  }
  ```

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

```
  package main

  import (
      "golang.org/x/tour/tree"
    "fmt"
  )

  // Walk walks the tree t sending all values
  // from the tree to the channel ch.
  func Walk(t *tree.Tree, ch chan int) {
      WalkRecursive(t, ch)
      close(ch)
  }

  func WalkRecursive(t *tree.Tree, ch chan int) {
      if t != nil {
          WalkRecursive(t.Left, ch)
          ch <- t.Value
          WalkRecursive(t.Right, ch)
      }
  }

  // Same determines whether the trees
  // t1 and t2 contain the same values.
  func Same(t1, t2 *tree.Tree) bool {
      ch1, ch2 := make(chan int), make(chan int)
      go Walk(t1, ch1)
      go Walk(t2, ch2)
      for {
          n1, ok1 := <- ch1
          n2, ok2 := <- ch2
      fmt.Printf("%v %v\n", n1, ok1)
      fmt.Printf("%v %v\n", n2, ok2)
          if ok1 != ok2 || n1 != n2 {
            return false
          }
          if !ok1 {
            break;
          }
      }
      return true
  }

  func main() {
    ch := make(chan int)
      go Walk(tree.New(1), ch)

    for i := 0; i < 10; i++ {

    }

      fmt.Println(Same(tree.New(1), tree.New(2)))
      fmt.Println(Same(tree.New(1), tree.New(1)))
      fmt.Println(Same(tree.New(2), tree.New(1)))
  }
```

- sync.Mutex: https://tour.golang.org/concurrency/9
