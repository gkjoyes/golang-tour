# Pointers

Pointers provide a way to share data across program boundaries. Having the ability to share and reference data with a pointer provides the benefit of efficiency. There is only one copy of the data and everyone can see it changing. The cost is that anyone can change the data which can cause side effects in the running program.

- Use pointers to share data.
- Values in Go are always pass by value.
- The `*` operator declares a pointer variable.
- `new` allocates zeroed storage in memory and returns a pointer to it.

## Escape Analysis

- When a value could be referenced after the function that constructs the value returns.
- When the compiler determines a value is too large to fit on the stack.
- When the compiler does not know the size of a value at compile time.
- When a value is decoupled through the use of function or interface values.

## Stack vs Heap

The stack is for data that needs to persist only for the lifetime of the function that constructs it and is reclaimed without any cost when the function exits. The heap is for data that needs to persist after the function that constructs it exits and is reclaimed by a sometimes costly garbage collection.

## Examples

- [Pass by Value](https://github.com/gkjoyes/golang-tour/blob/master/lesson/02/syntax/pointers/example1/example1.go)
- [Using a pointer to share data](https://github.com/gkjoyes/golang-tour/blob/master/lesson/02/syntax/pointers/example2/example2.go)
- [Points to the object](https://github.com/gkjoyes/golang-tour/blob/master/lesson/02/syntax/pointers/example3/example3.go)
- [Mechanics of escape analysis](https://github.com/gkjoyes/golang-tour/blob/master/lesson/02/syntax/pointers/example4/example4.go)
- [How stack grow with a program](https://github.com/gkjoyes/golang-tour/blob/master/lesson/02/syntax/pointers/example5/examples5.go)
