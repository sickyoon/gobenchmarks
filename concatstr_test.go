package main

import (
	"bytes"
	"fmt"
	"testing"
)

func concatNormal() string {
	return hello + space + world
}

func BenchmarkConcatStrNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatNormal()
	}
}

func concatConstantNormal() string {
	return chello + cspace + cworld
}

func BenchmarkConcatStrConstantNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatConstantNormal()
	}
}

func concatBuffer() string {
	var buffer bytes.Buffer
	buffer.WriteString(hello)
	buffer.WriteString(space)
	buffer.WriteString(world)
	return buffer.String()
}

func BenchmarkConcatStrBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatBuffer()
	}
}

func concatConstantBuffer() string {
	var buffer bytes.Buffer
	buffer.WriteString(chello)
	buffer.WriteString(cspace)
	buffer.WriteString(cworld)
	return buffer.String()
}

func BenchmarkConcatStrConstantBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatConstantBuffer()
	}
}

func concatCopy() string {
	bs := make([]byte, 11)
	copy(bs[0:], hello)
	copy(bs[5:], space)
	copy(bs[6:], world)
	return string(bs)
}

func BenchmarkConcatStrCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatCopy()
	}
}

func concatConstantCopy() string {
	bs := make([]byte, 11)
	copy(bs[0:], chello)
	copy(bs[5:], cspace)
	copy(bs[6:], cworld)
	return string(bs)
}

func BenchmarkConcatStrConstantCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatConstantCopy()
	}
}

func concatPrint() string {
	return fmt.Sprintf("%s%s%s", hello, space, world)
}

func BenchmarkConcatStrPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatPrint()
	}
}

func concatConstantPrint() string {
	return fmt.Sprintf("%s%s%s", chello, cspace, cworld)
}

func BenchmarkConcatStrConstantPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatConstantPrint()
	}
}

func TestConcat(t *testing.T) {
	ans := "hello world"
	if concatNormal() != ans {
		t.Errorf("invalid concatNormal()")
	}
	if concatConstantNormal() != ans {
		t.Errorf("invalid concatConstantNormal()")
	}
	if concatCopy() != ans {
		t.Errorf("invalid concatCopy()")
	}
	if concatConstantCopy() != ans {
		t.Errorf("invalid concatConstantCopy()")
	}
	if concatPrint() != ans {
		t.Errorf("invalid concatPrint()")
	}
	if concatConstantPrint() != ans {
		t.Errorf("invalid concatConstantPrint()")
	}
	if concatBuffer() != ans {
		t.Errorf("invalid concatBuffer()")
	}
	if concatConstantBuffer() != ans {
		t.Errorf("invalid concatConstantBuffer()")
	}
}
