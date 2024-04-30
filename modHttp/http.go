package modHttp

func Http_Initialize() error {
	showIP()
	return getSingleChi().Initialize()
}
func Http_Start() error {
	return g_singleChi.start()
}
