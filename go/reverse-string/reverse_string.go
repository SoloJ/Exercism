package reverse

//Reverse is the function that takes an input string nd returns it in reveres
func Reverse(in string) string {

	var out string
	for _, val := range in {
		out = string(val) + out
	}
	return out
}
