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

func getDataConfig() *CDataConfig {
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

func (pInst *CDataConfig) getConfig(key string) string {
	data, exists := pInst.configMap[key]
	if !exists {
		return ""
	}
	return data
}

func (pInst *CDataConfig) inputConfig(key, text string) error {
	var bExists = false
	var err error = nil
	_, bExists = pInst.configMap[key]

	if bExists {
		err = getSingleTableConfig().Update(key, text)
	} else {
		err = getSingleTableConfig().Insert(key, text)
	}

	if err != nil {
		return err
	}

	pInst.configMap[key] = text
	return nil
}
