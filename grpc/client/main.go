package main

import (
	"context"
	user "grpc/proto/gen"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial(":8200", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Não conectou: %s", err)
	}

	client := user.NewCnpjServiceClient(conn)

	req := user.CnpjRequest{
		Cnpj: "60701190000104",
	}

	response, err := client.GetCnpjInfo(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(response)
}