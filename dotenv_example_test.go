package dotenv_test

import (
	"fmt"
	"os"
	"strings"

	"kkn.fi/dotenv"
)

func ExampleNew() {
	env := dotenv.New()
	r := strings.NewReader("VAR=VAL\n")
	if err := env.Load(r); err != nil {
		fmt.Fprintf(os.Stderr, "error while running example: %v", err)
	}
	fmt.Println(os.Getenv("VAR"))
	// Output:
	// VAL
}
