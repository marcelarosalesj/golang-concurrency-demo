package service

import (
	"fmt"
	"net/http"
)

type CarService interface {
	FetchData()
}

const (
	CarServiceUrl = "https://myfakeapi.com/api/cars/1"
)

type fetchCarDataService struct{}

func NewCarService() CarService {
	return &fetchCarDataService{}
}
func (*fetchCarDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the URL %s", CarServiceUrl)
	// call the external API

	resp, _ := client.Get(CarServiceUrl)

	//Write response to the channel
	carDataChannel <- resp

}
