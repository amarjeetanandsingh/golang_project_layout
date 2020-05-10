package routeplan

import (
	"git.comp.com/foo/foo/internal/db/arango"
	"git.comp.com/foo/foo/internal/db/redis"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func RegisterAPI(e *echo.Echo) {
	routePlanApiGroup := e.Group("/foo/routeplans")

	routePlanApiGroup.GET("/:id", findRoutePlanById)
	routePlanApiGroup.POST("/", createRoutePlan)

	log.Println("Api configured: /foo/routeplans")
}

var redisClient = redis.NewClient()
var arangoClient = arango.NewClient()
var rpRepository = NewRoutePlanRepo(arangoClient, redisClient)
var rpService = NewRoutePlanService(rpRepository)

func findRoutePlanById(c echo.Context) error {
	id := c.Param("id")
	resp, err := rpService.FindRoutePlanById(id)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func createRoutePlan(c echo.Context) error {
	rpReq := RoutePlan{}
	if err := c.Bind(&rpReq); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := rpService.CreateRoutePlan(rpReq); err != nil {
		c.JSON(http.StatusExpectationFailed, err)
	}
	return c.JSON(http.StatusCreated, "")
}
