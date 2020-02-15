package strand

import "strings"

//ToRNA is the function tht the tests expect to pass arguments to
func ToRNA(dna string) string {

	var t = map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}
	answer := make([]string, len(dna))
	for i := 0; i < len(dna); i++ {
		answer[i] = t[string(dna[i])]
	}
	return strings.Join(answer, "")

}
