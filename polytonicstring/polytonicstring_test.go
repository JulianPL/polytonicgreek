package polytonicstring

import (
	"polytonicgreek/polytonicchar"
	"testing"
)

func TestNew(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{
			"Au)ta\\r o( e)k lime/nos",
			"\u0391\u1f50\u03c4\u1f70\u03c1 \u1f41 \u1f10\u03ba \u03bb\u03b9\u03bc\u1f73\u03bd\u03bf\u03c2",
		},
		{
			"o(",
			"\u1f41",
		},
	}
	for _, c := range cases {
		got := New(c.in).String()
		if got != c.want {
			t.Errorf("New(%q).String() == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestBetaToStr(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{
			"Au)ta\\r o( e)k lime/nos",
			"\u0391\u1f50\u03c4\u1f70\u03c1 \u1f41 \u1f10\u03ba \u03bb\u03b9\u03bc\u1f73\u03bd\u03bf\u03c2",
		},
	}
	for _, c := range cases {
		got := BetaToStr(c.in).String()
		if got != c.want {
			t.Errorf("BetaToStr(%q).String() == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestInString(t *testing.T) {
	cases := []struct {
		inChar *polytonicchar.PolytonicChar
		inStr  string
		exact  bool
		want   int
	}{
		{
			polytonicchar.New(polytonicchar.Tau, true),
			"Au)ta\\r o( e)k lime/nos",
			false,
			2,
		},
		{
			polytonicchar.New(polytonicchar.Tau, true),
			"Au)ta\\r o( e)k lime/nos",
			true,
			-1,
		},
	}
	for _, c := range cases {
		got := InString(c.inChar, New(c.inStr), c.exact)
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
		{
			"o( e(",
			"Au)ta\\r o( e)k lime/nos",
			false,
			6,
		},
		{
			"o( e(",
			"Au)ta\\r o( e)k lime/nos",
			true,
			-1,
		},
		{
			"o( e)",
			"Au)ta\\r o( e)k lime/nos",
			true,
			6,
		},
	}
	for _, c := range cases {
		got := SubString(New(c.inSub), New(c.inStr), c.exact)
		if got != c.want {
			t.Errorf("InString(New(%q), New(%q), %t) == %d, want %d", c.inSub, c.inStr, c.exact, got, c.want)
		}
	}
}
