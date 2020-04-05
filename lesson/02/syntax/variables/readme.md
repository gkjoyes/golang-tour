# Lesson 02: Language Syntax

## Variables

Variables are at the heart of the language and provide the ability to read from and write to memory. In Go, access to memory is type safe. This means the compiler takes type seriously and will not allow us to use variables
outside the scope of how they are declared.

- There is two-way to declare and initialize a variable:
  - Using the var keyword (`var foo int`).
  - Using the short-hand operator (`foo := 20`).

- We can't redeclare variables but can shadow them.

- Visibility
  - Variables declared inside a function or block cannot be accessed outside the function or block.
  - Package level variables with the first letter starting with lowercase are only available in that package.
  - Package level variables with the first letter starting with uppercase are for exporting.

- Naming conventions
  - __Pascal__ or __Camel__ case.
  - Capitalize acronyms (HTTP, URL etc...)
  - As short as reasonable.
  - Longer names for longer lives.

- Type conversion
  - T(v) converts the value v to the type T.
  - Use __strconv__ package for strings.s

## Examples

- [Declare and Initialize variable](https://github.com/george-kj/golang-tour/lesson/02/syntax/variables/example.go)

## Reference

- [Unlimited Go by William Kennedy](https://learning.oreilly.com/videos/ultimate-go-programming/9780135261651/9780135261651-UGP2_01_02_01)
