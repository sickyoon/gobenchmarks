package main

import (
	"bytes"
	"fmt"
	"testing"
)

func ConcatNormal() string {
	return hello + space + world
}

func ConcatBuffer() string {
	var buffer bytes.Buffer
	buffer.WriteString(hello)
	buffer.WriteString(space)
	buffer.WriteString(world)
	return buffer.String()
}

func ConcatCopy() string {
	bs := make([]byte, 11)
	copy(bs[0:], hello)
	copy(bs[5:], space)
	copy(bs[6:], world)
	return string(bs)
}

func ConcatPrint() string {
	return fmt.Sprintf("%s%s%s", hello, space, world)
}

func TestConcat(t *testing.T) {
	ans := "hello world"
	if ConcatNormal() != ans {
		t.Errorf("invalid ConcatNormal()")
	}
	if ConcatCopy() != ans {
		t.Errorf("invalid ConcatCopy()")
	}
	if ConcatPrint() != ans {
		t.Errorf("invalid ConcatPrint()")
	}
	if ConcatBuffer() != ans {
		t.Errorf("invalid ConcatBuffer()")
	}
}

func BenchmarkConcatStrNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatNormal()
	}
}

func BenchmarkConcatStrCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatCopy()
	}
}

func BenchmarkConcatStrBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatBuffer()
	}
}

func BenchmarkConcatStrPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatPrint()
	}
}
