package polytonicstring

import (
	"polytonicgreek/polytonicchar"
	"unicode"
)

type PolytonicString []*polytonicchar.PolytonicChar

func New(str string) PolytonicString {
	polyStr := PolytonicString(make([]*polytonicchar.PolytonicChar, 0))
	polyStr = append(polyStr, BetaToStr(str)...)
	return polyStr
}

func InString(char *polytonicchar.PolytonicChar, str PolytonicString, exact bool) int {
	for i, comp := range str {
		if polytonicchar.AreEqual(char, comp, exact) {
			return i
		}
	}
	return -1
}

func SubString(sub PolytonicString, str PolytonicString, exact bool) int {
	for i, _ := range str {
		if i+len(sub) > len(str) {
			break
		}
		found := true
		for j, _ := range sub {
			if !polytonicchar.AreEqual((sub)[j], (str)[i+j], exact) {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}

func WrapString(str PolytonicString, width int) *PolytonicString{
	space := polytonicchar.New(polytonicchar.Space, false)
	newline := polytonicchar.New(polytonicchar.Space, false)
	polytonicchar.SetVariant(newline, true)
	current := 0
	lastFound := -1
	for i, comp := range str {
		current += 1
		if polytonicchar.AreEqual(space, comp, true) {
			lastFound = i
		}
		if polytonicchar.AreEqual(newline, comp, true) {
			current = 0
			lastFound = -1
		}
		if current > width && lastFound >= 0{
			polytonicchar.SetVariant((str)[lastFound], true)
			current = i - lastFound
			lastFound = -1
		}
	}
    return &str
}

func BetaToStr(str string) PolytonicString {
	polyStr := PolytonicString(make([]*polytonicchar.PolytonicChar, 0))
	for _, beta := range str {
		betaAccents := map[rune]polytonicchar.Accent{
			'\\': polytonicchar.Grave,
			'/':  polytonicchar.Acute,
			'=':  polytonicchar.Circumflex,
		}
		if accent, ok := betaAccents[beta]; ok {
			if len(polyStr) > 0 {
				polytonicchar.SetAccent((polyStr)[len(polyStr)-1], accent)
			}
		}
		betaSpiritus := map[rune]polytonicchar.Spiritus{
			')': polytonicchar.Lenis,
			'(': polytonicchar.Asper,
		}
		if spiritus, ok := betaSpiritus[beta]; ok {
			if len(polyStr) > 0 {
				polytonicchar.SetSpiritus((polyStr)[len(polyStr)-1], spiritus)
			}
		}
		if '|' == beta {
			if len(polyStr) > 0 {
				polytonicchar.SetIota((polyStr)[len(polyStr)-1], true)
			}
		}
		betaPunctuations := map[rune]polytonicchar.Char{
			' ':  polytonicchar.Space,
			'\n': polytonicchar.Space,
			',':  polytonicchar.Comma,
			'.':  polytonicchar.Dot,
			'\'': polytonicchar.Apostrophe,
			':':  polytonicchar.Interpunct,
			';':  polytonicchar.Question,
		}
		if char, ok := betaPunctuations[beta]; ok {
			polyStr = append(polyStr, polytonicchar.New(char, false))
			if beta == '\n' {
				polytonicchar.SetVariant((polyStr)[len(polyStr)-1], true)
			}
		}
		betaChars := map[rune]polytonicchar.Char{
			'a': polytonicchar.Alpha,
			'b': polytonicchar.Beta,
			'g': polytonicchar.Gamma,
			'd': polytonicchar.Delta,
			'e': polytonicchar.Epsilon,
			'z': polytonicchar.Zeta,
			'h': polytonicchar.Eta,
			'q': polytonicchar.Theta,
			'i': polytonicchar.Iota,
			'k': polytonicchar.Kappa,
			'l': polytonicchar.Lambda,
			'm': polytonicchar.Mu,
			'n': polytonicchar.Nu,
			'c': polytonicchar.Xi,
			'o': polytonicchar.Omicron,
			'p': polytonicchar.Pi,
			'r': polytonicchar.Rho,
			's': polytonicchar.Sigma,
			't': polytonicchar.Tau,
			'u': polytonicchar.Upsilon,
			'f': polytonicchar.Phi,
			'x': polytonicchar.Chi,
			'y': polytonicchar.Psi,
			'w': polytonicchar.Omega,
		}
		if char, ok := betaChars[unicode.ToLower(beta)]; ok {
			polyStr = append(polyStr, polytonicchar.New(char, unicode.IsUpper(beta)))
			if unicode.ToLower(beta) == 's' {
				polytonicchar.SetVariant((polyStr)[len(polyStr)-1], true)
			}
			if len(polyStr) > 1 && polytonicchar.Name(polyStr[len(polyStr)-2]) == polytonicchar.Sigma {
				polytonicchar.SetVariant((polyStr)[len(polyStr)-2], false)
			}
		}
	}

	return polyStr
}

func (s PolytonicString) String() string {
	ret := ""
	for _, v := range s {
		ret += v.String()
	}
	return ret
}
