package polytonicchar

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

const replacementChar = 65533
const lowerCaseAlpha = 945
const upperCaseAlpha = 913
const greekExtended = 7936
const accentendLowerCaseWithoutSpiritus = 8048
const accentendUpperCaseWithoutSpiritus = 8122
const circumflexWithoutSpiritus = 8118
const iotaSubscriptumWithoutSpiritus = 8114
const asperOffset = 1
const iotaWithCircumflexWithoutSpiritusOffset = 1
const accentOffset = 2
const capitalOffset = 8
const letterOffset = 16
const iotaOffset = 128
const rhoOffset = 228
