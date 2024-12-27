package internal

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	DATABASE_URL   	string
	SERVER_PORT    	string
	JWT_SECRET_KEY 	string
	WHITELIST      	[]*regexp.Regexp
	MAIL_FROM				string
	MAIL_PASSWORD		string
	MAIL_SMTP 			string
	MAIL_PORT 			int
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	DATABASE_URL = getEnv("DATABASE_URL")
	SERVER_PORT = getEnv("SERVER_PORT")
	JWT_SECRET_KEY = getEnv("JWT_SECRET_KEY")
	WHITELIST = getCorsWhitelist()
	MAIL_FROM = getEnv("MAIL_FROM")
	MAIL_PASSWORD = getEnv("MAIL_PASSWORD")
	MAIL_SMTP = getEnv("MAIL_SMTP")
	MAIL_PORT = getEnvAsInt("MAIL_PORT")
}

func getCorsWhitelist() []*regexp.Regexp {
	whitelist := getEnv("WHITELIST_REQUESTS")
	domains := strings.Split(whitelist, ",")
	var regexList []*regexp.Regexp

	for _, domain := range domains {
		domain = strings.TrimPrefix(domain, "http://")
		domain = strings.TrimPrefix(domain, "https://")

		regexString := strings.ReplaceAll(domain, ".", "\\.")
		regexString = strings.ReplaceAll(regexString, "*", ".*")

		regexString = "^(http|https)://" + regexString + "$"
		regexList = append(regexList, regexp.MustCompile(regexString))
	}

	return regexList
}


func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("A variável de ambiente %s não foi encontrada", key)
	}
	return value
}

func getEnvAsInt(key string) int {
	value, err := strconv.Atoi(getEnv(key))
	if err != nil {
		log.Fatalf("Erro ao converter %s para int: %v", key, err)
	}
	return value
}
