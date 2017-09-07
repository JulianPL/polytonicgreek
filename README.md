# polytonicgreek

This project provides methods for converting betacode into proper polytonic greek letters and for basic string operations.

## TODO

Add packages for metric and maybe grammar
Add functionality for the main package
Remove the mess of the constants in polytonicchar.String()
Add testcases
Add examples

## polytonicchar

This package provides the basic struct for polytonic greek letters:
* name and capital determine the base letter
* iotaSubscriptum, spiritus and accent provide the Diacritics
* variant is for the difference of the two lower case sigmas and for ' ' vs '\n'

## polytonicstring

This package provides the betacode conversion and basic string operations
