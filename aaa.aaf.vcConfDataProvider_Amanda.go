package main

/* This virtual component makes possible the use of configuration file. Its interfaces can be used to fetch configuration data from a configuration file. The interface to use among the three, depends on the kind of data you're trying to fetch. */

const (
	iScalarData_vcConfDataProvider_Amanda func (string) (string, error) /* To fetch a scalar configuration data (number or string) use this interface.

	INPUT
	input 0: The name of the data to be fetched.

	OUTPT
	outpt 0: The value of the data asked to be fetched. Value would be an empty string if any error should occur during the fetch.

	outpt 1: Any error that occurs during the fetch. On successful fetch, value would be nil. On failed fetch, value would the error that occured. If the data is not set, value would be an error. */
	
	iSliceData_vcConfDataProvider_Amanda func (string) ([]string, error) /* To fetch an array configuration data use this interface. The interface requires one parametre which should be the name of the data to be fetched.

	INPUT
	input 0: The name of the data to be fetched.

	OUTPT
	outpt 0: The value of the data asked to be fetched. Value would be an empty slice if any error should occur during the fetch.
	
	outpt 1: Any error that occurs during the fetch. On successful fetch, value would be nil. On failed fetch, value would the error that occured. If the data is not set, value would be an error. */

	iMapData_vcConfDataProvider_Amanda func (string) (map[string]string, error) /* To fetch a map configuration data use this interface. The interface requires one parametre which should be the name of the data to be fetched.

	INPUT
	input 0: The name of the data to be fetched.

	OUTPT
	outpt 0: The value of the data asked to be fetched. Note indeces of map would all be in lowercase. Value would be an empty map if any error should occur during the fetch.
	
	outpt 1: Any error that occurs during the fetch. On successful fetch, value would be nil. On failed fetch, value would the error that occured. If the data is not set, value would be an error. */
)
