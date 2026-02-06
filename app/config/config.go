package config

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var JwtExpirationHours int
var CookieDomain string
var CookieSecure bool
var CookieSameSite http.SameSite

func mustGet(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("Variable requerida no definida: %s", key)
	}
	return v
}

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No se pudo cargar .env")
	}

	h, err := strconv.Atoi(mustGet("JWT_EXPIRATION_HOURS"))
	if err != nil {
		log.Fatal("JWT_EXPIRATION_HOURS inválido")
	}
	JwtExpirationHours = h

	CookieDomain = mustGet("COOKIE_DOMAIN")

	secureStr := mustGet("COOKIE_SECURE")
	CookieSecure = secureStr == "true"

	switch mustGet("COOKIE_SAME_SITE") {
	case "Strict":
		CookieSameSite = http.SameSiteStrictMode
	case "Lax":
		CookieSameSite = http.SameSiteLaxMode
	case "None":
		CookieSameSite = http.SameSiteNoneMode
	default:
		log.Fatal("COOKIE_SAME_SITE inválido")
	}
}