package conf

import (
	"fmt"
	"os"
	"strings"
)

var DefaultAdminUser string = ""
var DefaultAdminPassword string = ""

// InitFromEnv : get the config from environment variables
func InitFromEnv(isDevMode bool) bool {
	if isDevMode {
		fmt.Println("Dev mode is not authorized when initializing from env variables")
		os.Exit(1)
	}
	Port = os.Getenv("PORT")
	connStr := os.Getenv("DATABASE_URL")
	ConnStr = strings.Replace(connStr, "postgresql://", "", 1)
	EncodingKey = os.Getenv("KEY")
	DefaultAdminUser = os.Getenv("ADMIN_USER")
	DefaultAdminPassword = os.Getenv("ADMIN_PWD")
	mustRunAutoconf := false
	if DefaultAdminUser != "" && DefaultAdminPassword != "" {
		mustRunAutoconf = true
	}
	return mustRunAutoconf
}
