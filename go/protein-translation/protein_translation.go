package protein

import (
	"errors"
)

//FromCodon
var mapper = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGS": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var (
	//ErrStop is an errors the test file expects
	ErrStop = errors.New("Terminate Codon")
	//ErrInvalidBase is an errors the test file expects
	ErrInvalidBase = errors.New("Invalid Codon")
)

//FromCodon is the function for testing the RNA translation
func FromCodon(input string) (string, error) {
	if _, ok := mapper[input]; ok {
		if input == "UAA" || input == "UAG" || input == "UGA" {
			return "", ErrStop
		}
		return mapper[input], nil
	}
	return "", ErrInvalidBase
}

//FromRNA is the function for testing the RNA translation
func FromRNA(input string) ([]string, error) {
	b := make([]string, 0)
	for j := 0; j < len(input); j += 3 {
		h, l := FromCodon(input[j : j+3])
		if l == ErrStop {
			return b, nil
		}
		if l == ErrInvalidBase {
			return b, l
		}
		b = append(b, h)
	}
	return b, nil
}
