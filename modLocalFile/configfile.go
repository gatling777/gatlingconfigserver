package modLocalFile

type cConfigFile struct {
}

var g_singleConfigFile *cConfigFile = &cConfigFile{}

func getSingleConfigFile() *cConfigFile {
	return g_singleConfigFile
}

func (pInst *cConfigFile) initialize() error {
	return nil
}
