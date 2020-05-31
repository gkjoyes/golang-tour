# Strings

- Go source code is always UTF-8.
- A string holds arbitrary bytes.
- A string literal, absent byte-level escapes, always hold valid UTF-8 sequences. Those sequences represent Unicode code points, called runes.
- No guarantee is made in Go that characters in strings are normalized.
- Multiple runes can represent different characters.
  - The lower grave-accented letter `à` is a character, and it's also a code point(U+00E0), but it has other representations.
  - We can use the `combining` grave accent code point, U+0300, and attach it to the lower case letter a, U+0061, to create the same character `à`.
  - In general, a character may be represented by a number of different sequences of code points(runes), and therefore different sequence of UTF-8 bytes.

## Format numbers and strings

- With Go `fmt` package, you can format numbers and strings padded with spaces and zeros, in different bases, and with optional quotes.
- You submit a `template string` that contains the text you want to format plus some `annotation verbs` that tells the `fmt` functions how to format the trailing arguments.
- Go `vet` command help to catch formatting errors like missing arguments.
- Arbitrary values can be encoded with backslash escapes and can be used in any string literal.
- There are four different formats:
  - `\x` followed by exactly two hexadecimal digits.
  - `\` followed by exactly three octal digits.
  - `\u` followed by exactly four hexadecimal digits.
  - `\U` followed by exactly eight hexadecimal digits.
- The escapes `\u` and `\U` represent Unicode code points.

### Default format types

|  Value        | Verb  | Format        |  Description                              |
|---------------|-------|---------------|-------------------------------------------|
| []int64{0, 1} |   %v  |  [0 1]        | Default format                            |
| []int64{0, 1} |   %#v | []int64{0, 1} | Go-syntax format                          |
| []int64       |   %T  | []int64       | The type of the value                     |

### Integer

|  Value        | Verb  | Format        |  Description                              |
|---------------|-------|---------------|-------------------------------------------|
| 15            |   %d  | 15            | Base 10                                   |
| 15            |  %+d  | +15           | Always show the sign                      |
| 15            |  %4d  | ␣␣15          | Pad with spaces(width 4, right justified) |
| 15            |  %-4d | 15␣␣          | Pad with spaces(width 4, left justified)  |
| 15            |  %04d | 0015          | Pad with zeros(width 4)                   |
| 15            |  %04d | 0015          | Pad with zeros(width 4)                   |
| 15            |  %b   | 1111          | Base 2                                    |
| 15            |  %o   | 17            | Base 8                                    |
| 15            |  %x   | f             | Base 16, lowercase                        |
| 15            |  %X   | F             | Base 16, uppercase                        |
| 15            |  %#x  | 0xf           | Base 16, with leading 0x                  |

### Characters

|  Value                | Verb  | Format        |  Description                              |
|-----------------------|-------|---------------|-------------------------------------------|
| 65(Unicode letter A)  |   %c  | A             | Character                                 |
| 65(Unicode letter A)  |   %q  | 'A'           | Quoted character                          |
| 65(Unicode letter A)  |   %U  | U+0041        | Unicode                                   |
| 65(Unicode letter A)  |   %#U | U+0041 'A'    | Unicode with character                    |

### Boolean(true/false)

|  Value                | Verb  | Format        |  Description                              |
|-----------------------|-------|---------------|-------------------------------------------|
| true/false            |   %t  | true          | Format a boolean as `true` or `false`     |

### Pointer (hex)

|  Value                | Verb  | Format        |  Description                                         |
|-----------------------|-------|---------------|------------------------------------------------------|
|  var c int            |   %p  | 0xc00002c008  | Format a pointer in base 16 notation with leading 0x |

## Float

|  Value                | Verb  | Format        |  Description                              |
|-----------------------|-------|---------------|-------------------------------------------|
| 123.456               | %e    | 1.234560e+02  | Scientific notation                       |
| 123.456               | %f    | 123.456000    | Decimal point, no exponent                |
| 123.456               | %.2f  | 123.45        | Default width, precision 2                |
| 123.456               | %g    | 123.456       | Exponent as needed, necessary digits only |

## String

|  Value                | Verb  | Format        |  Description                              |
|-----------------------|-------|---------------|-------------------------------------------|
| café                  | %s    | café          | Plain string                              |
| café                  | %6s   | ␣␣café        | Width 6, right justify                    |
| café                  | %-6s  | café␣␣        | Width 6, left justify                     |
| café                  | %q    | "café"        | Quoted string                             |
| café                  | %x    | 636166c3a9    | Hex dump of byte values                   |
| café                  | % x   | 63 61 66 c3 a9| Hex dump with spaces                      |

## Special Values

|  Value      | Description                                                      |
|-------------|------------------------------------------------------------------|
|  \a         | U+0007 alert or bell                                             |
|  \b         | U+0008 backspace (Equivalent of hitting the backspace key)       |
|  \\         | U+005c backslash                                                 |
|  \t         | U+0009 horizontal tab                                            |
|  \n         | U+000A line feed or newline                                      |
|  \f         | U+000C form feed(It skips to the start of the next page)         |
|  \r         | U+000D carriage return                                           |
|  \v         | U+000b vertical tab                                              |
