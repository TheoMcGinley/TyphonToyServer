package main

/*
 * the following is a toy server, and is not intended as a
 * real-world example, as a bulletproof server, or even as
 * a perfect example of REST server design - it was built for fun
 * in an evening to better understand the internals of Monzo's typhon library
 */

import (
	"context"
	"errors"
	"fmt"
	"github.com/monzo/typhon"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
	"typhontoyserver/db"
)

func getID(URI string) (int, error) {
	uriParts := strings.Split(URI, "/")
	if len(uriParts) < 3 {
		return -1, errors.New("no ID provided")
	}

	id, err := strconv.Atoi(uriParts[2])
	if err != nil {
		return -1, err
	}

	return id, nil
}

func getComment(req typhon.Request) typhon.Response {
	commentID, err := getID(req.RequestURI)
	if err != nil {
		return typhon.Response{
			Error: err,
			Request: &req,
		}
	}

	return req.Response(db.Get(commentID))
}

func postComment(req typhon.Request) typhon.Response {
	commentID, err := getID(req.RequestURI)
	if err != nil {
		return typhon.Response{
			Error: err,
			Request: &req,
		}
	}

	// no body.Close() required!
	body, err := req.BodyBytes(true)
	if err != nil {
		return typhon.Response{
			Error: err,
			Request: &req,
		}
	}

	db.Post(commentID, string(body))
	return req.Response("")
}

func main() {
	router := typhon.Router{}

	// won't accept "/comments" without an ID
	router.GET("/comments/*", getComment)
	router.POST("/comments/*", postComment)

	service := router.Serve().
		Filter(typhon.ErrorFilter).
		Filter(typhon.H2cFilter)
	server, err := typhon.Listen(service, ":8888")
	if err != nil {
		panic(err)
	}
	log.Println("listening on " + server.Listener().Addr().String())
	fmt.Println(router)

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Printf("shutting down...")
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.Stop(c)
}
