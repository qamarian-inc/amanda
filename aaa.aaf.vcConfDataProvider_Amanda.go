package main

/* This virtual component can help provide configuration data of your app. */

const (
	iScalarData_vcConfDataProvider_Amanda func () (string, error)
	iSliceData_vcConfDataProvider_Amanda func () ([]string, error)
	iMapData_vcConfDataProvider_Amanda func () (map[string]string, error)
)
