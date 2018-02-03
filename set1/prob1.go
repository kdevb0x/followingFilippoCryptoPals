package set1

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func hex2base64(hs string) (string, error) {
	v, err := hex.DecodeString(hs)
	if err != nil {
		return "", err
	}
	log.Println("%s", err)
	return base64.StdEncoding.EncodeToString(v), nil

}
func xor(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("xor: not equal lengths")
	}
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}
	return res
}
