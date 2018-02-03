package set3

import (
	"testing"
	//"github.com/kidoda/followingFillipoLivecoding/set3"
)

func TestProblem17(t *testing.T) {
	generateMessage, checkMessagePadding := newCBCPaddingOracles([][]byte{
		decodeBase64(t, "MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc="),
		decodeBase64(t, "MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic="),
		decodeBase64(t, "MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw=="),
		decodeBase64(t, "MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg=="),
		decodeBase64(t, "MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl"),
		decodeBase64(t, "MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA=="),
		decodeBase64(t, "MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw=="),
		decodeBase64(t, "MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8="),
		decodeBase64(t, "MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g="),
		decodeBase64(t, "MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93"),
	})
	attackCBCPaddingOracle(generateMessage, checkMessagePadding)
}
