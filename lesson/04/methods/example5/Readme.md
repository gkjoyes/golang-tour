# Use semantics in a consistent way

This documentation trying to show the importance of using value or pointer semantics in a consistent way. Choose the semantic that is reasonable and practical for the given type and be consistent. One exception is an unmarshal operation since that always requires the address of a value.

> Let's look at the semantics used in some golang standard libraries:

## net

These are named types from the net package called [IP](https://github.com/golang/go/blob/2bed279721d684de828d0027db43a9d6283938a1/src/net/ip.go#L32) and [IPMask](https://github.com/golang/go/blob/2bed279721d684de828d0027db43a9d6283938a1/src/net/ip.go#L38) with a base type that is a slice of bytes. Since we use value semantics for reference types, the implementation is using value semantics for both.

```go
    type IP []byte

    type IPMask []byte
```

> These are some example which uses value semantics in the net package:

- [Mask](https://github.com/golang/go/blob/2bed279721d684de828d0027db43a9d6283938a1/src/net/ip.go#L248) is using a value receiver and returning a value of type IP.

    ```go
        func (ip IP) Mask(mask IPMask) IP {
            //code.
        }
    ```

- [isEmptyString](https://github.com/golang/go/blob/2bed279721d684de828d0027db43a9d6283938a1/src/net/ip.go#L371) accepts a value of type IP and returns a value of type string.

    ```go
        func ipEmptyString(ip IP) string {
            // code.
        }
    ```

## time

Should time use value or pointer semantics? If you need to modify a time value should you mutate the value or create a new one?

```go
    type Time struct{
        sec  int64
        nsec int32
        loc  *Location
    }
```

The factory function dictates the semantics that will be used. The [Now](https://github.com/golang/go/blob/b71eafbcece175db33acfb205e9090ca99a8f984/src/time/time.go#L1066) function returns a value of type `Time`. This means we should be using value semantics and copy Time values.

> These are some example in time library which follow value semantics:

- [Add](https://github.com/golang/go/blob/b71eafbcece175db33acfb205e9090ca99a8f984/src/time/time.go#L813) is using a value receiver and returning a value of type `Time`.

    ```go
        func (t Time) Add(d Duration) Time {
            // code.
        }
    ```

- [div](https://github.com/golang/go/blob/b71eafbcece175db33acfb205e9090ca99a8f984/src/time/time.go#L1435) accepts a value of type `Time` and returns values of built-in types.

    ```go
        func div(t Time, d Duration) (qmod2 int, r Duration) {
            // code.
        }
    ```

- The only one pointer semantics for the `Time` api are these unmarshal related functions.

    ```go
        func (t *Time) UnmarshalBinary(data []byte) error {
            // code.
        }
    ```

    ```go
        func (t *Time) GobDecode(data []byte) error {
            // code.
        }
    ```

    ```go
        func (t *Time) UnmarshalJSON(data []byte) error {
            // code.
        }
    ```

    ```go
        func (t *Time) UnmarshalText(data []byte) error {
            // code.
        }
    ```

## os

Factory function dictates the semantics that will be used. The [Open](https://github.com/golang/go/blob/8194187a2de9b693af6656ca956762437c2f9c64/src/os/file.go#L306) function returns a pointer of type [File](https://github.com/golang/go/blob/a2a3dd00c934fa15ad880ee5fe1f64308cbc73a7/src/os/types.go#L16). This means we should be using pointer semantics and share File values.

```go
    func Open(name string) (file *File, err error) {
        return OpenFile(name, O_RDONLY, 0)
    }
```

> These are some examples which follow pointer semantics:

- [Chdir](https://github.com/golang/go/blob/664d2707276b3523895d98f1233c3ad7b7297220/src/os/file_plan9.go#L498) is using a pointer receiver.

    ```go
        func (f *File) Chdir() error {
            // code.
        }
    ```

- [epipecheck](https://github.com/golang/go/blob/e64675a79fef5924f268425de021372df874010e/src/os/file_unix.go#L177) accepts a pointer of type `File`.

    ```go
        func epipecheck(file *File, e error) {
            // code.
        }
    ```
