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
		t.Fatalf(`Hail(%v) = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

