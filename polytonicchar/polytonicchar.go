//This package provides the basic struct for polytonic greek letters
package polytonicchar

type PolytonicChar struct {
	name            Char
	capital         bool
	iotaSubscriptum bool
	variant         bool
	spiritus        Spiritus
	accent          Accent
}

const None = 0

type Char int

const (
	Alpha Char = iota
	Beta
	Gamma
	Delta
	Epsilon
	Zeta
	Eta
	Theta
	Iota
	Kappa
	Lambda
	Mu
	Nu
	Xi
	Omicron
	Pi
	Rho
	_
	Sigma
	Tau
	Upsilon
	Phi
	Chi
	Psi
	Omega
	Space
	Comma
	Dot
	Apostrophe
	Interpunct
	Question
)

type Spiritus int

const (
	_ Spiritus = iota
	Lenis
	Asper
)

type Accent int

const (
	_ Accent = iota
	Grave
	Acute
	Circumflex
)


func New(n Char, c bool) *PolytonicChar {
	return &PolytonicChar{name: n, capital: c}
}

func Name(c *PolytonicChar) Char {
	return c.name
}

func Diacritics(c *PolytonicChar) (bool, Spiritus, Accent) {
	return c.iotaSubscriptum, c.spiritus, c.accent
}

func SetVariant(char *PolytonicChar, in bool) {
	char.variant = in
}

func SetIota(char *PolytonicChar, in bool) {
	char.iotaSubscriptum = in
}

func SetSpiritus(char *PolytonicChar, in Spiritus) {
	char.spiritus = in
}

func SetAccent(char *PolytonicChar, in Accent) {
	char.accent = in
}

func AreEqual(char1 *PolytonicChar, char2 *PolytonicChar, exact bool) bool {
	if exact {
		return *char1 == *char2
	}
	return char1.name == char2.name
}
