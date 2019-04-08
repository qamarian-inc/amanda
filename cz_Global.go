package main

const (
	// Modify the following constants as appropriate.

	SOFTWARE_NAME string = "Amanda Ware"// The name of the software.
	SOFTWARE_ID string = "ware_X" // The ID of the software.
	CONF_FILE string = "conf.yaml" // The filepath of the configuration file. Only YAML files supported, and extension of the file must be ".yaml".

	LOG_FILE_DIR_ENV_VAR string = "" // The environmental variable name of the directory where the log file of the software should be saved.
	ALERT_RAISER_EMAIL_ADDR_ENV_VAR string = "" // The environmental variable name of the email address alert raiser can use to send emails to the admin.
	ALERT_RAISER_EMAIL_PASS_ENV_VAR string = "" // The environmental variable name of the password of the email specified in 'ALERT_RAISER_EMAIL_ADDR_ENV_VAR'.
	ALERT_RAISER_EMAIL_SERVER_ADDR string = "" // The environmental variable name of the IP address/domain name of the email server of the email specified in 'ALERT_RAISER_EMAIL_ADDR_ENV_VAR'.
	ALERT_RAISER_EMAIL_SERVER_PORT string = "" // The environmental variable name of the port of the email server of the email specified in 'ALERT_RAISER_EMAIL_ADDR_ENV_VAR'.
	ALERT_DESTINATION_EMAIL string = "" // The environmental variable name of the email that alerts should be sent to.

	TLS_CERT_FILEPATH string = "" // The environmental variable name of the filepath of a TLS certificate that can be used.
	PRIV_KEY_FILEPATH string = "" // The environmental variable name of the filepath of a TLS private key that can be used.
)

var (
	// Services that the software provides should be added.
	software_Service []service

	// Admin services of the software should be added.
	admin_Service []service = []service {
		service {"/shutdown/{admin_Code}", shutdown_Manager}, // You can replace this service with a custom shutdown service.
	}
)

/* The difference between software services and admin services is that admin services are meant to be provided to only the admin, and this require providing a code to the admin services, to prove that a user is truly an admin. See file "" to learn more about admin code.

This framework uses a third-party router (github.com/gorilla/mux). You can learn more about the router, to understand how what type of service paths are supported.

*/