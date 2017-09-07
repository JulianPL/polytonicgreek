package polytonicchar

type PolytonicChar struct {
	name            Char
	capital         bool
	iotaSubscriptum bool
	variant         bool
	spiritus        Spiritus
	accent          Accent
}

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

func (c *PolytonicChar) String() string {
	if c.variant && c.name == Space {
		return "\n"
	}
	if c.name > Omega {
		letterOrder := map[Char]string{
			Space:      " ",
			Comma:      ",",
			Dot:        ".",
			Apostrophe: "'",
			Interpunct: "\u0387",
			Question:   ";",
		}
		if val, ok := letterOrder[c.name]; ok {
			return val
		}
		return string(replacementChar)
	}
	if c.variant && c.name == Sigma && !c.capital {
		return "\u03C2"
	}
	if c.spiritus != None {
		s := greekExtended
		if c.spiritus == Asper {
			s += asperOffset
		}
		if c.capital {
			s += capitalOffset
		}
		s += accentOffset * int(c.accent)
		if c.iotaSubscriptum {
			s += iotaOffset
			letterOrder := map[Char]int{
				Alpha: 0,
				Eta:   1,
				Omega: 2,
			}
			if val, ok := letterOrder[c.name]; ok {
				s += val * letterOffset
				return string(s)
			}
			return string(replacementChar)
		}
		letterOrder := map[Char]int{
			Alpha:   0,
			Epsilon: 1,
			Eta:     2,
			Iota:    3,
			Omicron: 4,
			Upsilon: 5,
			Omega:   6,
		}
		if val, ok := letterOrder[c.name]; ok {
			s += val * letterOffset
			return string(s)
		}
		if c.name == Rho {
			s += rhoOffset
			if c.capital {
				s -= 1 //it works: https://en.wikipedia.org/wiki/Greek_diacritics
				if c.spiritus == Lenis {
					return string(replacementChar) //not possible
				}
			}
			return string(s)
		}
		return string(replacementChar)
	}
	if c.accent == Circumflex {
		s := circumflexWithoutSpiritus
		if c.iotaSubscriptum {
			s += iotaWithCircumflexWithoutSpiritusOffset
		}
		letterOrder := map[Char]int{
			Alpha:   0,
			Eta:     1,
			Iota:    2,
			Upsilon: 3,
			Omega:   4,
		}
		if val, ok := letterOrder[c.name]; ok {
			s += val * letterOffset
			return string(s)
		}
		return string(replacementChar)
	}
	if c.iotaSubscriptum {
		s := iotaSubscriptumWithoutSpiritus
		letterOrder := map[Char]int{
			Alpha: 0,
			Eta:   1,
			Omega: 4,
		}
		if val, ok := letterOrder[c.name]; ok {
			s += val * letterOffset
		} else {
			return string(replacementChar)
		}
		if c.capital {
			s += 10
			return string(s)
		}
		accentOrder := map[Accent]int{
			Grave: 0,
			None:  1,
			Acute: 2,
		}
		if val, ok := accentOrder[c.accent]; ok {
			s += val
			return string(s)
		}
		return string(replacementChar)
	}
	if c.accent != None {
		if c.capital {
			s := accentendUpperCaseWithoutSpiritus
			letterOrder := map[Char]int{
				Alpha:   0,
				Epsilon: 1,
				Eta:     1,
				Iota:    2,
				Upsilon: 3,
				Omicron: 4,
				Omega:   4,
			}
			if val, ok := letterOrder[c.name]; ok {
				s += val * letterOffset
			} else {
				return string(replacementChar)
			}
			if c.name == Epsilon || c.name == Omicron {
				s -= 2
			}
			accentOrder := map[Accent]int{
				Grave: 0,
				Acute: 1,
			}
			if val, ok := accentOrder[c.accent]; ok {
				s += val
				return string(s)
			}
			return string(replacementChar)
		}
		s := accentendLowerCaseWithoutSpiritus
		letterOrder := map[Char]int{
			Alpha:   0,
			Epsilon: 1,
			Eta:     2,
			Iota:    3,
			Omicron: 4,
			Upsilon: 5,
			Omega:   6,
		}
		if val, ok := letterOrder[c.name]; ok {
			s += val * 2
		} else {
			return string(replacementChar)
		}
		accentOrder := map[Accent]int{
			Grave: 0,
			Acute: 1,
		}
		if val, ok := accentOrder[c.accent]; ok {
			s += val
			return string(s)
		}
		return string(replacementChar)
	}
	if c.capital {
		return string(upperCaseAlpha + c.name)
	}
	return string(lowerCaseAlpha + c.name)
}
