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
	env.Load(r)
	fmt.Println(os.Getenv("VAR"))
	// Output:
	// VAL
}
