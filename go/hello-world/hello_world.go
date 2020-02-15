package greeting

import (
	"fmt"
)

func HelloWorld() string {
	A := "Hello, World!"
	return A
}

func main() {
	fmt.Print(HelloWorld())

}
