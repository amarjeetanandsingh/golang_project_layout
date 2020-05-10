package trip

import (
	"git.comp.com/foo/foo/internal/db/arango"
	"git.comp.com/foo/foo/internal/db/redis"
	"git.comp.com/foo/foo/internal/routeplan"
)

func NewTripService(arangoClient *arango.Client, redisClient *redis.Client, rpService routeplan.RoutePlanService) TripService {
	repo := NewTripRepo(arangoClient, redisClient)
	return TripService{Db: repo, RoutePlanService: rpService}
}

type TripService struct {
	Db               TripRepoI
	RoutePlanService routeplan.RoutePlanService
}

type Trip struct {
	ID          string
	Name        string
	Source      string
	Destination string
}

func (rp *TripService) CreateTrip(doc Trip) error {
	if err := rp.Db.Create(doc); err != nil {
		return err
	}
	return nil
}

func (rp *TripService) FindTripById(id string) (Trip, error) {
	routePlan, err := rp.Db.FindById(id)
	if err != nil {
		return Trip{}, err
	}
	return routePlan, nil
}
