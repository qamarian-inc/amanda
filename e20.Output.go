package main

import "fmt"

func output (output_String string) {
	fmt.Println (fmt.Sprintf ("\n    $ (%s): %s", SOFTWARE_ID, output_String))
}
