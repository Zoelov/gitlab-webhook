package main

import (
	"net/http"
	"os"

	"github.com/zoelov/gitlab-webhook/utils/log"
	"github.com/zoelov/gitlab-webhook/utils/signal"
)

func main() {

	config := config.GetConfig()

	router := engine.NewRouter()

	routers.RegisterAPIRouter(router)

	addr := conf.GetStringOrDefault("http.addr", constants.HTTP_ADDR)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Errorf("Failed to ListenAndServe at %v, err = %v", addr, err)
			os.Exit(1)
		}
	}()

	log.Infoln("flora started!")

	sig := signal.WaitForExit() //signal exit: Ctrl+C or ...
	log.Infof("Got signal: %v, exit", sig)

}
