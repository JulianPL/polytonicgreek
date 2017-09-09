# polytonicgreek

This project provides methods for converting betacode into proper polytonic greek letters and for basic string operations.

## How to use

Compile the main package with go.
The main function asks for a number and a betastring and outputs the corresponding greek text with the given width (if possible).

## polytonicchar

This package provides the basic struct for polytonic greek letters:
* name and capital determine the base letter
* iotaSubscriptum, spiritus and accent provide the Diacritics
* variant is for the difference of the two lower case sigmas and for ' ' vs '\n'

## polytonicstring

This package provides the betacode conversion and basic string operations

## TODO

* Add packages for metric and maybe grammar
* Add Trema to polytonicchar
* Add examples
