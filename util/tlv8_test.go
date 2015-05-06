package util

import (
	"bytes"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTLV8SetByte(t *testing.T) {
	container := NewTLV8Container()
	container.SetByte(1, 0xF)
	assert.Equal(t, container.GetByte(1), byte(0xF))
}

func TestTLV8Bytes(t *testing.T) {
	data := "0102AFFA"
	b, _ := hex.DecodeString(data)
	buf := bytes.NewBuffer(b)
	container, err := NewTLV8ContainerFromReader(buf)
	assert.Nil(t, err)
	assert.Equal(t, container.GetBytes(1), []byte{0xAF, 0xFA})
}

func TestTLV8BytesFromMultipleSource(t *testing.T) {
	data := "0102AFFA0103BFFBAA"
	b, _ := hex.DecodeString(data)

	buf := bytes.NewBuffer(b)
	container, err := NewTLV8ContainerFromReader(buf)
	assert.Nil(t, err)
	assert.Equal(t, container.GetBytes(1), []byte{0xAF, 0xFA, 0xBF, 0xFB, 0xAA})
}

func TestTLV8SetMoreThanMaxBytes(t *testing.T) {
	container := NewTLV8Container()
	data := "00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF" // 384 bytes
	b, _ := hex.DecodeString(data)
	assert.Equal(t, len(b), 384)

	container.SetBytes(1, b)

	// split up in 255 chunks
	// 01(type)FF(length=255)bytes...01(type)81(length=129)bytes...
	expectedData := "01FF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEE0181FF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF" // 384 bytes
	expectedBytes, _ := hex.DecodeString(expectedData)
	assert.Equal(t, container.BytesBuffer().Bytes(), expectedBytes)
}

func TestTLV8SetBytes(t *testing.T) {
	container := NewTLV8Container()
	container.SetBytes(1, []byte{0xAF, 0xFA})
	assert.Equal(t, container.GetBytes(1), []byte{0xAF, 0xFA})
}

func TestTLV8BytesBuffer(t *testing.T) {
	container := NewTLV8Container()
	container.SetBytes(1, []byte{0xAF, 0xFA})

	assert.Equal(t, container.BytesBuffer().Bytes(), []byte{0x01, 0x02, 0xAF, 0xFA})
}

func TestTLV8String(t *testing.T) {
	container := NewTLV8Container()
	container.SetString(1, "Hello World")

	assert.Equal(t, container.GetString(1), "Hello World")
}