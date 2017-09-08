package polytonicchar

const replacementChar = string(65533)

func (in *PolytonicChar) String() string {
	if in.variant {
		return variantRepresentation(in)
	}

	if in.name > Omega {
		return punctuationRepresentation(in)
	}

	//START https://en.wikipedia.org/wiki/Greek_Extended
	if in.spiritus != None {
		return spiritedRepresentation(in)
	}
	if in.accent != None || in.iotaSubscriptum {
		return difficultRepresentation(in)
	}
	//END https://en.wikipedia.org/wiki/Greek_Extended

	return simpleRepresentation(in)
}

func variantRepresentation(in *PolytonicChar) string {
	if in.name == Space {
		return "\n"
	}
	if in.name == Sigma && !in.capital {
		const sigmaVariant = 962
		return string(sigmaVariant)
	}
	return replacementChar
}

//in is guaranteed to have variant == false and name>Omega
func punctuationRepresentation(in *PolytonicChar) string {
	nameToString := map[Char]string{
		Space:      " ",
		Comma:      ",",
		Dot:        ".",
		Apostrophe: "'",
		Interpunct: "\u0387",
		Question:   ";",
	}
	if out, ok := nameToString[in.name]; ok {
		return out
	}
	return replacementChar
}

//in is guaranteed to have variant == false, name<=Omega and spiritus != None
func spiritedRepresentation(in *PolytonicChar) string {
	if in.name == Rho {
		const startingPosition = 8164
		const asperOffset = 1
		const capitalOffset = 7

		if in.capital && in.spiritus == Lenis {
			return replacementChar
		}
		if in.iotaSubscriptum || in.accent != None {
			return replacementChar
		}
		out := startingPosition
		if in.spiritus == Asper {
			out += asperOffset
		}
		if in.capital {
			out += capitalOffset
		}
		return string(out)
	}

	const startingPosition = 7936
	const asperOffset = 1
	const accentOffset = 2
	const capitalOffset = 8
	const letterOffset = 16
	const iotaOffset = 128

	out := startingPosition
	if in.spiritus == Asper {
		out += asperOffset
	}
	out += accentOffset * int(in.accent)
	if in.capital {
		out += capitalOffset
	}

	if in.iotaSubscriptum {
		out += iotaOffset
		letterNumber := map[Char]int{
			Alpha: 0,
			Eta:   1,
			Omega: 2,
		}
		if letter, ok := letterNumber[in.name]; ok {
			out += letter * letterOffset
			return string(out)
		}
		return replacementChar
	}

	//no iota subscriptum
	letterNumber := map[Char]int{
		Alpha:   0,
		Epsilon: 1,
		Eta:     2,
		Iota:    3,
		Omicron: 4,
		Upsilon: 5,
		Omega:   6,
	}
	if letter, ok := letterNumber[in.name]; ok {
		out += letter * letterOffset
		return string(out)
	}

	return replacementChar
}

//in is guaranteed to have:
//  variant == false, name<=Omega, spiritus == None
//  and (accent != None or iotaSubscriptum == true)
func difficultRepresentation(in *PolytonicChar) string {
	if in.iotaSubscriptum {
		const startingPosition = 8114
		const noAccentOffset = 1
		const acuteOffset = 2
		const circumflexOffset = 5
		const capitalOffset = 9
		const letterOffset = 16

		out := startingPosition
		if in.accent == None {
			out += noAccentOffset
		}
		if in.accent == Acute {
			out += acuteOffset
		}
		if in.accent == Circumflex {
			out += circumflexOffset
		}
		if in.capital {
			out += capitalOffset
			if in.accent != None {
				return replacementChar
			}
		}
		letterNumber := map[Char]int{
			Alpha: 0,
			Eta:   1,
			Omega: 4,
		}
		if letter, ok := letterNumber[in.name]; ok {
			out += letter * letterOffset
			return string(out)
		}

		return replacementChar
	}
	if in.accent == Circumflex {
		const startingPosition = 8118
		const letterOffset = 16

		if in.capital {
			return replacementChar
		}
		out := startingPosition
		letterNumber := map[Char]int{
			Alpha:   0,
			Eta:     1,
			Iota:    2,
			Upsilon: 3,
			Omega:   4,
		}
		if letter, ok := letterNumber[in.name]; ok {
			out += letter * letterOffset
			return string(out)
		}

		return replacementChar
	}
	if in.capital {
		const startingPosition = 8122
		const shortVowelOffset = -2
		const AcuteOffset = 1
		const letterOffset = 16

		out := startingPosition
		if in.accent == Acute {
			out += AcuteOffset
		}
		if in.name == Epsilon || in.name == Omicron {
			out += shortVowelOffset
		}
		letterNumber := map[Char]int{
			Alpha:   0,
			Epsilon: 1,
			Eta:     1,
			Iota:    2,
			Upsilon: 3,
			Omicron: 4,
			Omega:   4,
		}
		if letter, ok := letterNumber[in.name]; ok {
			out += letter * letterOffset
			return string(out)
		}

		return replacementChar
	}
	const startingPosition = 8048
	const AcuteOffset = 1
	const letterOffset = 2

	out := startingPosition
	if in.accent == Acute {
		out += AcuteOffset
	}
	letterNumber := map[Char]int{
		Alpha:   0,
		Epsilon: 1,
		Eta:     2,
		Iota:    3,
		Omicron: 4,
		Upsilon: 5,
		Omega:   6,
	}
	if letter, ok := letterNumber[in.name]; ok {
		out += letter * letterOffset
		return string(out)
	}

	return replacementChar
}

func simpleRepresentation(in *PolytonicChar) string {
	const lowerCaseAlpha = 945
	const upperCaseAlpha = 913

	if in.capital {
		return string(upperCaseAlpha + in.name)
	}
	return string(lowerCaseAlpha + in.name)
}
