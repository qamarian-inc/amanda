package main

/* This virtual component can be used to produce customized console outputs. Simply call its interface "iOutput_AAAAAB ()", to output a customized output. 

	BACKUPS: The following are some actual components capable of backing up this virtual component:
	Actual component aaaaac (Easy-to-read Output)
*/

var (
	iOutput_AAAAAB func (string) /* This interface outputs customized output.

	INPUT
	input 0: The string to be outputted. */
)
