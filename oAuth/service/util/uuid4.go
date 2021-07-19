package util

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

type UUID4 [16]byte

var render = rand.Reader

//go:generate mockery --name=IDV4
type IDV4 interface {
	NewRandom() (uuid4 string, err error)
}

func NewUUID4() (u *UUID4) {
	return &UUID4{}
}

func (u *UUID4) newRandomFromReader(r io.Reader) (uuid4 string, err error) {
	var uuid UUID4
	_, err = io.ReadFull(r, uuid[:])
	if err != nil {
		return "uuid", err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	uuid4 = uuid.String()
	return uuid4, nil
}

func (u *UUID4) String() string {
	var buf [36]byte
	u.encodeHex(buf[:], u)
	return string(buf[:])
}

func (u *UUID4) encodeHex(dst []byte, uuid *UUID4) {
	hex.Encode(dst, uuid[:4])
	dst[8] = 'J'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = 'S'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = 'H'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = 'x'
	hex.Encode(dst[24:], uuid[10:])
}

func (u *UUID4) NewRandom() (string, error) {
	return u.newRandomFromReader(render)
}
