package main

import (
        "fmt"
        "github.com/gorilla/mux"
        "net/http"
)

var shutdown_Manager service_Provider = func (response_Res http.ResponseWriter, request *http.Request) { // This service provider is an admin service provider. It is considered an object, and its job is to gracefully shutdown this software.

        // If panic occurs while serving request, it is logged. { ...
        defer func () {
                panic_Reason := recover ()
                if panic_Reason != nil {
                        logger ("shutdown_Manager () paniced while serving a request.")
                }
        } ()
        // ... }

        // If admin code provided is invalid: request is not granted and error is returned.
        if admin_Code_Manager () != mux.Vars (request)["admin_Code"] { 
                response_Res.WriteHeader (http.StatusUnauthorized)
                fmt.Fprintf (response_Res, "Request not granted: incorrect access code.")
                return
        }

        response_Res.WriteHeader (http.StatusOK)
        fmt.Fprintf (response_Res, "Ok. Shutting down...")

        go main___Shutdown_Servers ()
}
