package modUtility

import "fmt"

func Utility_initialize() error {
    err := Utility_log_initialize()
    if err != nil {
        fmt.Prtinln("Utility log initialize error: ", err)
        return nil
    }


    return nil
}