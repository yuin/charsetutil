package charsetutil

import (
	"bytes"
	"strings"
	"testing"
)

func TestEncodeOk(t *testing.T) {
	expected := []byte{'\x82', '\xb1', '\x82', '\xf1', '\x82', '\xc9', '\x82', '\xbf', '\x82', '\xed'}
	assert := func(b []byte, err error) {
		if err != nil {
			t.Errorf("Failed: %s", err.Error())
		}
		if string(b) != string(expected) {
			t.Error("Failed")
		}
	}

	b, err := EncodeString("こんにちわ", "Windows-31J")
	assert(b, err)

	b, err = EncodeBytes([]byte("こんにちわ"), "Windows-31J")
	assert(b, err)

	b, err = Encode("こんにちわ", "Windows-31J")
	assert(b, err)

	b, err = EncodeReader(strings.NewReader("こんにちわ"), "Windows-31J")
	assert(b, err)

	b = MustEncodeString("こんにちわ", "Windows-31J")
	assert(b, nil)

	b = MustEncodeBytes([]byte("こんにちわ"), "Windows-31J")
	assert(b, nil)

	b = MustEncode("こんにちわ", "Windows-31J")
	assert(b, nil)

	b = MustEncodeReader(strings.NewReader("こんにちわ"), "Windows-31J")
	assert(b, nil)
}

func TestEncodeError(t *testing.T) {
	assert := func(b []byte, err error) {
		if b != nil || err == nil {
			t.Error("Failed")
		}
	}

	assertPanic := func(f func() []byte) {
		defer func() {
			if recover() == nil {
				t.Error("Should be failed")
			}
		}()
		b := f()
		if b != nil {
			t.Error("Failed")
		}
	}

	b, err := EncodeString("こんにちわ", "unknown")
	assert(b, err)

	b, err = EncodeBytes([]byte("こんにちわ"), "unknown")
	assert(b, err)

	b, err = Encode("こんにちわ", "unknown")
	assert(b, err)

	b, err = EncodeReader(strings.NewReader("こんにちわ"), "unknown")
	assert(b, err)

	assertPanic(func() []byte { return MustEncodeString("こんにちわ", "unknown") })

	assertPanic(func() []byte { return MustEncodeBytes([]byte("こんにちわ"), "unknown") })

	assertPanic(func() []byte { return MustEncode("こんにちわ", "unknown") })

	assertPanic(func() []byte { return MustEncodeReader(strings.NewReader("こんにちわ"), "unknown") })
}

func TestDecodeOk(t *testing.T) {
	source := []byte{'\x82', '\xb1', '\x82', '\xf1', '\x82', '\xc9', '\x82', '\xbf', '\x82', '\xed'}
	expected := "こんにちわ"

	assert := func(b string, err error) {
		if err != nil {
			t.Errorf("Failed: %s", err.Error())
		}
		if b != expected {
			t.Error("Failed")
		}
	}

	b, err := DecodeString(string(source), "Windows-31J")
	assert(b, err)

	b, err = DecodeBytes(source, "Windows-31J")
	assert(b, err)

	b, err = Decode(source, "Windows-31J")
	assert(b, err)

	b, err = DecodeReader(bytes.NewReader(source), "Windows-31J")
	assert(b, err)

	b = MustDecodeString(string(source), "Windows-31J")
	assert(b, nil)

	b = MustDecodeBytes(source, "Windows-31J")
	assert(b, nil)

	b = MustDecode(source, "Windows-31J")
	assert(b, nil)

	b = MustDecodeReader(bytes.NewReader(source), "Windows-31J")
	assert(b, nil)
}

func TestDecodeError(t *testing.T) {
	source := []byte{'\x82', '\xb1', '\x82', '\xf1', '\x82', '\xc9', '\x82', '\xbf', '\x82', '\xed'}
	assert := func(s string, err error) {
		if s != "" || err == nil {
			t.Error("Failed")
		}
	}

	assertPanic := func(f func() string) {
		defer func() {
			if recover() == nil {
				t.Error("Should be failed")
			}
		}()
		s := f()
		if s != "" {
			t.Error("Failed")
		}
	}

	b, err := DecodeString(string(source), "unknown")
	assert(b, err)

	b, err = DecodeBytes(source, "unknown")
	assert(b, err)

	b, err = Decode(source, "unknown")
	assert(b, err)

	b, err = DecodeReader(bytes.NewReader(source), "unknown")
	assert(b, err)

	assertPanic(func() string { return MustDecodeString(string(source), "unknown") })

	assertPanic(func() string { return MustDecodeBytes(source, "unknown") })

	assertPanic(func() string { return MustDecode(source, "unknown") })

	assertPanic(func() string { return MustDecodeReader(bytes.NewReader(source), "unknown") })
}

func TestGuess(t *testing.T) {
	sourceEuc := []byte{'\xa4', '\xa2', '\xa4', '\xa4', '\xa4', '\xa6', '\xa4', '\xa8', '\xa4', '\xaa', '\x0d', '\x0a', '\xa5', '\xbd', '\xc7', '\xbd', '\x0d', '\x0a', '\x74', '\x65', '\x73', '\x74', '\x0d', '\x0a', '\x8e', '\xb6', '\x8e', '\xb7', '\x8e', '\xb8', '\x8e', '\xb9', '\x8e', '\xba'}
	// sourceSjis := []byte{'\x82', '\xa0', '\x82', '\xa2', '\x82', '\xa4', '\x82', '\xa6', '\x82', '\xa8', '\x0d', '\x0a', '\x83', '\x5c', '\x94', '\x5c', '\x0d', '\x0a', '\x74', '\x65', '\x73', '\x74', '\x0d', '\x0a', '\xb6', '\xb7', '\xb8', '\xb9', '\xba'}

	assert := func(r CharsetGuess, charset, language string, err error) {
		if err != nil {
			t.Errorf("Failed:%+v", err)
		}
		if r.Charset() != charset {
			t.Errorf("'%s' expected, but got '%s'", charset, r.Charset())
		}
		if r.Language() != language {
			t.Errorf("'%s' expected, but got '%s'", language, r.Language())
		}
	}

	result, err := Guess(sourceEuc)
	assert(result, "EUC-JP", "ja", err)

	result, err = GuessBytes(sourceEuc)
	assert(result, "EUC-JP", "ja", err)

	result, err = GuessReader(bytes.NewReader(sourceEuc))
	assert(result, "EUC-JP", "ja", err)

	result, err = GuessString("ああｲｲ”haa")
	assert(result, "UTF-8", "", err)

}
