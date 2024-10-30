package ordscanner_test

import (
	"bufio"
	"fmt"
	"strings"

	ords "github.com/pschou/go-ordscanner"
)

func ExampleNew() {
	s := ords.New(
		bufio.NewScanner(strings.NewReader("a\nc\nd\n")),
		bufio.NewScanner(strings.NewReader("b\ne\nf\ng\n")),
		bufio.NewScanner(strings.NewReader("f\nz\n")),
	)

	for s.Scan() {
		fmt.Println(s.Text())
	}
	// Output:
	// a
	// b
	// c
	// d
	// e
	// f
	// f
	// g
	// z
}

func ExampleNewWithoutScanFirst() {
	list := []*bufio.Scanner{
		bufio.NewScanner(strings.NewReader("b\ne\nf\ng\n")),
		bufio.NewScanner(strings.NewReader("f\nz\n")),
		bufio.NewScanner(strings.NewReader("a\nc\nd\n")),
	}
	for _, scn := range list {
		scn.Scan()
	}

	s := ords.New(list...)
	fmt.Println(s.Text())
	// Output:
	// a
}
