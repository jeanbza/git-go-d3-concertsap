package common

import (
    "log"
    "runtime"
)

func CheckError(err error) {
    if err != nil {
        ReportError(err)
        // os.Exit(1)
    }
}

func ReportError(err error) {
    var stack [4096]byte
    runtime.Stack(stack[:], false)
    log.Printf("%q\n%s\n", err, stack[:])
}

func IssetInForm(form []string, index int) string {
    if len(form) < index+1 {
        return ""
    } else {
        return form[index]
    }
    
    return ""
}