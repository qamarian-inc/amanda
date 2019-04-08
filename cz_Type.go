package main

type service struct { // A data of this type represents an HTTP service that can be provided. "Service_Path" represents the HTTP path of the service, while "Provider" represents a function that implements the service.

        Service_Path string
        Provider service_Provider
}

type service_Provider func (http.ResponseWriter, *http.Request) // Service providers are functions that serve HTTP requests.

func (provider service_Provider) ServeHTTP (response http.ResponseWriter, request *http.Request) {
        provider (response, request)
}
