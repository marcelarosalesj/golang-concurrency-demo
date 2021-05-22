package service

import (
	"fmt"
	"net/http"
)

type OwnerService interface {
	FetchData()
}

const (
	OwnerServiceUrl = "https://myfakeapi.com/api/users/1"
)

type fetchOwnerDataService struct{}

func NewOwnerService() CarService {
	return &fetchOwnerDataService{}
}
func (*fetchOwnerDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the URL %s", OwnerServiceUrl)
	// call the external API

	resp, _ := client.Get(OwnerServiceUrl)

	//Write response to the channel
	ownerDataChannel <- resp

}
