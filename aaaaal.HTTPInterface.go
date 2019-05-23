package main

/* This component can be used to turn an Amada Ware into a cloud application. It enable components of your app to be accessed via HTTP. For example, you can make loading of a url (e.g. "https://localhost:90/componentX") trigger the execution of a component in your app.

        HOW IT WORKS
        When this component receives an HTTP request, it does the following:

                1. Checks its routing table to see who to direct the request to
                2. If it finds a controller that can handle the request, it handles the request over to the controller, but if there are no controller to handle the request, error 404 is returned.
                3. The controller can then decide what to do with the HTTP reuest.

                ** Controllers can be created in file "aaaaal.zController.go", and registered in file "aaaaal.zRouter.go".

        DEPENDENCIES
        Virtual component aaaaab (Customized output assistant)
        Virtual component aaaaaf (Configuration data provider)
        Virtual component aaaaah (Critical Event Zain)

        USAGE NOTES
        1. Always remember to call interface "iShutdown_AAAAAL ()" before shutting down an app using this component.
*/

import (
        "context"
	"fmt"
        "github.com/gorilla/mux"
	"net/http"
        "os"
	"strconv"
        "time"
)

func init () { 
	// Fetching network address and port to use. { ...
	netAddr, errA := iScalarData_AAAAAF ("AAAAAL.NetAddr")
	netPort, errB := iScalarData_AAAAAF ("AAAAAL.NetPort")

	if errA != nil {
		output := fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'AAAAAL.NetAddr', from the conf file: init () in aaaaal.HTTPInterface.go", errA.Error ())
		iOutput_AAAAAB (output)
		os.Exit (1)
	}
	if errB != nil {
		output := fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'AAAAAL.NetPort', from the conf file: init () in aaaaal.HTTPInterface.go", errB.Error ())
		iOutput_AAAAAB (output)
		os.Exit (1)
	}
	// ... }

	// Fetching max duration allowed for net input/output. { ...	
        netIODurationBeforeTimeout, errP := iScalarData_AAAAAF ("AAAAAL.MaxDurationForNetIO")
        if errP != nil {
                output := fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'AAAAAL.MaxDurationForNetIO', from the conf file: init () in aaaaal.HTTPInterface.go", errP.Error ())
                iOutput_AAAAAB (output)
                os.Exit (1)
        }

        // Casting duration from string to int. { ...
        intNetIOMaxDuration, errH := strconv.Atoi (netIODurationBeforeTimeout)
        if errH != nil {
                output := fmt.Sprintf ("Startup Error: %s ---> \n Value of 'AAAAAL.MaxDurationForNetIO' (from the conf file) is not a number: init () in aaaaal.HTTPInterface.go", errH.Error ())
                iOutput_AAAAAB (output)
                os.Exit (1)
        }
        // ... }
        // ... }

        // Fetching max size allowed for HTTP request headers. { ...
        httpMaxReqHeaderSize, errQ := iScalarData_AAAAAF ("AAAAAL.HttpMaxReqHeaderSize")
        if errQ != nil {
        	output := fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'AAAAAL.HttpMaxReqHeaderSize', from the conf file: init () in aaaaal.HTTPInterface.go", errQ.Error ())
                iOutput_AAAAAB (output)
                os.Exit (1)
        }

        // Casting max header size from string to int. { ...
        intReqHeaderSize, errI := strconv.Atoi (httpMaxReqHeaderSize)
        if errI != nil {
        	output := fmt.Sprintf ("Startup Error: %s ---> \n Value of 'AAAAAL.HttpMaxReqHeaderSize' (from the conf file) is not a number: init () in aaaaal.HTTPInterface.go", errI.Error ())
                iOutput_AAAAAB (output)
                os.Exit (1)
        }
        // ... }
        // ... }

        // Composing HTTP server details to be used. { ...
	dServerInfo_AAAAAL := &http.Server {	
                Addr:            netAddr + ":" +  netPort,
                ReadTimeout:     time.Duration (intNetIOMaxDuration) * time.Second,
                WriteTimeout:    time.Duration (intNetIOMaxDuration) * time.Second,
                MaxHeaderBytes:  intReqHeaderSize,
        }

        router := mux.NewRouter ()

        // Registering routing rules with HTTP router.
        for _, rule := range dRoutingRule_AAAAAL {
                router.HandleFunc (rule.route, rule.controller)
        }

        dServerInfo_AAAAAL.Handler = router
        // ... }

        // Determining if HTTP or HTTPS should be be used. { ...

        // Fetching filepath of TLS cert bundle. { ...
        certBundle, errO := iScalarData_AAAAAF ("AAAAAL.TLSCertBundle")
        if errO != nil {
        	output := fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'AAAAAL.TLSCertBundle', from the configuration file: init () in aaaaal.HTTPInterface.go", errO.Error ())
                iOutput_AAAAAB (output)
                os.Exit (1)
        }
        // ... }

        // Fetching filepath of TLS private key. { ...
        privateKey, errT := iScalarData_AAAAAF ("AAAAAL.TLSPrivateKey")
        if errT != nil {
        	output := fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'AAAAAL.TLSPrivateKey', from the configuration file: init () in aaaaal.HTTPInterface.go", errT.Error ())
                iOutput_AAAAAB (output)
                os.Exit (1)
        }
        // ... }

        httpProtocolInUse := "" // Declaration of a variable that'll be used later.

        if certBundle != "" && privateKey != "" {
                httpProtocolInUse = "HTTPS"
        } else {
                httpProtocolInUse = "HTTP"
                iOutput_AAAAAB ("Note: Your app will be using HTTP (not HTTPS) since the filepath of one or both of the certificate bundle and private key are not set.")
        }
        // ... }

        // Startup notification.
        iOutput_AAAAAB (fmt.Sprintf ("HTTP interface is starting up now! NETWORK ADDRESS: %s:%s (%s)", netAddr, netPort, httpProtocolInUse))

        // Starting server, as a different routine.
        go func () {
        	errJ := *new (error) // Declaration of a variable that'll be used later.

        	if httpProtocolInUse == "HTTPS" { // HTTPS startup.
                	errJ = dServerInfo_AAAAAL.ListenAndServeTLS (certBundle, privateKey) // Note, this function blocks.
        	} else { // HTTP startup.

	                errJ = dServerInfo_AAAAAL.ListenAndServe () // Note, this function blocks.
        	}
                
                // By the time execution reachs here, the server must have shutdown.
	        iOutput_AAAAAB ("State: Server has shutdown!")

	        // If server shutdowns due to an error, a log is recorded, and a critical event zain is notified.
        	if errJ != nil && errJ != http.ErrServerClosed {
                	iOutput_AAAAAB (fmt.Sprintf ("Runtime Error: %s ---> \n Running HTTP interface server", errJ.Error ()))

	                iRecord_AAAAAJ (fmt.Sprintf ("%s ---> \n Running HTTP interface server", errJ.Error ()))

        	        iBeInformed_AAAAAH (fmt.Sprintf ("HTTP interface server has shutdown due to an error: %s", errJ.Error ()))
        	}
        } ()
}

func iShutdown_AAAAAL () { // This interface should be called, to gracefully shutdown this component.

        // If a panic should occur, it is prevented from affecting other components.
        defer func () {
                recover ()
        } ()

        // Shutting down server gracefully. { ...
        errX := dServerInfo_AAAAAL.Shutdown (context.Background ())
        if errX != nil {
                errorMessage := fmt.Sprintf ("%s ---> \n Gracefully shutting down service server: iShutdown_AAAAAL ()", errX.Error ())
                iRecord_AAAAAJ (errorMessage)
        }
        // ... }
}

var (
	dServerInfo_AAAAAL *http.Server // Information needed to run the interface's server.
)

type controller_AAAAAL func (http.ResponseWriter, *http.Request)

func (controller controller_AAAAAL) ServeHTTP (res http.ResponseWriter, req *http.Request) {
        controller (res, req)
}

type routingRule_AAAAAL struct {
	route string
	controller controller_AAAAAL
}
