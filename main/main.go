package main

import (
	"fmt"
	"polytonicgreek/polytonicchar"
	"polytonicgreek/polytonicstring"
)

func main() {
	T := "Au)ta\\r o( e)k lime/nos trhxei=an a)tarpo\\n di' a)/krias prose/bh a)na/ xw=ron u(lh/enta, "
	T += "h(=| oi( A)qh/nh di=on u(forbo/n pe/frade, "
	T += "o(/ oi( ma/lista bio/toio oi)kh/wn kh/deto, "
	T += "ou(\\s di=os O)dusseu/s kth/sato.\n"

	T += "To\\n d' a)/ra h(/menon e)ni\\ prodo/mw| eu(=r', "
	T += "e)/nqa oi( au)lh\\ u(yhlh\\ kalh\\ te mega/lh te, "
	T += "peri/dromos, "
	T += " e)ni\\ xw/rw| periske/ptw|, "
	T += "de/dmhto. "
	T += "h(/n r(a subw/ths au)to\\s u(/essin a)poixome/noio a)/naktos, "
	T += "no/sfin despoi/nhs kai\\ Lae/rtao ge/rontos, "
	T += "r(utoi=sin la/essi dei/mato kai\\ a)xe/rdw| e)qri/gkwsen"

	str := polytonicstring.New(T)
	char := polytonicchar.New(polytonicchar.Omicron, false)
	polytonicchar.SetSpiritus(char, polytonicchar.Asper)
	fmt.Println(polytonicstring.InString(char, str, true))
	sub := polytonicstring.New("o(")
	polytonicstring.WrapString(str, 80)
	fmt.Println(polytonicstring.SubString(sub, str, true))
	fmt.Println(sub)
	fmt.Println(str)
}
