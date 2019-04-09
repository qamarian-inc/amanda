package main

import (
	"errors"
	"fmt"
	"github.com/blue-jay/core/email"
	"os"
	"runtime"
	"strconv"
)

/* This component helps sends distress email alerts to the admin of the software. On startup, the alert raiser is started. To ask alert raiser to send an email, to the admin, use function alert_Raiser___Raise_Alert ().

You can shutdown the alert raiser, using alert_Raiser___Shutdown ().
*/

func init () { // This function fetches all data needed by alert_Raiser (), then starts alert_Raiser () itself.

	// Fetching the email address alerts should be sent to.
	alert_Destination_Email := os.Getenv (ALERT_DESTINATION_EMAIL)
	if alert_Destination_Email == "" {
		output (fmt.Sprintf ("Startup Error: Enviromental variable '%s' (ALERT_DESTINATION_EMAIL) is not set: init () in c130_alert_Raiser.go", ALERT_DESTINATION_EMAIL))
		os.Exit (1)
	}

	// Fetching the address of the email server used to send alerts.
	alert_Email_Server_Addr := os.Getenv (ALERT_RAISER_EMAIL_SERVER_ADDR)
	if alert_Email_Server_Addr == "" {
		output (fmt.Sprintf ("Startup Error: Enviromental variable '%s' (ALERT_RAISER_EMAIL_SERVER_ADDR) is not set: init () in c130_alert_Raiser.go", ALERT_RAISER_EMAIL_SERVER_ADDR))
		os.Exit (1)
	}

	// Fetching the port of the email server used to send alerts. { ...
	alert_Email_Server_Port := os.Getenv (ALERT_RAISER_EMAIL_SERVER_PORT)
	if alert_Email_Server_Port == "" {
		output (fmt.Sprintf ("Startup Error: Enviromental variable '%s' (ALERT_RAISER_EMAIL_SERVER_PORT) is not set: init () in c130_alert_Raiser.go", ALERT_RAISER_EMAIL_SERVER_PORT))
		os.Exit (1)
	}
	server_Port_As_Int, errX := strconv.Atoi (alert_Email_Server_Port);
	if errX != nil {
		output (fmt.Sprintf ("Startup Error: Environmental variable '%s' (ALERT_RAISER_EMAIL_SERVER_PORT) is not a number, number required: init () in c130_alert_Raiser.go", ALERT_RAISER_EMAIL_SERVER_PORT))
		os.Exit (1)
	}
	// ... }

	// Fetching the email address alert_Raiser () uses to send alerts. { ...
	alert_Email_Username := os.Getenv (ALERT_RAISER_EMAIL_ADDR_ENV_VAR)
	if alert_Email_Username == "" {
		output (fmt.Sprintf ("Startup Error: Enviromental variable '%s' (ALERT_RAISER_EMAIL_ADDR_ENV_VAR) is not set: init () in c130_alert_Raiser.go", ALERT_RAISER_EMAIL_ADDR_ENV_VAR))
		os.Exit (1)
	}
	// ... }

	// Fetching the password of the email address alert_Raiser () uses to send alerts. { ...
	alert_Email_Password := os.Getenv (ALERT_RAISER_EMAIL_PASS_ENV_VAR)
	if alert_Email_Password == "" {
		output (fmt.Sprintf ("Startup Error: Enviromental variable '%s' (ALERT_RAISER_EMAIL_PASS_ENV_VAR) is not set: init () in c130_alert_Raiser.go", ALERT_RAISER_EMAIL_PASS_ENV_VAR))
		os.Exit (1)
	}
	// ... }

	// Starting alert_Raiser ().
	go alert_Raiser (alert_Destination_Email, alert_Email_Server_Addr, server_Port_As_Int, alert_Email_Username, alert_Email_Password)
}

func alert_Raiser (alert_Destination_Email, alert_Email_Server_Addr string, alert_Email_Server_Port int, alert_Email_Username, alert_Email_Password string) { // This function sends email alerts to the admin email provided. It runs as a daemon (goroutine), and alert_Raiser___Raise_Alert () can be used to communicate with it.


	// In case any panic occurs, the panic is logged.
	defer func () {
		panic_Reason := recover ()
		if panic_Reason != nil {
			defer logger (`alert_Raiser () paniced.`)
		}
	} ()

	defer output (`State: Alert raiser "alert_Raiser ()" is down.`)

	// Creating channel needed to receive new alerts meant to be sent.
	alert_Raiser___alert_Channel = make (chan *alert)
	// Creating channel used to signal shutdown.
	alert_Raiser___shutdown_Channel = make (chan bool)

	// Creating info needed to send alert emails.
	mailing_Info := email.Info {
		Hostname: alert_Email_Server_Addr,
		Port:     alert_Email_Server_Port,
		Username: alert_Email_Username,
		Password: alert_Email_Password,
		From:     alert_Email_Username,
	}

	// Continously sends email alert, until it has been asked to shutdown.
	for {
		select {

		// Sends email.
		case alert, ok := <- alert_Raiser___alert_Channel:

			if ok == true {
				mail_Subject := fmt.Sprintf ("%s (%s) Alert", SOFTWARE_ID, SOFTWARE_NAME)

				// Sending alert.
				errY := mailing_Info.Send (alert_Destination_Email, mail_Subject, alert.alert_Message)

				if errY != nil {
					error_Message := fmt.Sprintf ("%s ---> \n Sending alert: init () in c130_Alert_Raiser.go in %s", errY.Error (), SOFTWARE_ID)
					alert.err = errors.New (error_Message)
				}
			}

		// Shutdown.
		case _, _ = <- alert_Raiser___shutdown_Channel:

			close (alert_Raiser___alert_Channel)
			close (alert_Raiser___shutdown_Channel)
			return
		}

		runtime.Gosched ()
	}
}

func alert_Raiser___Raise_Alert (alert_Message string) (error) {
	// In case any panic occurs, the panic is prevented from getting out.
	defer func () {
		recover ()
	} ()

	new_Alert := alert {
		alert_Message: alert_Message,
		err: nil,
		status: false,
	}

	// Sending alert to alert_Raiser ().
	alert_Raiser___alert_Channel <- &new_Alert

	return new_Alert.err
}

func alert_Raiser___Shutdown () {
	// In case any panic occurs, the panic is prevented from getting out.
	defer func () {
		recover ()
	} ()

	// Signaling shutdown, to alert_Raiser ().
	alert_Raiser___shutdown_Channel <- true
}

type alert struct { // A data of this type represents an alert that needs to be sent to the admin.
	alert_Message string
	err error // Any error that occurs while trying to send an alert's message. If no error occurs, value will be nil.
	status bool // Value will be true if an alert's message has been successfully sent, otherwise, it will be false.
}

var (
	alert_Raiser___alert_Channel chan *alert // Any alert sent to this channel, will be emailed, by  alert_Raiser ().
	alert_Raiser___shutdown_Channel chan bool // alert_Raiser () will shutdown whenever alert_Raiser___Shutdown () sends a boolean value to this channel.
)
