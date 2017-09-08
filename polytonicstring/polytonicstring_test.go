package polytonicstring

import (
	//"polytonicgreek/polytonicchar"
	"testing"
)

func TestBetaToStr(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{
			"",
			"",
		},
		{
			"W)h|u=iA)eo(",
			"\u1f68\u1fc3\u1fe6\u03b9\u1f08\u03b5\u1f41",
		},
		{
			". :\n,';",
			". \u0387\n,';",
		},
		{
			"sys",
			"\u03c3\u03c8\u03c2",
		},
	}
	for _, c := range cases {
		got := BetaToStr(c.in).String()
		if got != c.want {
			t.Errorf("BetaToStr(%q).String() == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestNew(t *testing.T) {
	cases := []struct {
		in string
	}{
		{
			"Au)ta\\r o( e)k lime/nos",
		},
	}
	for _, c := range cases {
		got := New(c.in).String()
		want := BetaToStr(c.in).String()
		if got != want {
			t.Errorf("New(%q).String() == %q, want %q", c.in, got, want)
		}
	}
}

func TestInString(t *testing.T) {
	cases := []struct {
		inChar string
		inStr  string
		exact  bool
		want   int
	}{
		{"I=", "Au)ta\\r o( e)k lime/nos", false, 12},
		{"g", "Au)ta\\r o( e)k lime/nos", false, -1},
		{"A", "Au)ta\\r o( e)k lime/nos", true, 0},
		{"e)", "Au)ta\\r o( e)k lime/nos", true, 8},
		{"A(", "Au)ta\\r o( e)k lime/nos", true, -1},
		{"a", "Au)ta\\r o( e)k lime/nos", true, -1},
	}
	for _, c := range cases {
		got := InString(New(c.inChar)[0], New(c.inStr), c.exact)
		if got != c.want {
			t.Errorf("InString(%q, New(%q), %t) == %d, want %d", c.inChar, c.inStr, c.exact, got, c.want)
		}
	}
}

func TestSubString(t *testing.T) {
	cases := []struct {
		inSub string
		inStr string
		exact bool
		want  int
	}{
		{"o( e(", "Au)ta\\r o( e)k lime/nos", false, 6},
		{"o( e)", "Au)ta\\r o( e)k lime/nos", false, 6},
		{"o( e(", "Au)ta\\r o( e)k lime/nos", true, -1},
		{"o( e)", "Au)ta\\r o( e)k lime/nos", true, 6},
	}
	for _, c := range cases {
		got := SubString(New(c.inSub), New(c.inStr), c.exact)
		if got != c.want {
			t.Errorf("InString(New(%q), New(%q), %t) == %d, want %d", c.inSub, c.inStr, c.exact, got, c.want)
		}
	}
}

func TestWrapString(t *testing.T) {
	cases := []struct {
		in    string
		width int
		want  string
	}{
		{
			"oo o o o o o o o",
			5,
			"\u03bf\u03bf \u03bf\n\u03bf \u03bf \u03bf\n\u03bf \u03bf \u03bf",
		},
		{
			"oo o o o o o o o",
			4,
			"\u03bf\u03bf \u03bf\n\u03bf \u03bf\n\u03bf \u03bf\n\u03bf \u03bf",
		},
		{
			"oo\no o o o o o o",
			5,
			"\u03bf\u03bf\n\u03bf \u03bf \u03bf\n\u03bf \u03bf \u03bf\n\u03bf",
		},
		{
			"oooooo oooo ooo oo o",
			4,
			"\u03bf\u03bf\u03bf\u03bf\u03bf\u03bf\n\u03bf\u03bf\u03bf\u03bf\n\u03bf\u03bf\u03bf\n\u03bf\u03bf \u03bf",
		},
	}
	for _, c := range cases {
		got := WrapString(New(c.in), c.width).String()
		if got != c.want {
			t.Errorf("WrapString(New(%q), %v).String() == %q, want %q", c.in, c.width, got, c.want)
		}
	}
}
