package greetings

import (
	"regexp"
	"testing"
)

//It calls greetings.Hail with a name, checking for a valid return value
func TestHailWithName(t *testing.T) {
	name := "donizetti"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hail(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hail("donizetti") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHailWithEmpty(t *testing.T) {
	name := ""
	msg, err := Hail(name)
	if msg != "" || err == nil {
		t.Fatalf(`Hail("") = %q, %v, want ""`, msg, err)
	}
}
