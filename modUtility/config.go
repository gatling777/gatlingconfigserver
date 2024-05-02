package modUtility

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// x api key: for operator only;
const CG_Key_xApiKey = "XAPIKEY"
const CG_Key_HttpPort = "HTTPPORT"
const CG_Key_DBUrl = "DBURL"
const CG_Key_DBToken = "DBTOKEN"
const CG_Key_LocalDataPath = "LOCALDATAPATH"

const GD_Default_XApiKey = "aaXAPIKEYaa"
const GD_Default_HttpPort = 10000

var G_XApiKey = ""
var G_HttpPort = 0
var G_DBUrl = ""
var G_DBToken = ""
var G_AppPath = ""
var G_LocalDataPath = ""

func config_Initialize() error {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	G_AppPath = path[:index+1]

	G_XApiKey = getEnvKey(CG_Key_xApiKey)
	if G_XApiKey == "" {
		G_XApiKey = GD_Default_XApiKey
	}
	portStr := getEnvKey(CG_Key_HttpPort)
	if portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err == nil {
			G_HttpPort = port
		}
	}
	if G_HttpPort < 10 {
		G_HttpPort = GD_Default_HttpPort
	}
	G_DBUrl = getEnvKey(CG_Key_DBUrl)
	G_DBToken = getEnvKey(CG_Key_DBToken)

	G_LocalDataPath = getEnvKey(CG_Key_LocalDataPath)

	return nil
}

func getEnvKey(key string) string {
	portStr := os.Getenv(key)
	portStr = strings.TrimSpace(portStr)
	portStr = strings.Trim(portStr, "\r")
	portStr = strings.Trim(portStr, "\n")
	portStr = strings.Trim(portStr, "\r")

	return portStr
}
