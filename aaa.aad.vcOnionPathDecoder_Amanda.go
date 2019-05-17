package main

/* This virtual component decodes an onion path into its actual form. */

const iDecode_vcOnionPathDecoder_Amanda func (string) (string, error) /* This interface decodes an onion filepath.

	INPUT
	input 0: The onion filepath to be decoded.

	OUTPT
	outpt 0: The decoded form of input 0. On successful decoding, value would be the decoded form of input 0. On failed decoding, value would be an empty string.

	outpt 1: Any error that occurs during decoding. On successful decoding, value would be nil. On failed decoding, value would the error that occured. */
