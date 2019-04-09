package main

import (
        "context"
	"fmt"
        "github.com/gorilla/mux"
	"net/http"
        "os"
	"runtime"
	"strconv"
)

func main () { // This function helps the software start its server. All services added to variable 'software_Service' (in file ze30.Global.go), will be available via the server.

	output (fmt.Sprintf ("Ware '%s' (%s) is starting up (framework: Amanda)...", SOFTWARE_NAME, SOFTWARE_ID))

        // Initializing services requiring initialization at startup time.
        for _, service := range services_Init {
                service ()
        }

	// If a panic should occur, the panic is logged. { ...
        defer func () {
                panic_Reason := recover ()
                if panic_Reason != nil {
                        logger ("main () paniced")
                }
                os.Exit (1)
        } ()
        // ... }

	// Using the value of 'GOMAXPROCS' specified in the configuration file. { ...
	go_Max_Procs, errX := conf_Data_Provider ("Go_Max_Procs")
	if errX != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'Go_Max_Procs', from the configuration file '%s': main ()", CONF_FILE, errX.Error ()))
		os.Exit (1)
	}

	go_Max_Procs_As_Int, errY := strconv.Atoi (go_Max_Procs)
	if errY != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Converting value of 'Go_Max_Procs', from string to integer: main ()", CONF_FILE, errY.Error ()))
		os.Exit (1)
	}

	runtime.GOMAXPROCS (go_Max_Procs_As_Int)
	// ... }

	// Fetching network address and port of software's server. { ...
	software_Addr, errA := conf_Data_Provider ("Service_Addr.Addr")
	software_Port, errB := conf_Data_Provider ("Service_Addr.Port")

	if errA != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'Service_Addr.Addr', from the configuration file '%s': main ()", CONF_FILE, errA.Error ()))
		os.Exit (1)
	}
	if errB != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'Service_Addr.Port', from the configuration file '%s': main ()", CONF_FILE, errB.Error ()))
		os.Exit (1)
	}
	// ... }
	
	// Checking if TLS certificate and private key are available. { ...
	if os.Getenv (TLS_CERT_FILEPATH) == "" || os.Getenv (PRIV_KEY_FILEPATH) == "" {
		output (fmt.Sprintf ("Startup Error: Filepath of TLS certificate and/or private key, are not set --> \n Fetching filepaths of TLS certificate and private key, from environmetal variables %s and %s respectively: main ()", TLS_CERT_FILEPATH, PRIV_KEY_FILEPATH))
		os.Exit (1)
	}
	// ... }

	// Creating server. { ...
	main___server_Info := &http.Server {	
                Addr:            software_Addr + ":" + software_Port,
                ReadTimeout:     SERVER_NET_IO_TIMEOUT,
                WriteTimeout:    SERVER_NET_IO_TIMEOUT,
                MaxHeaderBytes:  SERVER_MAX_HEADER_SIZE,
        }

        routerX := mux.NewRouter ()
        for _, service := range software_Service {
                routerX.HandleFunc (service.Service_Path, service.Provider)
        }
        main___server_Info.Handler = routerX
        // ... }

        output ("State: Server will start up now!")

        // Starting server.
        errJ := main___server_Info.ListenAndServeTLS (os.Getenv (TLS_CERT_FILEPATH), os.Getenv (PRIV_KEY_FILEPATH)) // Note, this function blocks.
                
        output ("State: Server has shutdown!")

        // Log is recorded, if server shutdowns due to an error.
        if errJ != nil && errJ != http.ErrServerClosed {
                output (fmt.Sprintf ("Runtime Error: Server: %s", errJ.Error ()))
                logger (fmt.Sprintf ("%s ---> \n Running server: main ()", errJ.Error ()))

                // Gracefully shutting down server.
                main___Shutdown_Servers ()
        }

        // If server stopped because it was asked to, main () waits for os.Exit () to be called in main___Shutdown_Servers (), just to ensure graceful shutdown.
        if errJ == http.ErrServerClosed {
                for {
                        runtime.Gosched ()
                }
        }
}

func main___Shutdown_Servers () { // This function is an interface of main (); it tells main () to shutdown the server of this software. To gracefully shutdown this software, this function can be called.

        // If a panic should occur during shutdown, the panic is logged. { ...
        defer func () {
                panic_Reason := recover ()
                if panic_Reason != nil {
                        logger ("main___Shutdown_Servers () paniced during graceful shutdown")
                }
                os.Exit (1)
        } ()
        // ... }

        // Shutting down server gracefully. { ...
        errX := main___server_Info.Shutdown (context.Background ())
        if errX != nil {
                error_Message := fmt.Sprintf ("%s ---> \n Gracefully shutting down service server main___Shutdown_Servers ()", errX.Error ())
                logger (error_Message)
        }
        // ... }

        // Shutting down other components that should be explicitly shutdown. { ...
        logger___Shutdown ()
        alert_Raiser___Shutdown ()
        // ... }

        // Finally halting the software.
        os.Exit (0)
}

var main___server_Info *http.Server // Information needed to run the software's server.
