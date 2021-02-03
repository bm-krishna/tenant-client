package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bm-krishna/tenant-client/internal/client"
	"github.com/bm-krishna/tenant-service/pkg/api/tenant"
	// "github.com/bm-krishna/tenant-service/pkg/api/tenant"
)

type Service struct {
}

func (service *Service) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	urlPath := request.URL.Path
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stub, err := client.ConfigClient()
	if err != nil {
		log.Fatal("Failed to Config client")
		os.Exit(1)
	}
	log.Println(urlPath, "urlPath")
	if strings.Contains(urlPath, "/api") {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(rw, "Failed to Read body from Request", http.StatusBadRequest)
		}
		if err != nil {
			log.Fatal("Failed to Marshal payload to protoc Any")
		}
		tenantRequest := &tenant.TenantRequest{
			Request: body,
		}
		tenantRsp, err := stub.API(ctx, tenantRequest)
		if err != nil {
			log.Fatal("Failed to get Response", err)
		}
		data := tenantRsp.Response
		tenantResp := make(map[string]interface{})
		err = json.Unmarshal(data, &tenantResp)
		if err != nil {
			log.Fatal("Failed to unmarshal response")
		}
		log.Println("***************************")
		log.Println(tenantResp, "tenantResp")
		fmt.Fprintf(rw, "%s", tenantRsp.Response)
		return
	}
	fmt.Fprintf(rw, "%s", "Path is not found")
}
