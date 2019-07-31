package main

import (
	"go-music/repositories"
	"go-music/router"
	"log"
	"net/http"
)

func main()  {
	ctx, err := repositories.NewMusicContext("mysql", "root:12345678x@X@/shop?parseTime=true")
	if err != nil {
		log.Fatal("Error init db context: ", err)
	}
	routes := router.Routes{}
	handler := routes.Init(ctx)

	server := &http.Server{
		Addr: ":9005",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Panicf("listen: %s\n", err)
	}
}
