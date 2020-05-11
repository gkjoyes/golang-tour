# Inline Functions

One to the major objectives of using functions in a program is to save memory space, which becomes appreciable when a function is likely to be called many times. However, every time a function is called, it takes a lot of extra time in executing tasks such as jumping to the calling function. When a function is small, a substantial percentage of execution time may be spent in such overheads and sometimes maybe the time taken for jumping to the calling function will be greater than the time taken to execute that function.

One solution to this problem is to use macro definitions, commonly known as macros. But the major drawback with macros is that they are not really functions and therefore, the usual error checking process does not occur during compilation.

Modern programming languages have a different solution to this problem. To eliminate the time of calls to small functions, they propose a new function called inline function. An inline function is a function that is expanded in line when it is invoked thus saving time. The compiler replaces the function call with the corresponding function code that reduces the overhead of function call.

In Go, it's all about compiler decision. That means the compiler will decide whether this function need inlining or not. The compiler may not perform inlining in the following circumstances:

- If a function contains a loop.
- If a function is recursive.
- If a function contains static variables.
- If a function contains a switch command or goto statement.

## Advantage

- Function call overhead doesn't occur.
- It saves the overhead of a return call from a function.
- It saves the overhead of push/pop variables on the stack when the function is called.
- When we use the inline function it may enable the compiler to perform context-specific optimization on the function body.
- It increases the locality of reference by utilizing the instruction cache.
- An inline function may be useful for embedding systems because inline can yield less code than the function call preamble and return.

## Disadvantages

- Large Inline functions cause cache misses and affect efficiency negatively.
- Compilation overhead of copying the function body everywhere in the code at the time of compilation which is negligible in the case of small programs, but it may make a big difference in large codebases.
- If we require address of the function in a program, the compiler cannot perform inlining on such functions. Because for providing the address to a function, the compiler will have to allocate storage to it. But inline functions don't get storage, they are kept in the symbol table.
- Inline functions might cause thrashing because it might increase the size of the binary executable file. Thrashing in memory causes the performance of the computer to degrade and it also affects the efficiency of our code.
- The inline function may increase compile time overhead if someone tried to changes the code inside the inline function then all the calling location has to be recompiled again because the compiler would require to replace all the code once again to reflect the changes, otherwise it will continue with old functionality without any change.
