# Constants

Constants are a way to create a named identifier whose value can never change. They also provide an incredible amount of flexibility to the language. The way constants are implemented in Go is very unique.

- They exist only at compilation.
- Untyped constants can be implicitly converted where typed constants and variables can't.
- Think of untyped constants as having a Kind, not a Type.

## Examples

- [Declare and Initialize Constants](https://github.com/gkjoyes/golang-tour/blob/4f8064ff6bcd2e30fdc74f05507a29af13927bd8/lesson/02/syntax/constants/example1/example1.go)
- [Constants precisions](https://github.com/gkjoyes/golang-tour/blob/4f8064ff6bcd2e30fdc74f05507a29af13927bd8/lesson/02/syntax/constants/example2/example2.go)
- [iota](https://github.com/gkjoyes/golang-tour/blob/4f8064ff6bcd2e30fdc74f05507a29af13927bd8/lesson/02/syntax/constants/example3/example3.go)
- [Implicit conversion](https://github.com/gkjoyes/golang-tour/blob/4f8064ff6bcd2e30fdc74f05507a29af13927bd8/lesson/02/syntax/constants/example4/example4.go)
