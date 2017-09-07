package polytonicchar

import "testing"

func TestLowerCaseSimple(t *testing.T) {
	cases := []struct {
		in   Char
		want string
	}{
		{Alpha, "\u03B1"},
		{Delta, "\u03B4"},
		{Rho, "\u03C1"},
		{Sigma, "\u03C3"},
		{Omega, "\u03C9"},
	}
	for _, c := range cases {
		got := New(c.in, false).String()
		if got != c.want {
			t.Errorf("New(%q, false).String() == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestUpperCaseSimple(t *testing.T) {
	cases := []struct {
		in   Char
		want string
	}{
		{Alpha, "\u0391"},
		{Delta, "\u0394"},
		{Rho, "\u03A1"},
		{Sigma, "\u03A3"},
		{Omega, "\u03A9"},
	}
	for _, c := range cases {
		got := New(c.in, true).String()
		if got != c.want {
			t.Errorf("New(%q, true).String() == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGreekExtended(t *testing.T) {
	cases := []struct {
		in   *PolytonicChar
		want string
	}{
		{&PolytonicChar{name: Alpha, capital: false, spiritus: Lenis}, "\u1F00"},
		{&PolytonicChar{name: Eta, capital: false, iotaSubscriptum: true, accent: Circumflex, spiritus: Asper}, "\u1F97"},
		{&PolytonicChar{name: Rho, capital: true, spiritus: Asper}, "\u1FEC"},
		{&PolytonicChar{name: Omega, capital: false, accent: Circumflex, iotaSubscriptum: true}, "\u1FF7"},
		{&PolytonicChar{name: Omega, capital: true, iotaSubscriptum: true}, "\u1FFC"},
		{&PolytonicChar{name: Eta, capital: false, iotaSubscriptum: true, accent: Grave}, "\u1FC2"},
		{&PolytonicChar{name: Eta, capital: false, iotaSubscriptum: true, accent: Acute}, "\u1FC4"},
		{&PolytonicChar{name: Omicron, capital: false, accent: Grave}, "\u1F78"},
		{&PolytonicChar{name: Omicron, capital: true, accent: Acute}, "\u1FF9"},
	}
	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("(%q).String() == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestInterpunction(t *testing.T) {
	cases := []struct {
		in   Char
		want string
	}{
		{Comma, ","},
		{Dot, "."},
		{Apostrophe, "'"},
		{Interpunct, "\u0387"},
	}
	for _, c := range cases {
		got := New(c.in, false).String()
		if got != c.want {
			t.Errorf("New(%q, false).String() == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestAreEqual(t *testing.T) {
	cases := []struct {
		in1   *PolytonicChar
		in2   *PolytonicChar
		exact bool
		want  bool
	}{
		{
			&PolytonicChar{name: Alpha, capital: false, spiritus: Lenis},
			&PolytonicChar{name: Alpha, capital: true, spiritus: Asper, iotaSubscriptum: true},
			false,
			true,
		},
		{
			&PolytonicChar{name: Alpha, capital: false, spiritus: Lenis},
			&PolytonicChar{name: Gamma, capital: false, spiritus: Lenis},
			false,
			false,
		},
		{
			&PolytonicChar{name: Alpha, capital: false, spiritus: Lenis},
			&PolytonicChar{name: Alpha, capital: true, spiritus: Asper, iotaSubscriptum: true},
			true,
			false,
		},
		{
			&PolytonicChar{name: Alpha, capital: false, spiritus: Lenis},
			&PolytonicChar{name: Alpha, capital: false, spiritus: Lenis},
			true,
			true,
		},
		{
			&PolytonicChar{name: Alpha, capital: false, spiritus: Lenis},
			&PolytonicChar{name: Gamma, capital: false, spiritus: Lenis},
			true,
			false,
		},
	}
	for _, c := range cases {
		got := AreEqual(c.in1, c.in2, c.exact)
		if got != c.want {
			t.Errorf("HaveSameName(%q, %q, %t) == %t, want %t", c.in1, c.in2, c.exact, got, c.want)
		}
	}
}
