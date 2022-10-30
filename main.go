package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/difaal21/go-template/config"
	"github.com/difaal21/go-template/responses"
	"github.com/difaal21/go-template/server"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload" //for development
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

var (
	cfg          *config.Config
	httpResponse = responses.HttpResponse{}
	messageIndex = "Application running properly"
)

func init() {
	cfg = config.Load()
}

func main() {
	logger := logrus.New()
	logger.SetFormatter(cfg.Logger.Formatter)
	logger.SetReportCaller(true)

	// basicAuth := middleware.NewBasicAuth(cfg.BasicAuth.Username, cfg.BasicAuth.Password)

	// validate := validator.New()

	router := mux.NewRouter()
	router.HandleFunc("/", index)
	http.Handle("/", router)

	handler := cors.New(cors.Options{
		AllowedOrigins:   cfg.Application.AllowedOrigins,
		AllowedMethods:   []string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "Authorization", "X-RECAPTCHA-TOKEN"},
		AllowCredentials: true,
	}).Handler(router)

	srv := server.NewServer(logger, handler, cfg.Application.Port)
	srv.Start()

	// When we run this program it will block waiting for a signal. By typing ctrl-C, we can send a SIGINT signal, causing the program to print interrupt and then exit.
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm

	// closing service for a gracefull shutdown.
	srv.Close()

}

// func index(w http.ResponseWriter, r *http.Request) {
// 	response := responses.HttpResponseStatusCodesImpl{}
// 	response.Ok("").NewResponses(nil, "Application running properly")
// 	// resp := response.NewSuccessResponse(nil, response.StatOK, indexMessage)
// 	// response.JSON(w, resp)
// }

func index(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, httpResponse.Ok("").NewResponses(nil, messageIndex))
}
