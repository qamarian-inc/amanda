package main

import "time"

const (
	// Modify the following constants as appropriate.

	SOFTWARE_NAME string = "Amanda Ware"// The name of the software.
	SOFTWARE_ID string = "ware_X" // The ID of the software.
	CONF_FILE string = "zz.Conf.yml" // The filepath of the configuration file. Only YAML files are supported, and the extension of the file must be ".yaml".

	LOG_FILE_DIR_ENV_VAR string = "AMANDA_LOG_FILES_PATH" // The environmental variable name of the directory where the log file of the software should be saved.
	ALERT_RAISER_EMAIL_ADDR_ENV_VAR string = "AMANDA_ALERT_EMAIL_USERNAME" // The environmental variable name of the email address alert raiser can use to send emails to the admin.
	ALERT_RAISER_EMAIL_PASS_ENV_VAR string = "AMANDA_ALERT_EMAIL_PASSWORD" // The environmental variable name of the password of the email specified in 'ALERT_RAISER_EMAIL_ADDR_ENV_VAR'.
	ALERT_RAISER_EMAIL_SERVER_ADDR string = "AMANDA_ALERT_EMAIL_SERVER_ADDR" // The environmental variable name of the IP address/domain name of the email server of the email specified in 'ALERT_RAISER_EMAIL_ADDR_ENV_VAR'.
	ALERT_RAISER_EMAIL_SERVER_PORT string = "AMANDA_ALERT_EMAIL_SERVER_PORT" // The environmental variable name of the port of the email server of the email specified in 'ALERT_RAISER_EMAIL_ADDR_ENV_VAR'.
	ALERT_DESTINATION_EMAIL string = "AMANDA_ALERT_DESTINATION_EMAIL" // The environmental variable name of the email that alerts should be sent to.

	TLS_CERT_FILEPATH string = "AMANDA_TLS_CERT" // The environmental variable name of the filepath of a TLS certificate that can be used.
	PRIV_KEY_FILEPATH string = "AMANDA_PRIV_KEY" // The environmental variable name of the filepath of a TLS private key that can be used.

	SERVER_NET_IO_TIMEOUT time.Duration = time.Second * 60 // The maximum duration allowed for input-output tasks of the server.
	SERVER_MAX_HEADER_SIZE int = 1 << 20 // The maximum HTTPS header size allowed on the server (in bytes).
)

var (
	// Services that the software provides should be added to this slice.
	software_Service []service = []service {
		service {"/shutdown/{admin_Code}", shutdown_Manager}, // You can replace this service with a custom shutdown service.
	}

	// This framework uses a third-party router (github.com/gorilla/mux). You can learn more about the router, to understand what type of routes are supported.
)
