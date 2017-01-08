package main

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

func PrintNormal() {
	fmt.Printf("%s%s%s%d\n", hello, space, world, answer)
}

func BenchmarkPrintStrNormal(b *testing.B) {
	fmt.Println()
	for i := 0; i < b.N; i++ {
		PrintNormal()
	}
}

func PrintLine() {
	fmt.Println(hello, space, world, answer)
}

func BenchmarkPrintStrLine(b *testing.B) {
	fmt.Println()
	for i := 0; i < b.N; i++ {
		PrintLine()
	}
}

func PrintLog() {
	log.Printf("%s%s%s%d", hello, space, world, answer)
}

func BenchmarkPrintStrLog(b *testing.B) {
	fmt.Println()
	for i := 0; i < b.N; i++ {
		PrintLog()
	}
}

func PrintConcat() {
	fmt.Printf(hello + space + world + strconv.Itoa(answer) + "\n")
}

func BenchmarkPrintStrConcat(b *testing.B) {
	fmt.Println()
	for i := 0; i < b.N; i++ {
		PrintConcat()
	}
}

func PrintLineConcat() {
	fmt.Println(hello + space + world + strconv.Itoa(answer))
}

func BenchmarkPrintStrLineConcat(b *testing.B) {
	fmt.Println()
	for i := 0; i < b.N; i++ {
		PrintLineConcat()
	}
}

func PrintLogConcat() {
	log.Printf(hello + space + world + strconv.Itoa(answer))
}

func BenchmarkPrintStrLogConcat(b *testing.B) {
	fmt.Println()
	for i := 0; i < b.N; i++ {
		PrintLogConcat()
	}
}
