package main

import (
	"log"
	"net"
	user "grpc/proto/gen"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"google.golang.org/grpc"
)

type Server struct {
	user.UnimplementedCnpjServiceServer
}

type Cnpj struct {
	UF                           string `json:"uf"`
	CEP                          string `json:"cep"`
	Natureza_Juridica            string `json:"natureza_juridica"`
	Descricao_Situacao_Cadastral string `json:"descricao_situacao_cadastral"`
	CNPJ                         string `json:"cnpj"`
	CNAE                         string `json:"cnae_fiscal_descricao"`
}


func (s *Server) GetCnpjInfo(ctx context.Context, in *user.CnpjRequest) (*user.CnpjResponse, error) {
	log.Printf("Message Received", in.Cnpj)

	response, err := http.Get(fmt.Sprintf("https://brasilapi.com.br/api/cnpj/v1/%v", in.Cnpj))
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	cnpj := Cnpj{}

	jsonErr := json.Unmarshal(body, &cnpj)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return &user.CnpjResponse{
		Uf: cnpj.UF,
		Cep: cnpj.CEP,
		NaturezaJuridica: cnpj.Natureza_Juridica,
		DescricaoSituacaoCadastral: cnpj.Descricao_Situacao_Cadastral,
		Cnpj: cnpj.CNPJ,
		CnaeFiscalDescricao: cnpj.CNAE,
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", ":8200")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	
	grcpServer := grpc.NewServer()

	user.RegisterCnpjServiceServer(grcpServer, &Server{})

	log.Println("Listening on Port: 8200")

	if err := grcpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}