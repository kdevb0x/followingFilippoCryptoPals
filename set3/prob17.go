package set3

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"log"
	mathrand "math/rand"
	"strings"
)

func newCBCPaddingOracles(plaintexts [][]byte) (
	generateMessage func() []byte,
	checkMessagePadding func(message []byte) bool,
) {
	key := make([]byte, 16)
	rand.Read(key)
	b, _ := aes.NewCipher(key)
	generateMessage = func() []byte {
		msg := plaintexts[mathrand.Intn(len(plaintexts))]

		iv := make([]byte, 16)
		rand.Read(iv)

		ct := encryptCBC(padPKCS7(msg, 16), b, iv)
		return append(iv, ct...)

	}
	checkMessagePadding = func(message []byte) bool {
		iv, msg := message[:16], message[16:]
		res := unpadPKCS7(decryptCBC(msg, b, iv))
		return res != nil
	}
	return
}
func attackCBCPaddingOracle(ct []byte, checkMessagePadding func(ct []byte) bool) []byte {
	ct = ct[len(ct)-32:]
	log.Println(checkMessagePadding(ct))
	return nil
}
