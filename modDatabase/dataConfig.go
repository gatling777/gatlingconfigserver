/*
package data
should be in another package
*/
package modDatabase

import "fmt"

type CDataConfig struct {
	configMap map[string]string
}

var g_singleConfigManager *CDataConfig = &CDataConfig{configMap: make(map[string]string)}

func GetConfigManager() *CDataConfig {
	return g_singleConfigManager
}

func (pInst *CDataConfig) initialize() error {
	err := pInst.loadConfigData()
	if err != nil {
		return err
	}

	return nil
}

func (pInst *CDataConfig) loadConfigData() error {
	rows, err := getSingleTableConfig().Query()
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var name string
		var text string

		if err := rows.Scan(&name, &text); err != nil {
			fmt.Println("Error scanning row:", err)
			return err
		}

		pInst.configMap[name] = text
		//fmt.Println(project.ID, project.Name, project.Datetime1)
	}

	return nil
}
