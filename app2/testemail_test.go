package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("hello is not an email")
	}

	_, err = IsEmail("petch@mail.com")
	if err == nil {
		t.Error("petch@mail.com is an email")
	}

	_, err = IsEmail("petch@mail")
	if err != nil {
		t.Error("petch@mail is not an email")
	}
}
