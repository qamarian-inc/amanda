package main

/* This component decodes an onion filepath into its real form. If the real form of the filepath is a symlink, the symlink will be further evaluated into its extremely-real form. 

   To use this component, simply call its interface "iDecode_AAAAAE ()". */

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func iDecode_AAAAAE (onionPath string) (path string, err error) { /* This interface decodes an onion filepath, into its real form.

	In addition to decoding onion-formatted filepaths, if an onion path decodes to a symbolic link, this function further evaluates the symbolic link into its extremely-real form.

	EXPLANATION
	If an onion filepath (lets say "./file.ext") decodes to "/pathA/file.ext", and this filepath (/pathA/file.ext) is a symlink which points to "/pathB/file.ext" (another symlink), and "/pathB/file.ext" further points to "/pathC/file.ext" which is a real filepath, then the input of "./file.ext" would result into the output of "/pathC/file.ext".

	INPUT
	input 0: The onion filepath to be decoded.

	OUTPT
	outpt 0: The decoded form of input 0. On successful decoding, value would be the decoded form of input 0. On failed decoding, value would be an empty string.

	outpt 1: Any error that occurs during decoding. On successful decoding, value would be nil. On failed decoding, value would the error that occured. */

	// If a panic should occur, it is prevented from affecting other components. { ...
	err = errors.New ("iDecode_AAAAAE () paniced.") // This error will be returned if a panic should occur.

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
			errorMessage := fmt.Sprintf ("%s ---> \n Fetching directory of this program's file: iDecode_AAAAAE ()", errX.Error ())
			return "", errors.New (errorMessage)
		}

		pathSeparator := fmt.Sprintf ("%c", os.PathSeparator)

		onionPath = strings.Replace (onionPath, "*/", (filepath.Dir (programFileDir) + pathSeparator), 1)

	} else if strings.Index (onionPath, "./") == 0 {
	// If the filepath starts with "./", "./" is replaced with the present working directory.

		presentWorkingDir, errY := os.Getwd ()

		if errY != nil {
			errorMessage := fmt.Sprintf ("%s ---> \n Fetching present working directory: iDecode_AAAAAE ()", errY.Error ())
			return "", errors.New (errorMessage)
		}

		pathSeparator := fmt.Sprintf ("%c", os.PathSeparator)

		onionPath = strings.Replace (onionPath, "./", (presentWorkingDir + pathSeparator), 1)
	}

	// In case the real form of the onion filepath is a symlink, the symlink will be evaluated into its extreme real form.
	onionPath, errZ := filepath.EvalSymlinks (onionPath)

	if errZ != nil {
		errorMessage := fmt.Sprintf ("%s ---> \n Translating symlink into the actual file path: iDecode_AAAAAE ()", errZ.Error ())
		return "", errors.New (errorMessage)
	}

	return onionPath, nil
}
