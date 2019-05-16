package main

/* This component decodes an onion-formatted filepath into its real form. */

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func iDecode_OnionPathDecoder_Amanda (onionPath string) (path string, err error) { /* This interface decodes an onion-formatted filepath, into its real form.

	In addition to decoding onion-formatted filepaths, if a filepath's real form is a symbolic link, this function will evaluate the symbolic link into its extreme real form.

	For example: if the onion filepath has the real form "/home/user/file.ext", and this filepath is a symlink that points to "/bin/user/file.ext", and this "/bin/user/file.ext" is another symlink and points to "/usr/user/file.ext" which is the actual file, then the output of this function (for this example) will be "/usr/user/file.ext".

	If any error is encountered, this function returns an empty string and the error. 
	*/


	// If a panic should occur, it is prevented from affecting other components. { ...
	err = errors.New ("iDecode_OnionPathDecoder_Amanda () paniced.") // This error will be returned if a panic should occur.

	defer func () {
		panicReason := recover ()

		if panicReason != nil {
			return
		}
	} ()
	// ... }

	// If the filepath starts with "*/", "*/" is replaced with the actual directory of the program's file (the directory where the executable file of the program is located).
	if strings.Index (onionPath, "*/") == 0 {
		programFileDir, errX := os.Executable ()

		if errX != nil {
			errorMessage := fmt.Sprintf ("%s ---> \n Fetching directory of this program's file: iDecode_OnionPathDecoder_Amanda ()", errX.Error ())
			return "", errors.New (errorMessage)
		}

		onionPath = strings.Replace (onionPath, "*/", programFileDir, 1)

	} else if strings.Index (onionPath, "*/") == 0 {
	// If the filepath starts with "./", "./" is replaced with the present working directory.

		presentWorkingDir, errY := os.Getwd ()

		if errY != nil {
			errorMessage := fmt.Sprintf ("%s ---> \n Fetching present working directory: iDecode_OnionPathDecoder_Amanda ()", errY.Error ())
			return "", errors.New (errorMessage)
		}

		onionPath = strings.Replace (onionPath, "*/", presentWorkingDir, 1)
	}

	// In case the real form of the onion filepath is a symlink, the symlink will be evaluated into its extreme real form.
	onionPath, errZ := filepath.EvalSymlinks (onionPath)

	if errZ != nil {
		errorMessage := fmt.Sprintf ("%s ---> \n Translating symlink into the actual file path: iDecode_OnionPathDecoder_Amanda ()", errZ.Error ())
		return "", errors.New (errorMessage)
	}

	return onionPath, nil
}
