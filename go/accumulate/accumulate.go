package accumulate

//Accumulate is the function that the test expects to pas arguments to
func Accumulate(given []string, converter func(string) string) []string {
	out := make([]string, 0)
	for _, val := range given {
		out = append(out, converter(val))
	}
	return out
}
