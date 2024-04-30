package modUtility

import "fmt"

func Utility_log_initialize() error {
	return nil
}

func Utility_log(log string, args ...interface{}) {
	fmt.Println(log, args)
}
func Utility_error(log string, args ...interface{}) {
	fmt.Println(log, args)
}
