package main

import (
	_ "github.com/joho/godotenv/autoload"
	common "github.com/vipos89/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"

	pb "github.com/vipos89/oms/common/api"
)

var (
	httpAddr      = common.EnvString("HTTP_ADDR", ":8080")
	ordersSvcAddr = common.EnvString("ORDER_SVC_ADDR", "127.0.0.1:2000")
)

func main() {
	conn, err := grpc.Dial(ordersSvcAddr, grpc.WithTransportCredentials(
		insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	log.Printf("connected to %s", ordersSvcAddr)
	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRotes(mux)
	log.Printf("Starting server on %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
