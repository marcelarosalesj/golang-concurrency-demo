package main

import (
	"github.com/marcelarosalesj/golang-concurrency-demo/controller"
	router "github.com/marcelarosalesj/golang-concurrency-demo/http"
	"github.com/marcelarosalesj/golang-concurrency-demo/service"
)

var (
	carDetailsService    service.CarDetailsService       = service.NewCarDetailsService()
	carDetailsController controller.CarDetailsController = controller.NewCarDetailsController(carDetailsService)
	httpRouter           router.Router                   = router.NewChiRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/carDetails", carDetailsController.GetCarDetails)
	httpRouter.SERVE(port)
}
