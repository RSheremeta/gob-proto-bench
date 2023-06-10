package encode_decode

import (
	"bytes"
	"encoding/gob"

	"google.golang.org/protobuf/proto"
)

func EncodeGob(v any) []byte {
	var buf bytes.Buffer
	encdr := gob.NewEncoder(&buf)

	if err := encdr.Encode(v); err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func DecodeGob(b []byte, res any) {
	buf := bytes.NewBuffer(b)
	dcdr := gob.NewDecoder(buf)

	if err := dcdr.Decode(res); err != nil {
		panic(err)
	}
}

func EncodeProto(m proto.Message) []byte {
	res, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}

	return res
}

func DecodeProto(b []byte, res proto.Message) {
	if err := proto.Unmarshal(b, res); err != nil {
		panic(err)
	}
}
