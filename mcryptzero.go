package mcryptZero

import (
	"encoding/base64"
	"math/rand"
	"time"
)

func roR(b byte) byte {
	if (b % 2) == 0 {
		return b >> 1
	}

	return (b >> 1) + 128
}

func roL(b byte) byte {
	if b < 128 {
		return b << 1
	}

	return (b << 1) + 1
}

func encode8(d, k byte) byte {
	for i := 0; i < 8; i++ {
		if (k & 128) == 0 {
			d = roL(d) + 4
		} else {
			d = roL(d) - 3
		}
		k = k << 1
	}
	return d
}

func decode8(e, k byte) byte {
	for i := 0; i < 8; i++ {
		if (k & 1) == 0 {
			e = roR(e - 4)
		} else {
			e = roR(e + 3)
		}
		k = k >> 1
	}
	return e
}

func Encrypt(data, key []byte) []byte {
	result := make([]byte, len(data))
	dummy := byte(0)
	dummy2 := byte(0)
	dummy3 := byte(0)
	dummy4 := byte(0)
	for i, _ := range data {
		result[i] = data[i]
		result[i] = encode8(result[i], dummy)
		result[i] = encode8(result[i], dummy2)
		result[i] = encode8(result[i], dummy3)
		result[i] = encode8(result[i], dummy4)
		for j := 0; j < len(key); j++ {
			k := key[j]
			result[i] = encode8(result[i], k)
		}
		dummy4 = dummy3
		dummy3 = dummy2
		dummy2 = dummy
		dummy = result[i]
	}
	return result
}

func Decrypt(data, key []byte) []byte {
	result := make([]byte, len(data))
	dummy := byte(0)
	dummy2 := byte(0)
	dummy3 := byte(0)
	dummy4 := byte(0)
	ndummy := byte(0)
	for i, _ := range data {
		result[i] = data[i]
		ndummy = data[i]
		for j := len(key) - 1; j >= 0; j-- {
			k := key[j]
			result[i] = decode8(result[i], k)
		}
		result[i] = decode8(result[i], dummy4)
		result[i] = decode8(result[i], dummy3)
		result[i] = decode8(result[i], dummy2)
		result[i] = decode8(result[i], dummy)
		dummy4 = dummy3
		dummy3 = dummy2
		dummy2 = dummy
		dummy = ndummy
	}
	return result
}

func Crypt(data, key, salt []byte) string {
	encode1 := Encrypt(data, key)
	encode2 := Encrypt(data, salt)
	encode3 := Encrypt(encode1, encode2)
	return string(salt) + base64.StdEncoding.EncodeToString(encode3)
}

func SID(n int) string {
	salt := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	output := ""
	for i := 0; i < n; i++ {
		output = output + string(salt[rand.Intn(len(salt))])
	}
	return output
}

func init() {
	rand.Seed(time.Now().Unix())
}
