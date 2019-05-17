package main

/* This component makes possible the use of configuration file, in your app. The component uses the YAML syntax for its configuration file.

   USAGE
   	1. Ensure the conf file specified in "CONFILE_ConfDataProvider_Amanda", is available on the machine where the app would run.

   	2. Any data set in your conf file, can be fetched using any of interfaces:
   		iScalarData_ConfDataProvider_Amanda (),
   		iSliceData_ConfDataProvider_Amanda (), and
   		iMapData_ConfDataProvider_Amanda ().

   		The interface to use among the three, would depend on what kind of conf data you're trying to fetch.

   		This component comes with a default configuration file known as "aaa.aag.dConf.yml".

   DEPENDENCIES
	Virtual component aaa.aad (Onion path decoder)
*/

import (
	"errors"
	"fmt"
        viper_Interface "github.com/qamarian-inc/viper"
        "github.com/spf13/viper"
        "os"
)

const CONFILE_ConfDataProvider_Amanda string = "./aaa.aag.dConf.yml" // You can modify this constant if your app's configuration file wouldn't be file "aaa.aag.dConf.yml" or if file "aaa.aag.dConf.yml" wouldn't be present in the app's current working directory. Filepath format supported: onion.

func init () { // Initializes this component. The initialization basically means caching the configuration file.
	
	// Decoding the configuration filepath from its onion form into its real form.
	CONFILE_ConfDataProvider_Amanda, errD := iDecode_vcOnionPathDecoder_Amanda (CONFILE_ConfDataProvider_Amanda)

	if errD != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Decoding the configuration filepath from its onion form to its real form: init () in aaa.aag.ConfDataProvider_Amanda.go", errD.Error ()))
		os.Exit (1)
        }

        // Loading the configuration file.
	conf, errX := viper_Interface.New_Viper (CONFILE_ConfDataProvider_Amanda, "yaml")

	if errX != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Loading configuration file: init () in aaa.aag.ConfDataProvider_Amanda.go", errX.Error ()))
		os.Exit (1)
        }

        // Configuration is made gloabl, so as to become available to other functions in this component.
        dConf_ConfDataProvider_Amanda = conf
}

func iScalarData_ConfDataProvider_Amanda (dataName string) (string, error) { /* This interface helps fetch the value of a scalar data, from the configuration file.

	If the data is set in the configuration file, the value of the data and nil error, are returned.

	If the data is not set, an empty string and error "eDATANOTSET_ConfDataProvider_Amanda", are returned. */

	if ! dConf_ConfDataProvider_Amanda.IsSet (dataName) {
		return "", eDATANOTSET_ConfDataProvider_Amanda
	}

	return dConf_ConfDataProvider_Amanda.GetString (dataName), nil
}

func iSliceData_ConfDataProvider_Amanda (dataName string) (string, error) { /* This interface helps fetch the value of an array data, from the configuration file.

	If the configuration data is set in the configuration file, the value of the data and nil error, are returned.

	If the configuration data is not set, an empty string and error "eDATANOTSET_ConfDataProvider_Amanda", are returned. */

	if ! dConf_ConfDataProvider_Amanda.IsSet (dataName) {
		return "", eDATANOTSET_ConfDataProvider_Amanda
	}

	return dConf_ConfDataProvider_Amanda.GetStringSlice (dataName), nil
}

func iMapData_ConfDataProvider_Amanda (dataName string) (string, error) { /* This interface helps fetch the value of a hash map data, from the configuration file.

	If the configuration data is set in the configuration file, the value of the data and nil error, are returned.

	If the configuration data is not set, an empty string and error "eDATANOTSET_ConfDataProvider_Amanda", are returned. */

	if ! dConf_ConfDataProvider_Amanda.IsSet (dataName) {
		return "", eDATANOTSET_ConfDataProvider_Amanda
	}

	return dConf_ConfDataProvider_Amanda.GetStringMapString (dataName), nil
}

var (
	dConf_ConfDataProvider_Amanda *viper.Viper // Configuration data cache. It would be initialized by init ().
	eDATANOTSET_ConfDataProvider_Amanda error = errors.New ("The configuration data requested is not set: comp aaa.aag (ConfDataProvider_Amanda)")
)
