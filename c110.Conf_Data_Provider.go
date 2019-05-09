package main

import (
	"errors"
	"fmt"
        viper_Interface "github.com/qamarian-inc/viper"
        "github.com/spf13/viper"
        "os"
)

/* This component is responsible for providing the data in the configuration file, to other components. */

func init () { // Initializes this component. The initialization basically means caching the configuration file.
	
	// Translating the configuration file path from its onion form to its real form.
	conf_File, errD := onion_Filepath_Decoder (CONF_FILE)

	if errD != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Translating the configuration filepath from its onion form to its real form: init () in c110_Conf_Data_Provider.go ()", errD.Error ()))
		os.Exit (1)
        }

        // Loading the configuration file.
	configuration, errX := viper_Interface.New_Viper (CONF_FILE, "yaml")

	if errX != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Loading configuration file: init () in c110_Conf_Data_Provider.go ()", errX.Error ()))
		os.Exit (1)
        }

        // Configuration is made gloabl, so as to become available to other functions in this component.
        conf_Data_Provider___configuration = configuration
}

func conf_Data_Provider (data_Name string) (string, error) { /* This function provides the value of scalar data in the configuration file.

	If the data is set in the configuration file, the value of the data and nil error, are returned.

	If the data is not set, an empty string and error "conf_Data_Provider___DATA_NOT_SET", are returned. */

	// Error is returned if the data is not set, in the configuration file.
	if ! conf_Data_Provider___configuration.IsSet (data_Name) {
		return "", conf_Data_Provider___DATA_NOT_SET
	}

	return conf_Data_Provider___configuration.GetString (data_Name), nil
}

func conf_Data_Provider___Array_Data (data_Name string) (string, error) { /* This function provides the value of array data in the configuration file.

	If the configuration data is set in the configuration file, the value of the data and nil error, are returned.

	If the configuration data is not set, an empty string and error "conf_Data_Provider___DATA_NOT_SET", are returned. */

	// Error is returned if the data is not set, in the configuration file.
	if ! conf_Data_Provider___configuration.IsSet (data_Name) {
		return "", conf_Data_Provider___DATA_NOT_SET
	}

	return conf_Data_Provider___configuration.GetStringSlice (data_Name), nil
}

func conf_Data_Provider___Map_Data (data_Name string) (string, error) { /* This function provides the value of hash map data in the configuration file.

	If the configuration data is set in the configuration file, the value of the data and nil error, are returned.

	If the configuration data is not set, an empty string and error "conf_Data_Provider___DATA_NOT_SET", are returned. */

	// Error is returned if the data is not set, in the configuration file.
	if ! conf_Data_Provider___configuration.IsSet (data_Name) {
		return "", conf_Data_Provider___DATA_NOT_SET
	}

	return conf_Data_Provider___configuration.GetStringMapString (data_Name), nil
}

var (
	conf_Data_Provider___configuration *viper.Viper // Configuration data cache.
	conf_Data_Provider___DATA_NOT_SET error = errors.New ("The configuration data requested is not set: conf_Data_Provider ()")
)
