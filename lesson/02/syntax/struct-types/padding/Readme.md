# Structure Padding

- To align the data in memory, one or more empty bytes are inserted (or left empty) between addresses which are allocated for other structure members while memory allocation. This concept is called structure padding.

- The architecture of the computer processor is such a way that it can read 1 word(4 bytes in the 32-bit processor) from memory at a time. To make use of this advantage of processor, data are always aligned as 4 bytes package which leads to insert empty address between other member's address.

- Because of this structure padding, size of the structure is always not the same as what we think. For example, please consider below structure that has 5 members.

    ```go
        type student struct {
            id     int32
            name   string
            city   string
            passed bool
            phone  int32
        }
    ```

> 32-bit Architecture

- As per Go concepts, string data type occupies 2 words, int32 data type 4 bytes, and bool data type 1 byte. So, only 25 bytes(4 + 8 + 8 + 1 + 4) should be allocated for above structure.

- But this is wrong. Do you know why? In 32-bit architecture, we have 3 byte padding as shown in the image below. So the total size becomes 28 bytes.

![32-bit-architecture.png](https://github.com/gkjoyes/golang-tour/blob/master/lesson/02/syntax/struct-types/padding/images/32-bit-architecture.png)

> 64-bit Architecture

- When we get to the 64-bit architecture, it has 7-bit padding as shown below, and the total size changes to 48 bytes.

![64-bit-architecture.png](https://github.com/gkjoyes/golang-tour/blob/master/lesson/02/syntax/struct-types/padding/images/64-bit-architecture.png)

> Default

- If we did not specify the architecture at the time of the build, it would take the largest element size as the alignment boundary size.

- Consider the following structure, and see how padding works when ARCH is not specified during the build in the amd64 architecture, here we need 3 bytes of padding.

    ```go
        type student struct {
            id     int32
            passed bool
            phone  int32
        }
    ```

![Structure Padding](https://github.com/gkjoyes/golang-tour/blob/master/lesson/02/syntax/struct-types/padding/images/structure-padding.png)

## Examples

- [Structure padding in 32-bit architecture](https://github.com/gkjoyes/golang-tour/blob/master/lesson/02/syntax/struct-types/padding/example1/example1.go)
- [Structure padding in 64-bit architecture](https://github.com/gkjoyes/golang-tour/blob/master/lesson/02/syntax/struct-types/padding/example2/example2.go)
- [Structure padding without specifying ARCH](https://github.com/gkjoyes/golang-tour/blob/master/lesson/02/syntax/struct-types/padding/example2/example2.go)
