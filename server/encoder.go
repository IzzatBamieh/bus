package main

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

type Encoder struct {
}

func (encoder *Encoder) encode(writer io.Writer, value interface{}) error {
	return jsoniter.NewEncoder(writer).Encode(value)
}

func (encoder *Encoder) decode(reader io.Reader, value interface{}) error {
	return jsoniter.NewDecoder(reader).Decode(value)
}
