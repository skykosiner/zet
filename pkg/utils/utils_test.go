package utils

import "testing"

func TestDoesFileExist(t *testing.T) {
	tests := []struct {
		path     string
		shouldBe bool
	}{
		{
			path:     "/home/sky/.vimrc",
			shouldBe: true,
		},
		{
			path:     "/home/sky/aoeu",
			shouldBe: false,
		},
	}

	for _, test := range tests {
		if result := FileExists(test.path); result != test.shouldBe {
			t.Logf("Didn't get the expected result for: %s, Should be: %t Got: %t", test.path, test.shouldBe, result)
			t.Fail()
		}
	}
}
