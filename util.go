package raft_boltdb

import (
	"bytes"
	"encoding/binary"

	"github.com/ugorji/go/codec"
)

// Decode reverses the encode operation on a byte slice input
func decodeMsgPack(mh *codec.MsgpackHandle, buf []byte, out interface{}) error {
	r := bytes.NewBuffer(buf)
	dec := codec.NewDecoder(r, mh)
	return dec.Decode(out)
}

// Encode writes an encoded object to a new bytes buffer
func encodeMsgPack(mh *codec.MsgpackHandle, in interface{}) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(nil)
	enc := codec.NewEncoder(buf, mh)
	err := enc.Encode(in)
	return buf, err
}

// Converts bytes to an integer
func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

// Converts a uint to a byte slice
func uint64ToBytes(u uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, u)
	return buf
}
