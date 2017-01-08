package main

import (
	"bytes"
	"fmt"
	"testing"
)

var (
	hello  string
	space  string
	world  string
	answer string
)

func init() {
	hello = "hello"
	space = " "
	world = "world"
	answer = "hello world"
}

// ConcatNormal -
func ConcatNormal() string {
	return hello + space + world
}

// ConcatBuffer -
func ConcatBuffer() string {
	var buffer bytes.Buffer
	buffer.WriteString(hello)
	buffer.WriteString(space)
	buffer.WriteString(world)
	return buffer.String()
}

// ConcatCopy -
func ConcatCopy() string {
	bs := make([]byte, 11)
	copy(bs[0:], hello)
	copy(bs[5:], space)
	copy(bs[6:], world)
	return string(bs)
}

// ConcatPrint -
func ConcatPrint() string {
	return fmt.Sprintf("%s%s%s", hello, space, world)
}

func TestConcat(t *testing.T) {
	if ConcatNormal() != answer {
		t.Errorf("invalid ConcatNormal()")
	}
	if ConcatCopy() != answer {
		t.Errorf("invalid ConcatCopy()")
	}
	if ConcatPrint() != answer {
		t.Errorf("invalid ConcatPrint()")
	}
	if ConcatBuffer() != answer {
		t.Errorf("invalid ConcatBuffer()")
	}
}

// BenchmarkNormal -
func BenchmarkConcatStrNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatNormal()
	}
}

// BenchmarkCopy -
func BenchmarkConcatStrCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatCopy()
	}
}

// BenchmarkBuffer -
func BenchmarkConcatStrBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatBuffer()
	}
}

// BenchmarkPrint -
func BenchmarkConcatStrPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatPrint()
	}
}
