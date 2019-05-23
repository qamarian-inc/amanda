package main

/* This component makes possible the use of configuration file, in your app. The component uses the YAML syntax for its configuration file.

   DEPENDENCIES
	Virtual component "aaaaad" (Onion path decoder)

   USAGE NOTE
   	1. Ensure the conf file specified in "dCONFILE_AAAAAG", is available on the machine where the app would run.

   	2. Any data set in your conf file, can be fetched using any of interfaces:
   		iScalarData_AAAAAG (),
   		iSliceData_AAAAAG (), and
   		iMapData_AAAAAG ().

   		The interface to use among the three, would depend on what kind of conf data you're trying to fetch.

   	3. This component comes with a default configuration file known as "aaaaag.dConf.yml".
*/

import (
	"errors"
	"fmt"
        viper_Interface "github.com/qamarian-inc/viper"
        "github.com/spf13/viper"
        "os"
)

const dCONFILE_AAAAAG string = "./aaaaag.dConf.yml" // You can modify this constant if your app's configuration file wouldn't be file "aaaaag.dConf.yml" or if file "aaaaag.dConf.yml" wouldn't be present in the app's current working directory. Filepath format supported: onion.

func init () { // The initialization basically means caching the configuration file.
	
	// Decoding the configuration filepath from its onion form into its real form.
	conFilepath, errD := iDecode_AAAAAD (dCONFILE_AAAAAG)

	if errD != nil {
		iOutput_AAAAAB (fmt.Sprintf ("Startup Error: %s ---> \n Decoding the configuration filepath from its onion form to its real form: init () in aaaaag.ConfDataProvider.go", errD.Error ()))
		os.Exit (1)
        }

        // Loading the configuration file.
	conf, errX := viper_Interface.New_Viper (conFilepath, "yaml")

	if errX != nil {
		iOutput_AAAAAB (fmt.Sprintf ("Startup Error: %s ---> \n Loading configuration file: init () in aaaaag.ConfDataProvider.go", errX.Error ()))
		os.Exit (1)
        }

        // Configuration is made gloabl, so as to become available to other functions in this component.
        dConf_AAAAAG = conf

        // Indicating that this component is now ready to provide its services.
        dCompAvailable_AAAAAG = true
}

func iScalarData_AAAAAG (dataName string) (string, error) { /* This interface helps fetch the value of a scalar data, from the configuration file.

	INPUT
	input 0: The name of the data to be fetched.

	OUTPT
	outpt 0: The value of the data. Value would be an empty string if any error should occur during the fetch.
	
	outpt 1: Any error that occurs during the fetch. On successful fetch, value would be nil. On failed fetch, value would the error that occured. If the data is not set, value would be error "eDATANOTSET_AAAAAG". If the component has just started up, and it is yet to be available, value would be error "eCOMPNOTAVAILABLE_AAAAAG". */

	if dCompAvailable_AAAAAG == false {
		return "", eCOMPNOTAVAILABLE_AAAAAG
	}

	if ! dConf_AAAAAG.IsSet (dataName) {
		return "", eDATANOTSET_AAAAAG
	}

	return dConf_AAAAAG.GetString (dataName), nil
}

func iSliceData_AAAAAG (dataName string) (data []string, err error) { /* This interface helps fetch the value of an array data, from the configuration file.

	INPUT
	input 0: The name of the data to be fetched.

	OUTPT
	outpt 0: The value of the data. Value would be an empty slice if any error should occur during the fetch.
	
	outpt 1: Any error that occurs during the fetch. On successful fetch, value would be nil. On failed fetch, value would the error that occured. If the data is not set, value would be error "eDATANOTSET_AAAAAG". If the component has just started up, and it is yet to be available, value would be error "eCOMPNOTAVAILABLE_AAAAAG". */

	if dCompAvailable_AAAAAG == false {
		return data, eCOMPNOTAVAILABLE_AAAAAG
	}

	if ! dConf_AAAAAG.IsSet (dataName) {
		return []string {}, eDATANOTSET_AAAAAG
	}

	return dConf_AAAAAG.GetStringSlice (dataName), nil
}

func iMapData_AAAAAG (dataName string) (data map[string]string, err error) { /* This interface helps fetch the value of a hash map data, from the configuration file.

	INPUT
	input 0: The name of the data to be fetched.

	OUTPT
	outpt 0: The value of the data. Value would be an empty map if any error should occur during the fetch.
	
	outpt 1: Any error that occurs during the fetch. On successful fetch, value would be nil. On failed fetch, value would the error that occured. If the data is not set, value would be error "eDATANOTSET_AAAAAG". If the component has just started up, and it is yet to be available, value would be error "eCOMPNOTAVAILABLE_AAAAAG". */

	if dCompAvailable_AAAAAG == false {
		return data, eCOMPNOTAVAILABLE_AAAAAG
	}

	if ! dConf_AAAAAG.IsSet (dataName) {
		return data, eDATANOTSET_AAAAAG
	}

	return dConf_AAAAAG.GetStringMapString (dataName), nil
}

var (
	dConf_AAAAAG *viper.Viper // Configuration data cache. It would be initialized by init ().
	dCompAvailable_AAAAAG bool = false // This data signifies if this component is availble to provide its service or not.

	eCOMPNOTAVAILABLE_AAAAAG error = errors.New ("This component is not available yet!")
	eDATANOTSET_AAAAAG error = errors.New ("The configuration data requested is not set!")
)
