package config

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Logger struct {
		Formatter logrus.Formatter
	}
	Application struct {
		Port           string
		Name           string
		AllowedOrigins []string
	}
}

func (cfg *Config) logFormatter() {
	formatter := &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			filename := fmt.Sprintf("%s:%d", f.File, f.Line)
			return funcname, filename
		},
	}

	cfg.Logger.Formatter = formatter
}

func (cfg *Config) app() {
	appName := os.Getenv("APP_NAME")
	port := os.Getenv("PORT")

	fmt.Println("appName"+appName, "port"+port)
	rawAllowedOrigins := strings.Trim(os.Getenv("ALLOWED_ORIGINS"), " ")

	allowedOrigins := make([]string, 0)
	if rawAllowedOrigins == "" {
		allowedOrigins = append(allowedOrigins, "*")
	} else {
		allowedOrigins = strings.Split(rawAllowedOrigins, ",")
	}

	cfg.Application.Port = port
	cfg.Application.Name = appName
	cfg.Application.AllowedOrigins = allowedOrigins
}

func Load() *Config {
	cfg := new(Config)
	cfg.logFormatter()
	cfg.app()
	return cfg
}
