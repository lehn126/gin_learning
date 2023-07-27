package app

import (
	"fmt"
	"gin_app/app/config"
	"gin_app/app/util/conf"
	"gin_app/app/util/file"
	"os"
	"strings"
)

type RuntimeType struct {
	RootPath      string
	ResourcesDir  string
	Profile       string
	AppConfigPath string
}

var Runtime = RuntimeType{}

// Get the full config file path by file Name (without postfix) and file postfix
func (rt *RuntimeType) GetFullConfFilePath(fileName string, postfix string, useProfile bool) string {
	arr := []string{rt.RootPath, rt.ResourcesDir} // Get file in resources folder

	if useProfile && rt.Profile != "" {
		arr = append(arr, fmt.Sprintf("%s-%s.%s", fileName, rt.Profile, postfix))
	} else {
		arr = append(arr, fmt.Sprintf("%s.%s", fileName, postfix))
	}
	path := strings.Join(arr, string(os.PathSeparator))
	if !file.IsExist(path) {
		panic(fmt.Sprintf("config file %s not exist: %s", fileName, path))
	}
	return path
}

// Use yaml file by default
func (rt *RuntimeType) LoadConfFile(fileName string, confInstancePointer any) {
	rt.AppConfigPath = rt.GetFullConfFilePath(fileName, "yml", true)
	conf.LoadConfFromYml(rt.AppConfigPath, confInstancePointer)
}

// Default directory used to store configuration files
const DEFAULT_RESOURCES_DIR = "resources"

// Default app config file name (without postfix)
const DEFAULT_CONF_APP_CONFIG string = "app-config"

func Init() {
	// Load app-config.yml
	Runtime.LoadConfFile(DEFAULT_CONF_APP_CONFIG, &config.APP_CONFIG)
	Runtime.LoadConfFile(DEFAULT_CONF_APP_CONFIG, &config.DB_CONFIG)
}
