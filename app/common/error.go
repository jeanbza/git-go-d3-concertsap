package common

import (
	"fmt"
	// "os"
)

func CheckError(err error) {
    if err != nil {
        ReportError(err)
        // os.Exit(1)
    }
}

func ReportError(err error) {
    fmt.Println("Fatal error caught by common.CheckError: ", err.Error())
}

func IssetInForm(form []string, index int) string {
    if len(form) < index+1 {
        return ""
    } else {
        return form[index]
    }
    
    return ""
}