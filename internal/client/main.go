package client

import (
	"log"
	"os"

	"github.com/bm-krishna/tenant-service/pkg/api/tenant"
	"google.golang.org/grpc"
)

func ConfigClient() (tenant.TenantClient, error) {
	// client will send request to server by using grpc.Dail wich takes server address
	address := os.Getenv("ADDRESS")
	if address == "" {
		address = "7171"
	}
	clientConnection, err := grpc.Dial("172.17.0.1:"+address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Server Didn't connect", err)
		return nil, err
	}
	log.Println(clientConnection, "clientConnection")
	// register client with connect
	stub := tenant.NewTenantClient(clientConnection)

	// create request
	// payload := make(map[string]interface{})
	// payload["name"] = "mohan
	return stub, nil
}
