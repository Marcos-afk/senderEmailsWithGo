package internal

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

var (
	DATABASE_URL   string
	SERVER_PORT    string
	JWT_SECRET_KEY string
	WHITELIST      []*regexp.Regexp
)

func LoadEnvs() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}


	DATABASE_URL = getEnv("DATABASE_URL")
	SERVER_PORT = getEnv("SERVER_PORT")
	JWT_SECRET_KEY = getEnv("JWT_SECRET_KEY")
	WHITELIST = getCorsWhitelist()
}

func getCorsWhitelist() []*regexp.Regexp {
	whitelist := getEnv("WHITELIST_REQUESTS")
	domains := strings.Split(whitelist, ",")
	var regexList []*regexp.Regexp

	for _, domain := range domains {
		regexString := strings.ReplaceAll(domain, ".", "\\.")
		regexString = strings.ReplaceAll(regexString, "*", ".*")
		regexList = append(regexList, regexp.MustCompile("^" + regexString + "$"))
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
