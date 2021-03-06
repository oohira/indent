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

NOTE: Some operating system already has the other `indent` program that formats C source code.
You might have to add `export PATH=$GOPATH/bin:$PATH` to enable the newly installed one.

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
