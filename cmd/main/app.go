package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/shokhzodbek/internal/user"
	"github.com/shokhzodbek/pkg/logging"
)

func IndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name:=p.ByName("name")
	rw.Write([]byte(fmt.Sprintf("Hello %s",name)))

}
func main() {
	logger:=logging.GetLoger()
	logger.Trace()
	fmt.Println("Create router")
	// router created
	router := httprouter.New()
	// register user handler
	fmt.Println("User handler")
	handler:=user.New()
    handler.Register(router)
	// start application
	
	start(router)	
	
}

func start( router *httprouter.Router)  {
	logger:=logging.GetLoger()
	logger.Info()
	fmt.Println("Start application")
	listen,err :=net.Listen("tcp",":8000")
    if err!=nil {
		panic(err)
	}
	server:=&http.Server{
		Handler: router,
		WriteTimeout: 15*time.Second,
		ReadTimeout: 15*time.Second,
	}

	log.Fatalln(server.Serve(listen))
}