package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	s = "hello"
	t.Log(len(s))
	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s))
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共合国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}
