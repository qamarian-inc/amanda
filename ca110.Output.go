package main

/* This component formats outputs going to the console. */

import "fmt"

func output (output_String string) {
	fmt.Println (fmt.Sprintf ("\n    $ (Amanda): %s", output_String))
}
