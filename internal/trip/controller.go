package trip

import (
	"git.comp.com/foo/foo/internal/db/arango"
	"git.comp.com/foo/foo/internal/db/redis"
	"git.comp.com/foo/foo/internal/routeplan"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func RegisterAPI(e *echo.Echo) {
	routePlanApiGroup := e.Group("/foo/trips")

	routePlanApiGroup.GET("/:id", findTripById)
	routePlanApiGroup.POST("/", createTrip)

	log.Println("Api configured: /foo/trips")
}

var arangoClient = arango.NewClient()
var redisClient = redis.NewClient()
var rpRepository = routeplan.NewRoutePlanRepo(arangoClient, redisClient)
var rpService = routeplan.NewRoutePlanService(rpRepository)
var tripService = NewTripService(arangoClient, redisClient, rpService)

func findTripById(c echo.Context) error {
	id := c.Param("id")
	resp, err := tripService.RoutePlanService.FindRoutePlanById(id)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, err)
	}
	if id != resp.ID {
		log.Println(id, resp.ID)
	}

	return c.JSON(http.StatusOK, resp)
}

func createTrip(c echo.Context) error {
	rpReq := Trip{}
	if err := c.Bind(&rpReq); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := tripService.CreateTrip(rpReq); err != nil {
		c.JSON(http.StatusExpectationFailed, err)
	}
	return c.JSON(http.StatusCreated, "")
}
