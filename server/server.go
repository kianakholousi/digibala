package server

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func init() {
	e = echo.New()
}

func StartServer() {
	e.HideBanner = true
	bs, _ := os.ReadFile("server/banner.txt")
	fmt.Println(string(bs))

	supplierRoutes(e)
	addressRoutes(e)
	giftCardRoutes(e)
	productRoutes(e)
	promotionRoutes(e)
	categoryRoutes(e)
	adminRoutes(e)
	socialRoutes(e)
	brandRoutes(e)
	currencyRoutes(e)
	paymentRoutes(e)
	shippingRoutes(e)
	faqRoutes(e)

	log.Fatal(e.Start(":8080"))
}
