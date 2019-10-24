// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package set1

import (
	"bytes"
	"encoding/hex"
	"io/ioutil"
	"testing"
)

func TestProb1(t *testing.T) {
	res, err := hex2base64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if err != nil {
		t.Fatal(err)
	}
	if res != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Error("wrong string", res)
	}
}
func TestProb2(t *testing.T) {
	res := xor(hexDecode(t, "1c0111001f010100061a024b53535009181c"),
		hexDecode(t, "686974207468652062756c6c277320657965"))
	if !bytes.Equal(res, hexDecode(t, "746865206b696420646f6e277420706c6179")) {
		t.Errorf("bad result %x", res)
	}
}

func hexDecode(t *testing.T, s string) []byte {
	v, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal("failed to decode hex", s)
	}
	return v
}
func corpusFromFile(t *testing.T, name string) map[rune]float64 {
	text, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatal("failed to open corpus file:", err)
	}

	return buildCorpus(string(text))
}

func TestProblem3(t *testing.T) {
	c := corpusFromFile(t, "_testdata/aliceinwonderland.txt")
	for char, val := range c {
		t.Logf("%c: %.5f", char, val)
	}
	res := findSingleXORKey(hexDecode(t, "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"), c)
	t.Logf("%s", res)
}

func TestRepeatingKeyXOR(t *testing.T) {
	phrase := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

	cipher := `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`

	res := hex.EncodeToString(repeatingKeyXOR([]byte{'I', 'C', 'E'}, []byte(phrase)))
	if res != cipher {
		t.Logf("%s doesn't match expected cipher", res)
		t.Fail()
	}
}

func TestHammingDistance(t *testing.T) {
	if res := hammingDistance([]byte("this is a test"), []byte("wokka wokka!!!")); res != 37 {
		t.Fatal("wrong Hamming distance:", res)
	}
}
