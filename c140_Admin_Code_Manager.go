package main

import (
        "github.com/qamarian-inc/strings"
        "os"
)

func init () { // This function generates a new admin code, and saves the code in a file named the ID of the software.

	// Generating code for admin services. { ...
        admin_Code, errX :=  strings.CryptoSafe_Unique_Random_32 ()
        if errX != nil {
                output ("Startup Error: %s ---> \n Generating admin code: init () in c140_Admin_Code_Manager.go" + errX.Error ())
                os.Exit (1)
        }

        // Making code a global data (visible to other functions in this component).
        admin_Code_Manager___admin_Code = admin_Code
        // ... }

        // Saving admin code in file. { ...
        code_File, err2 := os.OpenFile (SOFTWARE_ID, os.O_CREATE|os.O_WRONLY, 0330) 
        if err2 != nil {
                output ("Startup Error: %s ---> \n Saving admin code in file: init () in c140_Admin_Code_Manager.go ()" + err2.Error ())
                os.Exit (1)
        }

        _, err3 := code_File.WriteString (admin_Code_Manager___admin_Code)
        if err3 != nil {
                output ("Startup Error: %s ---> \n Saving admin code in file: init () in c140_Admin_Code_Manager.go ()" + err3.Error ())
                os.Exit (1)
        }

        code_File.Close ()
        // ... }
}

func admin_Code_Manager () (string) {
	return admin_Code_Manager___admin_Code
}

var (
        admin_Code_Manager___admin_Code string
)
