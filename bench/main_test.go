package main

import (
	"testing"
	"crypto/sha512"
	"encoding/hex"
	"bytes"
	"strings"
	"fmt"
	"github.com/mattn/go-pipeline"
)

func CalcPassHash1(account, password string) string {
	salt := sha512.Sum512([]byte(account))
	str := password + ":" + hex.EncodeToString(salt[:])
	hash := sha512.Sum512([]byte(str))
	return hex.EncodeToString(hash[:])
}

func CalcPassHash2(account, password string) string {
	buf := bytes.NewBuffer([]byte{})
	salt := sha512.Sum512([]byte(account))
	salt_hex := make([]byte, hex.EncodedLen(len(salt)))
	hex.Encode(salt_hex, salt[:])
	buf.WriteString(password)
	buf.WriteString(":")
	buf.Write(salt_hex)
	hash := sha512.Sum512(buf.Bytes())
	return hex.EncodeToString(hash[:])
}

func CalcPassHash3(account, password string) string {
	salt := sha512.Sum512([]byte(account))

	hash := sha512.Sum512(
		[]byte(strings.Join([]string{password, hex.EncodeToString(salt[:])}, ":")),
	)
	return hex.EncodeToString(hash[:])
}

var joinChar = []byte{':'}

func CalcPassHash4(account, password string) string {
	salt := sha512.Sum512([]byte(account))

	saltHash := make([]byte, hex.EncodedLen(len(salt)))
	hex.Encode(saltHash, salt[:])
	hash := sha512.Sum512(
		bytes.Join(
			[][]byte{
				[]byte(password),
				saltHash,
			},
			joinChar,
		),
	)
	return hex.EncodeToString(hash[:])
}

func CalcPassHashOpenSSL(account, password string) string {
	salt, err := pipeline.Output(
		[]string{"printf", "%s", account},
		[]string{"openssl", "dgst", "-sha512"},
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	src := password + ":" + string(salt)
	hash, err := pipeline.Output(
		[]string{"printf", "%s", src},
		[]string{"openssl", "dgst", "-sha512"},
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(hash)
}

func TestCalcPassHashes(t *testing.T) {
	password, account := "jawoefiao23", "suzuki-kyosuke"
	t.Log(CalcPassHashOpenSSL(account, password))
	t.Log(CalcPassHash1(account, password))
	t.Log(CalcPassHash2(account, password))
	t.Log(CalcPassHash3(account, password))
	t.Log(CalcPassHash4(account, password))
}

func BenchmarkCalcPassHash(b *testing.B) {
	password, account := "jawoefiao23", "suzuki-kyosuke"
	for i := 0; i < b.N; i++ {
		CalcPassHashOpenSSL(account, password)
	}
}

func BenchmarkCalcPassHash1(b *testing.B) {
	password, account := "jawoefiao23", "suzuki-kyosuke"
	for i := 0; i < b.N; i++ {
		CalcPassHash1(account, password)
	}
}

func BenchmarkCalcPassHash2(b *testing.B) {
	password, account := "jawoefiao23", "suzuki-kyosuke"
	for i := 0; i < b.N; i++ {
		CalcPassHash2(account, password)
	}
}

func BenchmarkCalcPassHash3(b *testing.B) {
	password, account := "jawoefiao23", "suzuki-kyosuke"
	for i := 0; i < b.N; i++ {
		CalcPassHash3(account, password)
	}
}

func BenchmarkCalcPassHash4(b *testing.B) {
	password, account := "jawoefiao23", "suzuki-kyosuke"
	for i := 0; i < b.N; i++ {
		CalcPassHash4(account, password)
	}
}