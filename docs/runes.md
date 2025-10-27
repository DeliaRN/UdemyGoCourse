# Go Runes: Simple, Clear Explanation
## What is a Rune in Go?

**Rune = Unicode Code Point:**

In Go, a `rune` is just an alias for the type int32. It represents a single Unicode character.
_Example:_
The letter 'A', the emoji '😊', or the Chinese character '你'—each is a single rune.

## Why Not Just Use byte?
A byte is 8 bits and can only represent values from `0` to `255`. That’s fine for basic English letters (ASCII), but not enough for all the world’s characters and symbols.

A rune (int32) can represent over a million different characters, thanks to Unicode.


## Strings in Go:
Go strings are sequences of bytes, not runes.
If you have a string with non-ASCII characters, each character might be more than one byte.

_Example:_

```bash

s := "Go😊"
fmt.Println(len(s))      // Number of bytes
fmt.Println([]rune(s))   // Converts string to runes (Unicode code points)
fmt.Println(len([]rune(s))) // Number of runes (characters)
```

`"Go😊"` is 5 bytes (because '😊' is 4 bytes in UTF-8).
But it’s 3 runes: 'G', 'o', and '😊'.

### Why Care?
If you want to work with characters (not just bytes), use runes.
Useful for text processing, counting characters, or handling non-English text.

___

_Summary_

Rune = int32 = Unicode character

Use runes for characters, bytes for raw data.

Strings are bytes, but you can convert to runes for proper character handling.


## More Examples & Use Cases
- <h3>Spanish Characters Example</h3>


```bash
s := "Niño" // The 'ñ' is a special character
fmt.Println(len(s))         // Number of bytes
fmt.Println(len([]rune(s))) // Number of runes (characters)
```
`"Niño"` is 5 bytes (because 'ñ' is 2 bytes in UTF-8).

But it’s 4 runes: `'N'`, `'i'`, `'ñ'`, `'o'`.

___
- <h3>Accented Characters</h3>

```bash
s := "café" // The 'é' is accented
fmt.Println(len(s))         // Number of bytes
fmt.Println(len([]rune(s))) // Number of runes (characters)
```
`"café"` is 5 bytes (`'é'` is 2 bytes).
4 runes: `'c'`, `'a'`, `'f'`, `'é'`.


- <h3>Chinese Characters</h3>

```bash
s := "你好" // Chinese for "hello"
fmt.Println(len(s))         // Number of bytes
fmt.Println(len([]rune(s))) // Number of runes (characters)
```
`"你好"` is 6 bytes (each character is 3 bytes).
2 runes: `'你'`, `'好'`.

___

- <h3>Emojis</h3>

```bash
s := "👍🏽" // Thumbs up with skin tone
fmt.Println(len(s))         // Number of bytes
fmt.Println(len([]rune(s))) // Number of runes (characters)
```

`"👍🏽"` is 8 bytes (because emoji + modifier).

2 runes: `'👍'`, `'🏽'`. <- color (not processed by md)

___
.

Those are some common cases about runes

.
___


## Iterating Over Runes

If you want to process each character (not byte):

```bash
s := "café"
for i, r := range s {
    fmt.Printf("Character %d: %c (Unicode: %U)\n", i, r, r)
}
```
Output:
```
Character 0: c (Unicode: U+0063)
Character 1: a (Unicode: U+0061)
Character 2: f (Unicode: U+0066)
Character 3: é (Unicode: U+00E9)
```

___
## Use Case: Counting Characters

```bash
s := "¡Hola, señor!"
fmt.Println("Characters:", len([]rune(s)))
// Correctly counts all letters, including 'ñ' and '¡'
```

## Slicing Strings by Characters

If you want the first 3 characters, not bytes:
```bash
s := "mañana"
runes := []rune(s)
firstThree := string(runes[:3])
fmt.Println(firstThree) // "mañ"
```

Summary Table

| String   | Bytes | Runes | Notes                |
|----------|-------|-------|----------------------|
| "Go😊"   | 5     | 3     | Emoji is 4 bytes     |
| "Niño"   | 5     | 4     | 'ñ' is 2 bytes       |
| "café"   | 5     | 4     | 'é' is 2 bytes       |
| "你好"   | 6     | 2     | Each char is 3 bytes |
| "👍🏽"   | 8     | 2     | Emoji + modifier     |
