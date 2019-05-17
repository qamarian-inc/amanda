package main

/* This virtual component makes possible the use of configuration file. Its interfaces can be used to fetch configuration data from a configuration file. The interface to use among the three, depends on the kind of data you're trying to fetch. */

const (
	iScalarData_vcConfDataProvider_Amanda func () (string, error) // To fetch a scalar configuration data (number or string) use this interface.
	iSliceData_vcConfDataProvider_Amanda func () ([]string, error) // To fetch an array configuration data use this interface.
	iMapData_vcConfDataProvider_Amanda func () (map[string]string, error) // To fetch a map configuration data use this interface.
)
