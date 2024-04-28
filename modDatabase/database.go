package modDatabase

func DB_Initialize() error {
	err := getSingleTableConfig().initialize()
	if err != nil {
		return err
	}

	err = GetConfigManager().initialize()

	return err
}
