package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/lemjoe/Grapho/internal/models"
)

func LookupAndParseEnvInt(envName string, defaultVal int) (int, bool) {
	env, exists := os.LookupEnv(envName)
	if !exists {
		return defaultVal, false
	}
	parsedInt, err := strconv.Atoi(env)
	if err != nil {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not valid: %w", envName, err))
		return defaultVal, false
	}
	return parsedInt, true
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateDefaultConfig() error {
	input, err := os.ReadFile("./.env.default")
	if err != nil {
		return err
	}

	err = os.WriteFile("./.env", input, 0644)
	if err != nil {
		return err
	}
	return nil
}

func InitConfig(confPath string) (models.ConfigDB, models.ConfigApp, error) {
	if confPath != "" {
		if fileExists(confPath) { //if from dot env
			if err := godotenv.Load(confPath); err != nil {
				return models.ConfigDB{}, models.ConfigApp{}, fmt.Errorf("InitConfig: unable to read file '%s'", confPath)
			}

		} else {
			if err := CreateDefaultConfig(); err != nil {
				return models.ConfigDB{}, models.ConfigApp{}, fmt.Errorf("unable to create default config. You should create it manually")
			}
			godotenv.Load(confPath)
		}

	}

	// Application config
	defaultConfApp := models.ConfigApp{
		Port:        "4007",
		JwtSecret:   "default_secret",
		AdminPasswd: "admin",
		MainLog:     "main.log",
	}
	APP_PORT, exist := os.LookupEnv("APP_PORT")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "APP_PORT"))
	} else {
		defaultConfApp.Port = APP_PORT
	}
	JWT_SECRET, exist := os.LookupEnv("JWT_SECRET")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "JWT_SECRET"))
	} else {
		defaultConfApp.JwtSecret = JWT_SECRET
	}
	ADMIN_PASSWD, exist := os.LookupEnv("ADMIN_PASSWD")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "ADMIN_PASSWD"))
	} else {
		defaultConfApp.AdminPasswd = ADMIN_PASSWD
	}
	MAIN_LOG, exist := os.LookupEnv("MAIN_LOG")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "MAIN_LOG"))
	} else {
		defaultConfApp.MainLog = MAIN_LOG
	}

	// Database config
	defaultConfDB := models.ConfigDB{
		DbType: "cloverdb",
		Path:   "./db",
	}
	DB_TYPE, exist := os.LookupEnv("DB_TYPE")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "DB_TYPE"))
	} else {
		defaultConfDB.DbType = DB_TYPE
	}
	DB_PATH, exist := os.LookupEnv("DB_PATH")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "DB_PATH"))
	} else {
		defaultConfDB.Path = DB_PATH
	}
	DB_PORT, exist := os.LookupEnv("DB_PORT")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "DB_PORT"))
	} else {
		defaultConfDB.Port = DB_PORT
	}
	DB_HOST, exist := os.LookupEnv("DB_HOST")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "DB_HOST"))
	} else {
		defaultConfDB.Host = DB_HOST
	}
	DB_NAME, exist := os.LookupEnv("DB_NAME")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "DB_NAME"))
	} else {
		defaultConfDB.DBName = DB_NAME
	}
	DB_USER, exist := os.LookupEnv("DB_USER")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "DB_USER"))
	} else {
		defaultConfDB.User = DB_USER
	}
	DB_PASSWD, exist := os.LookupEnv("DB_PASSWD")
	if !exist {
		fmt.Printf("warn: %s\n", fmt.Errorf("env '%s' not found", "DB_PASSWD"))
	} else {
		defaultConfDB.Password = DB_PASSWD
	}
	return defaultConfDB, defaultConfApp, nil
}
