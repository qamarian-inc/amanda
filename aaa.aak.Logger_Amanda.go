package main

/* This component is a logger. When different goroutines simultanously ask it to record logs, it queues all the request, then records them one-by-one, instead of trying to write them to the log file all at once (thereby degrading the performance of the app using it). In short, this logger is designed not to degrade the performance of apps using it, unlike most other loggers. 

   This component has special case init and deinit functions, and they should be registered with component "aaa.aaa".

   DEPENDENCIES
   	Virtual component aaa.aab (Customized output assistant)
   	Virtual component aaa.aad (Onion path decoder)
   	Virtual component aaa.aaf (Configuration data provider)
   	Virtual component aaa.aah (Critical Event Zain)
*/

import (
	"fmt"
	qamarian_Logger "github.com/qamarian-inc/logger"
	"os"
)

func iInit_Logger_Amanda () { // Initializes this component. Remember to register this with component "aaa.aaa".

	// Getting the filepath of the log file. { ...
	logfilePath, errT := iScalarData_vcConfDataProvider_Amanda ("LogfilePath_aaa_aak")

	if errT != nil {
		iOutput_vcCustomizedOutput_Amanda (fmt.Sprintf ("Startup Error: %s ---> Getting the filepath of app's log file: iInit_Logger_Amanda ()", errT.Error ()))
		os.Exit (1)
	}
	// ... }

	// The filepath of the app's log file is expected to be onion-formatted, and this section does the decoding of the filepath into its real form. { ...
	logfilePath, errV := iDecode_vcOnionPathDecoder_Amanda (logfilePath)

	if errV != nil {
		iOutput_vcCustomizedOutput_Amanda (fmt.Sprintf ("Startup Error: %s ---> Decoding the filepath of app's log file, from its onion form into its real form: iInit_Logger_Amanda ()", errV.Error ()))
		os.Exit (1)
	}
	// ... }

	dLoggingInfo_Logger_Amanda = &qamarian_Logger.Logging_Info {Log_File: logfilePath} // Creating data needed to run the external logger.

	// Starting the logger.
	go logger_Amanda ()
}

func logger_Amanda () {
	// If a panic should occur, it is prevented from affecting other routines.
	defer func () {
		recover ()
	} ()

        errM := dLoggingInfo_Logger_Amanda.Logger (0)

        if errM != nil {
                iOutput_vcCustomizedOutput_Amanda ("State: Logger has shutdown due to an error. Error: " + errM.Error ())
                iReport_vcCriticalEventZain_Amanda ("Logger has shutdown due to an error. Error: " + errM.Error ())
        }
}

func iRecord_Logger_Amanda (new_Log string) (error) { /* Call this interface, to record a new log.

	INPUT
	input 0: The log to be recorded.

	OUTPT
	outpt 0: Any error that occurs during the logging. If logging succeeds, value would be nil. If logging fails, value would the error that occured. */

	// If a panic should occur, it is prevented from affecting caller of this function.
	defer func () {
		recover ()
	} ()

	// Recording log.
	errX := dLoggingInfo_Logger_Amanda.Log (new_Log)

	return errX
}

func iDeinit_Logger_Amanda () (error) { // Remember to register this with component "aaa.aaa".

	// If a panic should occur, it is prevented from affecting caller of this function.
	defer func () {
		recover ()
	} ()

	errX := dLoggingInfo_Logger_Amanda.Shutdown ()

	return errX
}

var (
	dLoggingInfo_Logger_Amanda *qamarian_Logger.Logging_Info
)
