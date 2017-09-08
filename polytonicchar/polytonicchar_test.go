package polytonicchar

import "testing"

func TestLowerCaseSimple(t *testing.T) {
	cases := []struct {
		in   Char
		want string
	}{
		//Between Rho and Sigma is varSigma
		//Alpha, Rho, Sigma and Omega are interesting edge cases
		{Alpha, "\u03B1"},
		{Delta, "\u03B4"},
		{Rho, "\u03C1"},
		{Sigma, "\u03C3"},
		{Omega, "\u03C9"},
	}
	for _, c := range cases {
		got := New(c.in, false).String()
		if got != c.want {
			t.Errorf("New(%v, false).String() == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestUpperCaseSimple(t *testing.T) {
	cases := []struct {
		in   Char
		want string
	}{
		//Between Rho and Sigma is a gap
		//Alpha, Rho, Sigma and Omega are interesting edge cases
		{Alpha, "\u0391"},
		{Delta, "\u0394"},
		{Rho, "\u03A1"},
		{Sigma, "\u03A3"},
		{Omega, "\u03A9"},
	}
	for _, c := range cases {
		got := New(c.in, true).String()
		if got != c.want {
			t.Errorf("New(%v, true).String() == %v, want %v", c.in, got, c.want)
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
		{&PolytonicChar{name: Sigma, capital: false, variant: true}, "\u03C2"},
	}
	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("(%v).String() == %v, want %v", *c.in, got, c.want)
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
			t.Errorf("New(%v, false).String() == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestName(t *testing.T) {
	cases := []struct {
		in   *PolytonicChar
		want Char
	}{
		{
			&PolytonicChar{name: Alpha, capital: false, spiritus: Lenis},
			Alpha,
		},
		{
			&PolytonicChar{name: Eta, capital: false, iotaSubscriptum: true, accent: Circumflex, spiritus: Asper},
			Eta,
		},
	}
	for _, c := range cases {
		got := Name(c.in)
		if got != c.want {
			t.Errorf("(%v).String() == %v, want %v", *c.in, got, c.want)
		}
	}
}

func TestDiacritics(t *testing.T) {
	cases := []struct {
		in    *PolytonicChar
		want1 bool
		want2 Spiritus
		want3 Accent
	}{
		{
			&PolytonicChar{name: Alpha, capital: false, spiritus: Lenis},
			false, Lenis, None,
		},
		{
			&PolytonicChar{name: Eta, capital: false, iotaSubscriptum: true, accent: Circumflex, spiritus: Asper},
			true, Asper, Circumflex,
		},
	}
	for _, c := range cases {
		got1, got2, got3 := Diacritics(c.in)
		if got1 != c.want1 || got2 != c.want2 || got3 != c.want3 {
			t.Errorf("(%v).String() == %v, %v, %v, want %v, %v, %v", *c.in, got1, got2, got3, c.want1, c.want2, c.want3)
		}
	}
}

func TestSetVariant(t *testing.T) {
	cases := []struct {
		in   bool
		want bool
	}{
		{
			true,
			true,
		},
		{
			false,
			false,
		},
	}
	for _, c := range cases {
        char := New(Alpha, false)
        SetVariant(char, c.in)
		got := char.variant
		if got != c.want {
			t.Errorf("char := New(Alpha, false);SetVariant(char, %v);char.variant == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestSetIota(t *testing.T) {
	cases := []struct {
		in   bool
		want bool
	}{
		{
			true,
			true,
		},
		{
			false,
			false,
		},
	}
	for _, c := range cases {
        char := New(Alpha, false)
        SetIota(char, c.in)
		got := char.iotaSubscriptum
		if got != c.want {
			t.Errorf("char := New(Alpha, false);SetIota(char, %v);char.iotaSubscriptum == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestSetSpiritus(t *testing.T) {
	cases := []struct {
		in   Spiritus
		want Spiritus
	}{
		{
			Asper,
			Asper,
		},
		{
			None,
			None,
		},
	}
	for _, c := range cases {
        char := New(Alpha, false)
        SetSpiritus(char, c.in)
		got := char.spiritus
		if got != c.want {
			t.Errorf("char := New(Alpha, false);SetSpiritus(char, %v);char.spiritus == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestSetAccent(t *testing.T) {
	cases := []struct {
		in   Accent
		want Accent
	}{
		{
			Circumflex,
			Circumflex,
		},
		{
			Grave,
			Grave,
		},
	}
	for _, c := range cases {
        char := New(Alpha, false)
        SetAccent(char, c.in)
		got := char.accent
		if got != c.want {
			t.Errorf("char := New(Alpha, false);SetAccent(char, %v);char.accent == %v, want %v", c.in, got, c.want)
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
			&PolytonicChar{name: Omega, capital: false, spiritus: Lenis},
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
			&PolytonicChar{name: Omega, capital: false, spiritus: Lenis},
			true,
			false,
		},
	}
	for _, c := range cases {
		got := AreEqual(c.in1, c.in2, c.exact)
		if got != c.want {
			t.Errorf("HaveSameName(%v, %v, %v) == %v, want %v", c.in1, c.in2, c.exact, got, c.want)
		}
	}
}
