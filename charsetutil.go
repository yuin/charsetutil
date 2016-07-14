package charsetutil

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"strings"
)

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func DecodeReader(s io.Reader, enc string) (string, error) {
	reader, err := charset.NewReaderLabel(enc, s)
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func MustDecodeReader(s io.Reader, enc string) string {
	ret, err := DecodeReader(s, enc)
	panicIfError(err)
	return ret
}

func DecodeBytes(s []byte, enc string) (string, error) {
	return DecodeReader(bytes.NewReader(s), enc)
}

func MustDecodeBytes(s []byte, enc string) string {
	ret, err := DecodeReader(bytes.NewReader(s), enc)
	panicIfError(err)
	return ret
}

func DecodeString(s, enc string) (string, error) {
	return DecodeReader(strings.NewReader(s), enc)
}

func MustDecodeString(s, enc string) string {
	ret, err := DecodeReader(strings.NewReader(s), enc)
	panicIfError(err)
	return ret
}

func Decode(s []byte, enc string) (string, error) {
	return DecodeReader(bytes.NewReader(s), enc)
}

func MustDecode(s []byte, enc string) string {
	ret, err := DecodeReader(bytes.NewReader(s), enc)
	panicIfError(err)
	return ret
}

func EncodeReader(s io.Reader, enc string) ([]byte, error) {
	e, _ := charset.Lookup(enc)
	if e == nil {
		return nil, errors.New(fmt.Sprintf("unsupported charset: %q", enc))
	}
	var buf bytes.Buffer
	writer := transform.NewWriter(&buf, e.NewEncoder())
	_, err := io.Copy(writer, s)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func MustEncodeReader(s io.Reader, enc string) []byte {
	ret, err := EncodeReader(s, enc)
	panicIfError(err)
	return ret
}

func EncodeBytes(s []byte, enc string) ([]byte, error) {
	return EncodeReader(bytes.NewReader(s), enc)
}

func MustEncodeBytes(s []byte, enc string) []byte {
	ret, err := EncodeReader(bytes.NewReader(s), enc)
	panicIfError(err)
	return ret
}

func EncodeString(s, enc string) ([]byte, error) {
	return EncodeReader(strings.NewReader(s), enc)
}

func MustEncodeString(s, enc string) []byte {
	ret, err := EncodeReader(strings.NewReader(s), enc)
	panicIfError(err)
	return ret
}

func Encode(s string, enc string) ([]byte, error) {
	return EncodeReader(strings.NewReader(s), enc)
}

func MustEncode(s string, enc string) []byte {
	ret, err := EncodeReader(strings.NewReader(s), enc)
	panicIfError(err)
	return ret

}
