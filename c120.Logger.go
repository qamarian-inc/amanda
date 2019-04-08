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

	// Getting the log files directory. { ...
	if os.Getenv (LOG_FILE_DIR_ENV_VAR) == "" {
		output (fmt.Sprintf ("Startup Error: Environmental variable '%s' (LOG_FILE_DIR_ENV_VAR) is not set: init () in c120_Logger.go", LOG_FILE_DIR_ENV_VAR))
		os.Exit (1)
	}
	// ... }

	// Creating data needed to run the external logger. { ...
	log_File_Path := fmt.Sprintf ("%s%c%s.log", os.Getenv (LOG_FILE_DIR_ENV_VAR), os.PathSeparator, SOFTWARE_ID)
	
	logger___Logging_Info = &qamarian_Logger.Logging_Info {Log_File: log_File_Path}
	// .. }

	// Starting the logger.
	go func () {
		const LOGGER_BUFFER_SIZE = 0
                errM := logger___Logging_Info.Logger (LOGGER_BUFFER_SIZE)

                if errM != nil {
                        output ("State: Logger has shutdown. Reason: " + errM.Error ())
                }
        } ()
}

func logger (new_Log string) (error) { // Call this function, to record a new log. On success, nil error is returned, otherwise a non-nil error is returned.

	// If a panic should occur, it is prevented from affecting caller of this function.
	defer func () {
		recover ()
	} ()

	// Recording log.
	errX := logger___Logging_Info.Log (new_Log)

	return errX
}

func logger___Shutdown () (error) { // Call this function, to gracefully shutdown the logger.

	// If a panic should occur, it is prevented from affecting caller of this function.
	defer func () {
		recover ()
	} ()

	errX := logger___Logging_Info.Shutdown ()

	return errX
}

var (
	logger___VAR_NOT_SET error = errors.New ("The environmental variable is not set.")
	logger___Logging_Info *qamarian_Logger.Logging_Info
)
