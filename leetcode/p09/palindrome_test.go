package p09

import "testing"

func TestName(t *testing.T) {

	if true != isPalindrome(121) {
		t.Error()
	}

	if false != isPalindrome(-121) {
		t.Error()
	}

	if false != isPalindrome(10) {
		t.Error()
	}
}
