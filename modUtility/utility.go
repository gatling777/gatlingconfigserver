package modUtility

import "fmt"

func Utility_initialize() error {
	err := Utility_log_initialize()
	if err != nil {
		fmt.Println("Utility log initialize error: ", err)
		return nil
	}
	err = config_Initialize()
	if err != nil {
		fmt.Println("Utility config initialize error: ", err)
		return nil
	}

	return nil
}
