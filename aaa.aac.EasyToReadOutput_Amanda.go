package main

/* This component can be used to output data to the console, in an easy-to-read format. */

import (
	"fmt"
)

func iOutput_EasyToReadOutput_Amanda (outputString string) {
	fmt.Println (fmt.Sprintf ("\n    $ (Amanda): %s", outputString))
}