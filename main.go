package main

import (
	"configServer/modDatabase"
	"configServer/modHttp"
	"configServer/modUtility"
	"fmt"
)

func main() {
	err := modUtility.Utility_initialize()
	if err != nil {
		fmt.Println("utility init failed: ", err)
		return
	}

	err = modDatabase.DB_Initialize()
	if err != nil {
		fmt.Println("database init failed: ", err)
		return
	} else {
		fmt.Println("database init successful")
	}

	err = modHttp.Http_Initialize()
	if err != nil {
		fmt.Println("http init failed: ", err)
		return
	}

	modHttp.Http_Start()

}
