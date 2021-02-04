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


https://viewer.diagrams.net/?highlight=0000ff&edit=_blank&layers=1&nav=1#RzZbbjpswEIafhstWgAkhl002baVtpaqRerj04llw1zCR42ygT19TBogDm%2B5WVZrc4Pl9mvlmYttjq6J6p%2Fk2%2F4gClBf6ovLYjReGgc8S%2B2mUulVmSdQKmZaCBg3CRv6Ebiapeylg5ww0iMrIrSumWJaQGkfjWuPBHXaPyt11yzMYCZuUq7H6VQqTt2oSzgf9Pcgs73YO4kXbU%2FBuMEWyy7nAw5HE1h5baUTTtopqBaqB13Fp5719ord3TENpnjOBrb9s5rfFsvy8KJKHWx2xWLyiVR652lPA5KypOwIa96WAZhHfY8tDLg1stjxteg8251bLTaGsFdgmLQfaQPWkn0EfvS0bwAKMru2QbkJMwKhiIp%2FswxF%2FkvIj9BFpnDKe9SsPUGyDuLyAUXh9jPpgzzAKwktCiq4PEntGIV0W0uz6IEXJnyFFl2QUjxj5za%2B7D%2F4jqll4ZQdTMkICwl5eZJZY2s%2FSpYTa5JhhydUHxC2x%2BQHG1HT18r1BlxxU0nxrpr%2BekfWdFmvaN9WxUZPROtZ4c561dR73OoUzQc6nc6JBcSMf3fWnCNPUTyjtzn0u45MDtD8ruiUM1xkYmnWSp96Nv0%2FdfLrK%2FWCU0ktXeexfWZUvRqiYNe9qA7%2FjEkKW2Yiajd%2B4aHZG4wOsUKEe%2Fhz3UqkTiSuZldZMLS%2Bw%2BrKhKe2D8A11FFKIZpvJXLjZ%2BheHTjJzC5VNXGLBRD7Cl%2BfDmsNztK304VHP1r8A
