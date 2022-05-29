package util

import (
	"crypto/md5"
	"encoding/hex"
)

type Header struct {
	RequestID   string
	RequestName string
	Recipient   string
}

func (head *Header) SetID() {
	head.RequestID = GetMD5Hash(head.RequestName)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

