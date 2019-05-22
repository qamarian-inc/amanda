package main

/* This component is a logger.

	When different goroutines simultanously ask it to record logs, it queues all the requests, then records them one-by-one, instead of trying to write them to the log file all at once (thereby degrading the performance of the app using it).

	In short, this logger is designed not to degrade the performance of apps using it, unlike most other loggers.

	DEPENDENCIES
   	Virtual component aaaaab (Customized output assistant)
   	Virtual component aaaaad (Onion path decoder)
   	Virtual component aaaaaf (Configuration data provider)
   	Virtual component aaaaah (Critical Event Zain)

	USAGE NOTES
	1. Ensure all dependencies needed by this component are present in your app's source code, before building your code.

	2. When about starting an app whose source code contains this component, do the following:
   		
   		- create a log file, on the computer on which the app would run
   		- set the logfile's path as conf data "AAAAAK_LogfilePath", in the conf file of dependency aaaaaf (remember "aaaaaf" was listed as one of the dependencies needed by this component)

   	3. To tell this component to record a log, use interface "iRecord_AAAAAK ()".

   	4. Always remember to call interface "iShutdown_AAAAAK ()" before shutting down an app using this component.
*/

import (
	"fmt"
	qamarian_Logger "github.com/qamarian-inc/logger"
	"os"
	"strings"
)

func init () {

	// Getting the filepath of the log file. { ...
	logfilePath, errT := iScalarData_AAAAAF ("AAAAAK_LogfilePath")

	if errT != nil {
		iOutput_AAAAAB (fmt.Sprintf ("Startup Error: %s ---> \n Getting the filepath of app's log file: iInit_AAAAAK ()", errT.Error ()))
		os.Exit (1)
	}
	// ... }

	logfilePath = strings.Trim (logfilePath, " ")

	// Halts app, if an empty string is set as the logfile's path.
	if logfilePath == "" {
		iOutput_AAAAAB (`Startup Error: An empty string is set as the logfile's path, in the conf file: iInit_AAAAAK ()`)
		os.Exit (1)
	}

	// The filepath of the app's log file is expected to be onion-formatted, and this section does the decoding of the filepath into its real form. { ...
	logfilePath, errV := iDecode_AAAAAD (logfilePath)

	if errV != nil {
		iOutput_AAAAAB (fmt.Sprintf ("Startup Error: %s ---> Decoding the filepath of app's log file, from its onion form into its real form: iInit_AAAAAK ()", errV.Error ()))
		os.Exit (1)
	}
	// ... }

	dLoggingInfo_AAAAAK = &qamarian_Logger.Logging_Info {Log_File: logfilePath} // Creating data needed to run the external logger.

	// Starting the logger.
	go cLogger ()
}

func cLogger () {
	// If a panic should occur, it is prevented from affecting other routines.
	defer func () {
		recover ()
	} ()

        errM := dLoggingInfo_AAAAAK.Logger (0)

        if errM != nil {
                iOutput_AAAAAB ("State: Logger has shutdown due to an error. Error: " + errM.Error ())
                iBeInformed_AAAAAH ("Logger has shutdown due to an error. Error: " + errM.Error ())
        }
}

func iRecord_AAAAAK (new_Log string) (error) { /* Call this interface, to record a new log.

	INPUT
	input 0: The log to be recorded.

	OUTPT
	outpt 0: Any error that occurs during the logging. If logging succeeds, value would be nil. If logging fails, value would the error that occured. */

	// If a panic should occur, it is prevented from affecting caller of this function.
	defer func () {
		recover ()
	} ()

	// Recording log.
	errX := dLoggingInfo_AAAAAK.Log (new_Log)

	return errX
}

func iShutdown_AAAAAK () { // This interface must to be called to ensure graceful shutdown of this component. Shutting down of an app without calling this interface, can leave some logs unrecorded.

	// If a panic should occur, it is prevented from affecting caller of this function.
	defer func () {
		recover ()
	} ()

	dLoggingInfo_AAAAAK.Shutdown ()
}

var (
	dLoggingInfo_AAAAAK *qamarian_Logger.Logging_Info
)
