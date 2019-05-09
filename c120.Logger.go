package main

import (
	"errors"
	"fmt"
	qamarian_Logger "github.com/qamarian-inc/logger"
	"os"
)

/* This component implements a synchronous logger which does not trigger the creation of more than one OS thread.

On startup, the init () function starts the logger (as a goroutine), and logs can be sent to this logger, using function logger ().

To gracefully shutdown the logger, call function: logger___Shutdown (). This function should always be called before halting a software, to ensure all pending logs are recorded, before shutdown.
*/

func init () { // Initializes this component.

	// If the built-in logger of this framework (in other words, this component) is not the default logger in use, this component would not bother initializing, as it assumes it would not be used.
	if RECORD_LOG != logger___Record {
		return
	}

	// Getting the filepath of the log file. { ...
	log_File_Path, errT := conf_Data_Provider ("Log_File_Path")

	if errT != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> Getting the filepath of app's log file: init () in c120_Logger.go", errT.Error ()))
		os.Exit (1)
	}
	// ... }

	// The filepath of the app's log file is expected to be onion-formatted, and this section does the decoding of the filepath into its real form. { ...
	log_File_Path, errV := onion_Filepath_Decoder (log_File_Path)

	if errT != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> Decoding the filepath of app's log file, from its onion form into its real form: init () in c120_Logger.go", errT.Error ()))
		os.Exit (1)
	}
	// ... }

	logger___Logging_Info = &qamarian_Logger.Logging_Info {Log_File: log_File_Path} // Creating data needed to run the external logger.

	// Starting the logger.
	go logger ()
}

func logger () {
	// If a panic should occur, it is prevented from affecting other routines.
	defer func () {
		recover ()
	} ()

        errM := logger___Logging_Info.Logger (0)

        if errM != nil {
                output ("State: Logger has shutdown due to an error. Error: " + errM.Error ())
                CRITICAL_EVENT_ACTION ("Logger has shutdown due to an error. Error: " + errM.Error ())
        }
}

func logger___Record (new_Log string) (error) { // Call this function, to record a new log. On success, nil error is returned, otherwise a non-nil error is returned.

	// If the built-in logger of this framework (in other words, this component) is not the default logger in use, this function tells the caller.
	if RECORD_LOG != logger___Record {
		return logger___NOT_RUNNING_NOT_DEFAULT
	}

	// If a panic should occur, it is prevented from affecting caller of this function.
	defer func () {
		recover ()
	} ()

	// Recording log.
	errX := logger___Logging_Info.Log (new_Log)

	return errX
}

func logger___Shutdown () (error) { // Call this function, to gracefully shutdown the logger.

	// If the built-in logger of this framework (in other words, this component) is not the default logger of the app, this function ignores the shutdown request since this component would not be running.
	if RECORD_LOG != logger___Record {
		return nil
	}

	// If a panic should occur, it is prevented from affecting caller of this function.
	defer func () {
		recover ()
	} ()

	errX := logger___Logging_Info.Shutdown ()

	return errX
}

var (
	logger___Logging_Info *qamarian_Logger.Logging_Info
	logger___NOT_RUNNING_NOT_DEFAULT errors = errors.New ("This logger is not running since it is not the default logger of the app.")
)
