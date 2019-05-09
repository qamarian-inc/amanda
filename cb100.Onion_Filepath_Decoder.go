package main

/* This component decodes an onion-formatted filepath into its real form. */

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func onion_Filepath_Decoder (onion_Path string) (string, error) { /* This function decodes an onion-formatted filepath, to its real form.

	In addition to decoding onion-formatted filepaths, if a filepath's real form is a symbolic link, this function will evaluate the symbolic link into its extreme real form.

	For example: if the onion filepath has the real form "/home/user/file.ext", and this filepath is a symlink that points to "/bin/user/file.ext", and this "/bin/user/file.ext" is another symlink and points to "/usr/user/file.ext" which is the actual file, then the output of this function (for this example) will be "/usr/user/file.ext".

	If any error is encountered, this function returns an empty string and the error. */

	// If the filepath starts with "*/", "*/" is replaced with the actual directory of the program's file (the directory where the executable file of the program is located).
	if strings.IndexOf (onion_Path, "*/") == 0 {
		program_File_Dir, errX := os.Executable ()

		if errX != nil {
			error_Message := fmt.Sprintf ("%s ---> \n Fetching directory of this program's file: onion_File_Path_To_OS_File_Path ()", errX.Error ())
			return "", errors.New (error_Message)
		}

		onion_Path = strings.Replace (onion_Path, "*/", program_File_Dir, 1)

	} else if strings.IndexOf (onion_Path, "*/") == 0 {
	// If the filepath starts with "./", "./" is replaced with the present working directory.

		present_Working_Dir, errY := os.Gwd ()

		if errY != nil {
			error_Message := fmt.Sprintf ("%s ---> \n Fetching present working directory: onion_File_Path_To_OS_File_Path ()", errY.Error ())
			return "", errors.New (error_Message)
		}

		onion_Path = strings.Replace (onion_Path, "*/", present_Working_Dir, 1)
	}

	// In case the real form of the onion filepath is a symlink, the symlink will be evaluated into its extreme real form.
	onion_Path, errA = filepath.EvalSymlinks (onion_Path)

	if errA != nil {
		error_Message := fmt.Sprintf ("%s ---> \n Translating symlink into the actual file path: onion_File_Path_To_OS_File_Path ()", errZ.Error ())
		return "", errors.New (error_Message)
	}

	return onion_Path, nil
}
