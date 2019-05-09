package main

var (
	// Modify the following constants as appropriate.

	CONF_FILE string = "z.Conf.yml" // Change this to the filepath of the app's configuration file (an onion-formatted filepath is allowed). Only YAML files are currently supported, and the extension of the file must be ".yaml.

	RECORD_LOG func (string) error = logger___Record // This decides what the default logger of the app should be. If you do not want this framework to use the built-in logger, change this to the logging function of your desired logger.

	SHUTDOWN_LOGGER func () error = logger___Shutdown // If the built-in logger is not your default logger, and your default logger has a shutdown function, set its shutdown function as the value of this constant.

	// If you want this framework to do something whenever a critical event occurs (for example if the server crashes unexpectedly), create a function, and set it as the value of this constant. Note, the function must be able to take at least one string argument (as this framework will always provide information about critical events).
	CRITICAL_EVENT_ACTION func (string) = func (event_Info string) {
		output ("Critical event: " + event_Info)
	}
)

var (   // ADD YOUR APP SERVICES HERE

	// Services that the software provides should be added to this slice.
	software_Service []service = []service {
		service {"/shutdown/{admin_Code}", shutdown_Manager}, // You can replace this service with a custom shutdown service.
	}

	// This framework uses a third-party router (github.com/gorilla/mux). You can learn more about the router, to understand what kind of routes are supported.

	// Services that require initialization, at startup time, should have the name of their initialization function here.
	services_Init []func () = []func () {
		shutdown_Manager___Init, // The initialization function of service "shutdown_Manager".
	}
)
