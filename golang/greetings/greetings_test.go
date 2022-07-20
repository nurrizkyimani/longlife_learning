package greetings

import (
	"fmt"
	"testing"
)

func TestHelloV2(t *testing.T) {
	name := "imani"
	//want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := HelloV2(name)

	fmt.Println(msg)

	if msg != name || err != nil {
		t.Fatalf(`Hello("imani") = %q, %v, want match for , nil`, msg, err)
	}
}

func TestHelloV2Empty(t *testing.T) {
	name := ""
	msg, err := HelloV2(name)
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") =%q, %v, want "", error`, msg, err)
	}
}
