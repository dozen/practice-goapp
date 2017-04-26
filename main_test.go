package main

import (
	"testing"
	"github.com/go-redis/redis"
	"github.com/dozen/practice-goapp/ruby-marshal"
	"bytes"
)

func TestRubyUnMarshaller(t *testing.T) {
	Key := "rack:session:c84c037568458e2fa16d246d96482311ab06b82f3fc5679fb97375a71072a228"
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB: 1,
	})

	b, e := c.Get(Key).Bytes()
	if e != nil {
		t.Error(e.Error())
	}
	ruby_marshal.NewDecoder(bytes.NewReader(b))
}