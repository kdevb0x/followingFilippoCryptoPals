// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package set1

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"unicode/utf8"
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
func buildCorpus(text string) map[rune]float64 {
	c := make(map[rune]float64)
	for _, char := range text {
		c[char]++
	}
	total := utf8.RuneCountInString(text)
	for char := range c {
		c[char] = c[char] / float64(total)
	}
	return c
}
