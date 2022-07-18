package token

import (
	"bytes"
	"math/rand"
)

func GenToken(l int) string {
	s := "0123456789qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM"
	tkn := bytes.Buffer{}
	for i := 0; i < l; i++ {
		idx := rand.Int() % len(s)
		tkn.WriteByte(s[idx])
	}
	return tkn.String()
}
