package main

/* This virtual component can be used to record logs. */

var iRecord_vcLogger_Amanda func (string) (error)  /* Call this interface, to record a log.

	INPUT
	input 0: The log to be recorded.

	OUTPT
	outpt 0: Any error that occurs during the logging. If logging succeeds, value would be nil. If logging fails, value would the error that occured. */
