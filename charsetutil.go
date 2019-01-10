package charsetutil

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/gogs/chardet"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

// CharsetGuess is a guessd charcter set
type CharsetGuess interface {
	// Charset returns a guessed charcter set
	Charset() string

	// Language returns a guessed language
	Language() string

	// Confidence returns a confidence of this guess
	Confidence() int
}

type charsetGuess struct {
	*chardet.Result
}

func (g *charsetGuess) Charset() string {
	return g.Result.Charset
}

func (g *charsetGuess) Language() string {
	return g.Result.Language
}

func (g *charsetGuess) Confidence() int {
	return g.Result.Confidence
}

// GuessBytes guesses a character set of given bytes
func GuessBytes(s []byte) (CharsetGuess, error) {
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(s)
	if err != nil {
		return nil, err
	}
	return &charsetGuess{result}, err
}

// Guess guesses a character set of given bytes
func Guess(s []byte) (CharsetGuess, error) {
	return GuessBytes(s)
}

// GuessBytes guesses a character set of given Reader
func GuessReader(s io.Reader) (CharsetGuess, error) {
	detector := chardet.NewTextDetector()
	buf := make([]byte, 128)
	if _, err := s.Read(buf); err != nil {
		return nil, err
	}
	result, err := detector.DetectBest(buf)
	if err != nil {
		return nil, err
	}
	return &charsetGuess{result}, err
}

// GuessBytes guesses a character set of given string
func GuessString(s string) (CharsetGuess, error) {
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest([]byte(s))
	if err != nil {
		return nil, err
	}
	return &charsetGuess{result}, err
}

// DecodeReader converts given Reader to a UTF-8 string
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

// MustDecodeReader converts given Reader to a UTF-8 string and panics if errros occur.
func MustDecodeReader(s io.Reader, enc string) string {
	ret, err := DecodeReader(s, enc)
	panicIfError(err)
	return ret
}

// DecodeBytes converts given bytes to a UTF-8 string
func DecodeBytes(s []byte, enc string) (string, error) {
	return DecodeReader(bytes.NewReader(s), enc)
}

// MustDecodeBytes converts given bytes to a UTF-8 string and panics if errros occur.
func MustDecodeBytes(s []byte, enc string) string {
	ret, err := DecodeReader(bytes.NewReader(s), enc)
	panicIfError(err)
	return ret
}

// DecodeString converts given string to a UTF-8 string
func DecodeString(s, enc string) (string, error) {
	return DecodeReader(strings.NewReader(s), enc)
}

// MustDecodeString converts given string to a UTF-8 string and panics if errros occur.
func MustDecodeString(s, enc string) string {
	ret, err := DecodeReader(strings.NewReader(s), enc)
	panicIfError(err)
	return ret
}

// DecodeBytes converts given bytes to a UTF-8 string
func Decode(s []byte, enc string) (string, error) {
	return DecodeReader(bytes.NewReader(s), enc)
}

// MustDecodeBytes converts given bytes to a UTF-8 string and panics if errros occur.
func MustDecode(s []byte, enc string) string {
	ret, err := DecodeReader(bytes.NewReader(s), enc)
	panicIfError(err)
	return ret
}

// EncodeReader converts a Reader to bytes encoded with given encoding
func EncodeReader(s io.Reader, enc string) ([]byte, error) {
	e, _ := charset.Lookup(enc)
	if e == nil {
		return nil, fmt.Errorf("unsupported charset: %q", enc)
	}
	var buf bytes.Buffer
	writer := transform.NewWriter(&buf, e.NewEncoder())
	_, err := io.Copy(writer, s)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MustEncodeReader converts a Reader to bytes encoded with given encoding and panics if errors occur
func MustEncodeReader(s io.Reader, enc string) []byte {
	ret, err := EncodeReader(s, enc)
	panicIfError(err)
	return ret
}

// EncodeBytes converts bytes to bytes encoded with given encoding
func EncodeBytes(s []byte, enc string) ([]byte, error) {
	return EncodeReader(bytes.NewReader(s), enc)
}

// MustEncodeBytes converts a bytes to bytes encoded with given encoding and panics if errors occur
func MustEncodeBytes(s []byte, enc string) []byte {
	ret, err := EncodeReader(bytes.NewReader(s), enc)
	panicIfError(err)
	return ret
}

// EncodeString converts a string to bytes encoded with given encoding
func EncodeString(s, enc string) ([]byte, error) {
	return EncodeReader(strings.NewReader(s), enc)
}

// MustEncodeString converts a bytes to bytes encoded with given encoding and panics if errors occur
func MustEncodeString(s, enc string) []byte {
	ret, err := EncodeReader(strings.NewReader(s), enc)
	panicIfError(err)
	return ret
}

// Encode converts a string to bytes encoded with given encoding
func Encode(s string, enc string) ([]byte, error) {
	return EncodeReader(strings.NewReader(s), enc)
}

// MustEncode converts a bytes to bytes encoded with given encoding and panics if errors occur
func MustEncode(s string, enc string) []byte {
	ret, err := EncodeReader(strings.NewReader(s), enc)
	panicIfError(err)
	return ret

}
