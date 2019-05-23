package main

/* This component can be used to output data to the console, in an easy-to-read format. Simply call its interface "iOutput_AAAAAC ()", to use. */

import "fmt"

func iOutput_AAAAAC (outputString string) { /* This interface outputs customized output.
	
	INPUT
	input 0: The string to be outputted. */

	fmt.Println (fmt.Sprintf ("\n    $ (Amanda): %s", outputString))
}
