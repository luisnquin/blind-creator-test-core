package environment

import (
	"log"
	"os"
	"strconv"
)

var (
	DbHost            = ""
	DbPort            = ""
	DbName            = ""
	DbUser            = ""
	DbPass            = ""
	DbEngine          = ""
	BasicAuthUsername = ""
	BasicAuthPassword = ""
	ServerPort        = 5000
	CorsWhitelist     = "*"
	ShowGormLogs      = false
	ApplyMigrations   = false
)

func InitializeEnv() {
	// godotenv.Load()
	DbHost = validateAndReturnVariable("DB_HOST")
	DbPort = validateAndReturnVariable("DB_PORT")
	DbName = validateAndReturnVariable("DB_NAME")
	DbUser = validateAndReturnVariable("DB_USER")
	DbPass = validateAndReturnVariable("DB_PASS")
	DbEngine = validateAndReturnVariable("DB_ENGINE")
	BasicAuthUsername = validateAndReturnVariable("BASIC_AUTH_USERNAME")
	BasicAuthPassword = validateAndReturnVariable("BASIC_AUTH_PASSWORD")
	ServerPort, _ = strconv.Atoi(validateAndReturnVariable("SERVER_PORT"))
	CorsWhitelist = validateAndReturnVariable("CORS_WHITELIST")
	ShowGormLogs, _ = strconv.ParseBool(validateAndReturnVariable("SHOW_GORM_LOGS"))
	ApplyMigrations, _ = strconv.ParseBool(validateAndReturnVariable("APPLY_MIGRATIONS"))
}

func validateAndReturnVariable(variableName string) string {
	variableValue := os.Getenv(variableName)
	if variableValue == "" {
		log.Fatal("an error happened getting the variable: '" + variableName + "' from the environment")
	} else {
		return variableValue
	}
	return ""
}
