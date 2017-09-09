# polytonicgreek

This project provides methods for converting betacode into proper polytonic greek letters and for basic string operations.

## How to use

Compile the main package with go.

The main function asks for a number and a betastring and outputs the corresponding greek text with the given width (if possible).

Betacode works mainly as stated in [Wikipedia](https://en.wikipedia.org/wiki/Beta_Code) with the following exceptions:

* *l can not be used instead of L for any latin letter l
* There is neither Digamma nor Lunate Sigma
* Medial and final Sigma are both represented by s
* Hyphen, Dash and Numeral are not implemented (and probably won't be in the near future)
* Diaeresis, macron and breve are not yet implemented

## polytonicchar

This package provides the basic struct for polytonic greek letters:
* name and capital determine the base letter
* iotaSubscriptum, spiritus and accent provide the Diacritics
* variant is for the difference of the two lower case sigmas and for ' ' vs '\n'

## polytonicstring

This package provides the betacode conversion and basic string operations

## TODO

* Add packages for metric and maybe grammar
* Add Diaeresis to polytonicchar
* Add examples
