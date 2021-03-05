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

Compile and install the application:

- cd hello
- go list -f '{{.Target}}' // To discover the install path
- go install
- hello

Study notes:

- Tour of go: https://tour.golang.org/list
- Why types look the way they do? https://blog.golang.org/declaration-syntax
- defers: https://tour.golang.org/flowcontrol/13
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
