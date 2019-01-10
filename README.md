## charsetutil - An easiest way to convert character set encodings in Go

charsetutil provides easiest way to convert character set encodings in Go.

## Install

```bash
go get github.com/yuin/charsetutil
```

## Utilities

- `Decode*` : Converts from the specified charset to UTF-8.
- `Encode*` : Converts from the UTF-8 to specified charset.
- `Guess*` : Guesses a charcter set.

- `MustDecode*` : Same as `Decode*`, but panics when errors occur
- `MustEncode*` : Same as `Encode*`, but panics when errors occur

```go
b, err = EncodeString("こんにちわ", "Windows-31J")
b, err = Encode("こんにちわ", "Windows-31J")
b, err = EncodeBytes([]byte("こんにちわ"), "Windows-31J")
b, err = EncodeReader(strings.NewReader("こんにちわ"), "Windows-31J")
b = MustEncodeString("こんにちわ", "Windows-31J")
b = MustEncode("こんにちわ", "Windows-31J")
b = MustEncodeBytes([]byte("こんにちわ"), "Windows-31J")
b = MustEncodeReader(strings.NewReader("こんにちわ"), "Windows-31J")

s, err = DecodeString(string(source), "Windows-31J")
s, err = Decode(source, "Windows-31J")
s, err = DecodeBytes(source, "Windows-31J")
s, err = DecodeReader(bytes.NewReader(source), "Windows-31J")
s = MustDecodeString(string(source), "Windows-31J")
s = MustDecode(source, "Windows-31J")
s = MustDecodeBytes(source, "Windows-31J")
s = MustDecodeReader(bytes.NewReader(source), "Windows-31J")

cs, err := GuessString(string(source))
cs, err := GuessBytes(source)
cs, err := GuessReader(bytes.NewReader(source))
cs, err := Guess(source)
```

## Supported character sets

See [Encoding spec on WHATWG](https://encoding.spec.whatwg.org/#names-and-labels)

## Author

Yusuke Inuzuka

## License

[BSD License](http://opensource.org/licenses/BSD-2-Clause)

