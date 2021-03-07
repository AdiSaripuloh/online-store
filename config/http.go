package config

import (
	"github.com/gin-gonic/gin"
	"os"
)

type HttpConf struct {
	HttpPort      string
	HttpDebugMode string
	HttpDebug     bool
	DBDriver      string
	DBMigration   *bool
	DBSeed        *bool
}

const (
	MYSQL           = "mysql"
	DefaultDbDriver = MYSQL
)

var HttpConfig *HttpConf
var supportedDBDriver = []string{MYSQL}

func NewHttpConfig(httpPort string, dbMigration *bool, dbSeed *bool) {
	httpDebug := getHttpDebug()
	httpDebugMode := getHttpDebugMode()
	dbDriver := getDBDriver()
	HttpConfig = &HttpConf{
		HttpPort:      httpPort,
		HttpDebugMode: httpDebugMode,
		HttpDebug:     httpDebug,
		DBDriver:      dbDriver,
		DBMigration:   dbMigration,
		DBSeed:        dbSeed,
	}
}

func getHttpDebugMode() string {
	debugMode := os.Getenv("APP_MODE")
	if debugMode == "" {
		return gin.ReleaseMode
	}
	return debugMode
}

func getHttpDebug() bool {
	httpDebug := os.Getenv("APP_MODE")
	if httpDebug == "" {
		return false
	}
	return httpDebug == gin.DebugMode
}

func getDBDriver() string {
	dbDriver := os.Getenv("APP_MODE")
	if dbDriver == "" {
		return DefaultDbDriver
	}
	if !isSupportedDBDriver(dbDriver) {
		return DefaultDbDriver
	}
	return DefaultDbDriver
}

func isSupportedDBDriver(dbDriver string) bool {
	for _, driver := range supportedDBDriver {
		if driver == dbDriver {
			return true
		}
	}
	return false
}
