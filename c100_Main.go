package main

import (
        "context"
	"fmt"
        "github.com/gorilla/mux"
	"net/http"
        "os"
	"runtime"
	"strconv"
        "time"
)

func main () { // This function helps the software start a service server and an admin server. All software services and admin services added to 'software_Service' and 'admin_Service' (both in file cz.Global.go), will be available via the service server and the admin server, respectively.

	output (fmt.Sprintf ("Ware %s (%s) is starting up (framework: Amanda)...", SOFTWARE_NAME, SOFTWARE_ID))

	// If a panic should occur, the panic is logged. { ...
        defer func () {
                panic_Reason := recover ()
                if panic_Reason != nil {
                        logger ("main () paniced")
                }
                os.Exit (1)
        } ()
        // ... }

	// Using value of 'GOMAXPROCS' specified in the configuration file. { ...
	go_Max_Procs, errX := conf_Data_Provider ("go_Max_Procs")
	if errX != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'go_Max_Procs', from the configuration file: main ()", errX.Error ()))
		os.Exit (1)
	}

	go_Max_Procs_As_Int, errY := strconv.Atoi (go_Max_Procs)
	if errY != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Converting value of 'go_Max_Procs', from string to integer: main ()", errY.Error ()))
		os.Exit (1)
	}

	runtime.GOMAXPROCS (go_Max_Procs_As_Int)
	// ... }

	// Fetching network address and port of service server. { ...
	service_Addr, errA := conf_Data_Provider ("service_Addr.addr")
	service_Port, errB := conf_Data_Provider ("service_Addr.port")

	if errA != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'service_Addr.addr', from the configuration file: main ()", errA.Error ()))
		os.Exit (1)
	}
	if errB != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'service_Addr.port', from the configuration file: main ()", errB.Error ()))
		os.Exit (1)
	}
	// ... }
	
	// Fetching network address and port of admin server. { ...
	admin_Addr, errC := conf_Data_Provider ("admin_Addr.addr")
	admin_Port, errD := conf_Data_Provider ("admin_Addr.port")

	if errC != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'admin_Addr.addr', from the configuration file: main ()", errC.Error ()))
		os.Exit (1)
	}
	if errD != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'admin_Addr.port', from the configuration file: main ()", errD.Error ()))
		os.Exit (1)
	}
	// ... } 

	// Checking if TLS certificate and private key are available. { ...
	if os.Getenv (TLS_CERT_FILEPATH) == "" || os.Getenv (PRIV_KEY_FILEPATH) == "" {
		output (fmt.Sprintf ("Startup Error: Filepath of TLS certificate and/or private key, are not set --> \n Fetching filepaths of TLS certificate and private key, from environmetal variables (%s and %s): main ()", TLS_CERT_FILEPATH, PRIV_KEY_FILEPATH))
		os.Exit (1)
	}
	// ... }

	// Creating service server. { ...
	main___Service_Server_Info := &http.Server {	
                Addr:            service_Addr + ":" + service_Port,
                ReadTimeout:     main___Network_IO_Timeout,
                WriteTimeout:    main___Network_IO_Timeout,
                MaxHeaderBytes:  1 << 20,
        }

        routerX := mux.NewRouter ()
        for _, service := range software_Service {
                routerX.HandleFunc (service.Service_Path, service.Provider)
        }
        main___Service_Server_Info.Handler = routerX
        // ... }

        // Starting service server, as a different routine. { ...
        go func () { 
                // Note, this function blocks.
                errJ := main___Service_Server_Info.ListenAndServeTLS (os.Getenv (TLS_CERT_FILEPATH), os.Getenv (PRIV_KEY_FILEPATH))
                
                output ("State: Service server shutting down...")

                // Log is recorded, if server shutdowns due to an error.
                if errJ != nil && errJ != http.ErrServerClosed { 
                        logger (fmt.Sprintf ("%s ---> \n Running service server: main ()", errJ.Error ()))
                        output (fmt.Sprintf ("Runtime Error: Service Server:", errJ.Error ()))
                }
        } ()
        // ... }

        output ("State: Service server has started!")

        // Creating service server. { ...
	main___Admin_Server_Info := &http.Server {	
                Addr:            admin_Addr + ":" + admin_Port,
                ReadTimeout:     main___Network_IO_Timeout,
                WriteTimeout:    main___Network_IO_Timeout,
                MaxHeaderBytes:  1 << 20,
        }

        routerY := mux.NewRouter ()
        for _, service := range admin_Service {
                routerY.HandleFunc (service.Service_Path, service.Provider)
        }
        main___Admin_Server_Info.Handler = routerY
        // ... }

        // Starting admin server. Note: this function blocks. { ...
        output ("State: Starting admin server...")

        errK := main___Admin_Server_Info.ListenAndServeTLS (os.Getenv (TLS_CERT_FILEPATH), os.Getenv (PRIV_KEY_FILEPATH))
        // ... }

        output ("State: Admin server shutting down...")

        // Log is recorded, if admin server shutdowns due to an error.
        if errK != nil && errK != http.ErrServerClosed {
                logger (fmt.Sprintf ("%s ---> \n Running admin server: main ()", errK.Error ()))
                output ("Runtime Error: Admin Server: " + errK.Error ())
                return
        }

        // If server stopped because it was asked to, main () waits for os.Exit () to be called in main___Shutdown_Servers (), just to ensure graceful shutdown.
        if errK == http.ErrServerClosed {
                for {
                        runtime.Gosched ()
                }
        }
}

func main___Shutdown_Servers () { // This function is an interface for main (). This function tells main () to shutdown the service server and the admin server of this software. To gracefully shutdown this software, this function can be called.

        // If a panic should occur during shutdown, the panic is logged. { ...
        defer func () {
                panic_Reason := recover ()
                if panic_Reason != nil {
                        logger ("main___Shutdown_Servers () paniced during graceful shutdown")
                }
                os.Exit (1)
        } ()
        // ... }

        // Shutting down service server. { ...
        errX := main___Service_Server_Info.Shutdown (context.Background ())
        if errX != nil {
                error_Message := fmt.Sprintf ("%s ---> \n Gracefully shutting down service server main___Shutdown_Servers ()", errX.Error ())
                logger (error_Message)
        }
        // ... }

        // Shutting down admin server. { ...
        err2 := main___Admin_Server_Info.Shutdown (context.Background ())
        if err2 != nil {
                error_Message := fmt.Sprintf ("%s ---> \n Gracefully shutting down admin server main___Shutdown_Servers ()", err2.Error ())
                logger (error_Message)
        }
        // ... }

        // Shutting down all necessary components. { ...
        logger___Shutdown ()
        alert_Raiser___Shutdown ()
        // ... }

        // Finally halting the software.
        os.Exit (0)
}

var (
        main___Service_Server_Info *http.Server // Server that checks authentication state of users.
        main___Admin_Server_Info *http.Server // Server that provides admin services.

        main___Network_IO_Timeout time.Duration = time.Second * 60
)
