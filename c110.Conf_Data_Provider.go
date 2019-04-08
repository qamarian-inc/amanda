package main

import (
	"errors"
	"fmt"
        viper_Interface "github.com/qamarian-inc/viper"
        "github.com/spf13/viper"
        "os"
)

/* This component makes the use of configuration files possible */

func init () { // Initializes this component.
	
	configuration, errX := viper_Interface.New_Viper (CONF_FILE, "yaml")

	if errX != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Loading configuration file: init () in c110_Conf_Data_Provider.go ()", errX.Error ()))
		os.Exit (1)
        }

        // Configuration is made gloabl, so as to become available to other functions in this component.
        conf_Data_Provider___configuration = configuration
}

func conf_Data_Provider (data_Name string) (string, error) { /* This function provides the value of the configuration data requested.

	If a configuration data is set in the configuration file, the value of the data and nil error, are returned.

	If a configuration data is not set, an empty string and error "conf_Data_Provider___DATA_NOT_SET", are returned. */

	// Error is returned if the data is not set, in the configuration file.
	if ! conf_Data_Provider___configuration.IsSet (data_Name) {
		return "", conf_Data_Provider___DATA_NOT_SET
	}

	return conf_Data_Provider___configuration.GetString (data_Name), nil
}

var (
	conf_Data_Provider___configuration *viper.Viper
	conf_Data_Provider___DATA_NOT_SET error = errors.New ("The configuration data is not set: conf_Data_Provider ()")
)
