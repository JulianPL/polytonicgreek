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

func TestVariants(t *testing.T) {
	cases := []struct {
		in   *PolytonicChar
		want string
	}{
		{&PolytonicChar{name: Space, capital: false, variant: true}, "\n"},
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
		{Space, " "},
		{Comma, ","},
		{Dot, "."},
		{Apostrophe, "'"},
		{Interpunct, "\u0387"},
		{Question, ";"},
	}
	for _, c := range cases {
		got := New(c.in, false).String()
		if got != c.want {
			t.Errorf("New(%v, false).String() == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestGreekExtended(t *testing.T) {
	cases := []struct {
		in   *PolytonicChar
		want string
	}{
		{&PolytonicChar{name: Rho, capital: false, iotaSubscriptum: false, spiritus: Asper, accent: None}, "\u1FE5"},
		{&PolytonicChar{name: Rho, capital: true, iotaSubscriptum: false, spiritus: Asper, accent: None}, "\u1FEC"},
		{&PolytonicChar{name: Eta, capital: false, iotaSubscriptum: true, spiritus: Asper, accent: Circumflex}, "\u1F97"},
		{&PolytonicChar{name: Omega, capital: true, iotaSubscriptum: true, spiritus: Lenis, accent: None}, "\u1FA8"},
		{&PolytonicChar{name: Alpha, capital: false, iotaSubscriptum: false, spiritus: Lenis, accent: None}, "\u1F00"},
		{&PolytonicChar{name: Omega, capital: false, iotaSubscriptum: true, spiritus: None, accent: Circumflex}, "\u1FF7"},
		{&PolytonicChar{name: Eta, capital: false, iotaSubscriptum: true, spiritus: None, accent: Grave}, "\u1FC2"},
		{&PolytonicChar{name: Eta, capital: false, iotaSubscriptum: true, spiritus: None, accent: Acute}, "\u1FC4"},
		{&PolytonicChar{name: Alpha, capital: true, iotaSubscriptum: true, spiritus: None, accent: None}, "\u1FBC"},
		{&PolytonicChar{name: Upsilon, capital: false, iotaSubscriptum: false, spiritus: None, accent: Circumflex}, "\u1FE6"},
		{&PolytonicChar{name: Iota, capital: true, iotaSubscriptum: false, spiritus: None, accent: Acute}, "\u1FDB"},
		{&PolytonicChar{name: Omicron, capital: true, iotaSubscriptum: false, spiritus: None, accent: Grave}, "\u1FF8"},
		{&PolytonicChar{name: Epsilon, capital: false, iotaSubscriptum: false, spiritus: None, accent: Acute}, "\u1F73"},
		{&PolytonicChar{name: Eta, capital: false, iotaSubscriptum: false, spiritus: None, accent: Grave}, "\u1F74"},
		{&PolytonicChar{name: Omicron, capital: false, accent: Grave}, "\u1F78"},
		{&PolytonicChar{name: Omicron, capital: true, accent: Acute}, "\u1FF9"},
	}
	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("(%v).String() == %v, want %v", *c.in, got, c.want)
		}
	}
}

func TestInvalidChars(t *testing.T) {
	const replacementChar = string(65533)
	cases := []struct {
		in *PolytonicChar
	}{
		{&PolytonicChar{name: Sigma, capital: true, variant: true}},
		{&PolytonicChar{name: Space + 1000, capital: false}},
		{&PolytonicChar{name: Rho, capital: true, spiritus: Lenis}},
		{&PolytonicChar{name: Rho, capital: false, spiritus: Lenis, iotaSubscriptum: true}},
		{&PolytonicChar{name: Rho, capital: false, spiritus: Lenis, accent: Grave}},
		{&PolytonicChar{name: Iota, capital: false, iotaSubscriptum: true, spiritus: Lenis, accent: Grave}},
		{&PolytonicChar{name: Gamma, capital: false, spiritus: Asper}},
		{&PolytonicChar{name: Eta, capital: true, iotaSubscriptum: true, spiritus: None, accent: Grave}},
		{&PolytonicChar{name: Iota, capital: false, iotaSubscriptum: true}},
		{&PolytonicChar{name: Upsilon, capital: true, iotaSubscriptum: false, spiritus: None, accent: Circumflex}},
		{&PolytonicChar{name: Epsilon, capital: false, iotaSubscriptum: false, spiritus: None, accent: Circumflex}},
		{&PolytonicChar{name: Delta, capital: true, accent: Acute}},
		{&PolytonicChar{name: Gamma, capital: false, accent: Acute}},
	}
	for _, c := range cases {
		got := c.in.String()
		if got != replacementChar {
			t.Errorf("(%v).String() == %v, want %v", *c.in, got, replacementChar)
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
		got := SetVariant(New(Alpha, false), c.in).variant
		if got != c.want {
			t.Errorf("SetVariant(New(Alpha, false), %v).variant == %v, want %v", c.in, got, c.want)
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
		got := SetIota(New(Alpha, false), c.in).iotaSubscriptum
		if got != c.want {
			t.Errorf("SetIota(New(Alpha, false), %v).iotaSubscriptum == %v, want %v", c.in, got, c.want)
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
		got := SetSpiritus(New(Alpha, false), c.in).spiritus
		if got != c.want {
			t.Errorf("SetSpiritus(New(Alpha, false), %v).spiritus == %v, want %v", c.in, got, c.want)
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
		got := SetAccent(New(Alpha, false), c.in).accent
		if got != c.want {
			t.Errorf("SetAccent(New(Alpha, false), %v).accent == %v, want %v", c.in, got, c.want)
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
