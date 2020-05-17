# indent

A tiny program just to indent lines.

## Usage

```
indent [-n int] [file]

  -n int
        number of spaces (default 4)
  file
        read from a file instead of standard input (only 1 file is supported)
```

## Installation

```
$ go get github.com/oohira/indent
```

## Example

Indent standard input.

```shell
$ echo "Hello\nWorld" | indent
    Hello
    World
```

Indent a file with 8 space characters.

```shell
$ echo "Hello\nWorld" > hello.txt
$ indent -n 8 hello.txt
        Hello
        World
```

Indent Clipboard text (macOS only).

```shell
$ echo "Hello\nWorld" | pbcopy
$ pbpaste | indent
    Hello
    World
```
