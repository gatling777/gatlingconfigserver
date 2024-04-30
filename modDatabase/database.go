package modDatabase

func DB_Initialize() error {
	err := getSingleTableConfig().initialize()
	if err != nil {
		return err
	}

	err = getDataConfig().initialize()

	return err
}

func DB_GetConfig(key string) string {
	return getDataConfig().getConfig(key)
}
func DB_InputConfig(key, text string) error {
	return getDataConfig().inputConfig(key, text)
}
