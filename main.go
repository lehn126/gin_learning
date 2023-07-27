package main

import (
	"flag"
	"fmt"
	"gin_app/app"
	"gin_app/app/api"
	"gin_app/app/config"
	"gin_app/app/util/db"
	"gin_app/app/util/file"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

// Get additional command line arguments.
// dir: (optional) Specify the configuration file directory. If not appear, use "conf" directory by default.
// tag: (optional) Specify the profile of configuration file. will read "app-config-dev.yml" when set profile to "dev".
func loadCmdArguments() {
	ptr_ds := flag.String("d", "", "runtime root path")
	ptr_ps := flag.String("p", "", "config file profile name")
	ptr_rs := flag.String("r", "", "resources directory name")

	flag.Parse()

	dir := *ptr_ds
	profile := *ptr_ps
	resourcesDir := *ptr_rs

	if dir != "" {
		if !file.IsExist(dir) {
			panic(fmt.Sprint("given runtime root path is not exist:", dir))
		}
		app.Runtime.RootPath = dir
		log.Println("runtime root path specified to:", app.Runtime.RootPath)
	} else {
		_, f, _, _ := runtime.Caller(0)
		app.Runtime.RootPath = filepath.Dir(f)
		log.Println("no runtime root path specified, use current running path:", app.Runtime.RootPath)
	}

	if resourcesDir != "" {
		pathArray := []string{app.Runtime.RootPath, resourcesDir}
		resourcesPath := strings.Join(pathArray, string(os.PathSeparator))
		if !file.IsExist(resourcesPath) {
			panic(fmt.Sprint("given resources directory path is not exist:", resourcesPath))
		}
		app.Runtime.ResourcesDir = resourcesDir
		log.Println("resources directory specified to:", app.Runtime.ResourcesDir)
	} else {
		app.Runtime.ResourcesDir = app.DEFAULT_RESOURCES_DIR
		log.Println("no resources directory specified, use default directory:", app.DEFAULT_RESOURCES_DIR)
	}

	if profile != "" {
		app.Runtime.Profile = profile
		log.Println("config profile specified to:", app.Runtime.Profile)
	} else {
		app.Runtime.Profile = ""
		log.Println("no config profile specified")
	}
}

/* func StartWithNetHttp() {
	addr := fmt.Sprintf("%s:%d",
		config.APP_CONFIG.Server.HostName,
		config.APP_CONFIG.Server.Port)
	api.RegisterHandlers()
	log.Printf("start http service: %s\n", addr)
	e := http.ListenAndServe(addr, nil)
	if e != nil {
		log.Fatalln("meet error when start http service:", e)
	}
} */

func StartWithGin() {
	addr := fmt.Sprintf("%s:%d",
		config.APP_CONFIG.Server.HostName,
		config.APP_CONFIG.Server.Port)
	// Set gin runtime mode
	gin.SetMode(config.APP_CONFIG.Server.Env)

	r := gin.Default()
	api.RegisterHandlersGin(r)
	log.Printf("start http service: %s\n", addr)
	e := r.Run(addr)
	if e != nil {
		log.Fatalln("meet error when start http service:", e)
	}
}

func main() {
	loadCmdArguments()
	app.Init()
	log.Printf("%+v\n", config.APP_CONFIG)
	log.Printf("%+v\n", config.DB_CONFIG)

	defer db.CloseDBConnection()
	db.OpenDBConnection()

	StartWithGin()
}
