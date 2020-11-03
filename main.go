package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

func main() {

	greetmsgempty := hello("")
	fmt.Println(aurora.Red(greetmsgempty))

	greetmsg := hello("john")
	fmt.Println(aurora.Yellow(greetmsg))
}
