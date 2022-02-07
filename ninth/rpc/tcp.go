package rpc

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

// TODO(Protocol, Header)
const lenBytes = 8

func ReadMsg(conn net.Conn) (bs []byte, err error) {
	msgLenBytes := make([]byte, lenBytes)
	length, err := conn.Read(msgLenBytes)
	defer func() {
		if msg := recover(); msg != nil {
			err = errors.New(fmt.Sprintf("%v", msg))
		}
	}()
	if err != nil {
		return nil, err
	}

	if length != lenBytes {
		return nil, errors.New("could not read the length data")
	}

	dataLen := binary.BigEndian.Uint64(msgLenBytes)
	bs = make([]byte, dataLen)
	_, err = io.ReadFull(conn, bs)
	return bs, err
}

func EncodeProto(m interface{}) ([]byte, error) {
	respData, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal response failed")
		return nil, err
	}
	return EncodeMsg(respData), nil
}

func EncodeMsg(msg []byte) []byte {
	encode := make([]byte, lenBytes+len(msg))
	binary.BigEndian.PutUint64(encode[:lenBytes], uint64(len(msg)))
	copy(encode[8:], msg)
	return encode
}
