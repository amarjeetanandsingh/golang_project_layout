package routeplan

import (
	"git.comp.com/foo/foo/internal/db/arango"
	"git.comp.com/foo/foo/internal/db/redis"
)

func NewRoutePlanRepo(arng *arango.Client, red *redis.Client) RoutePlanRepoI {
	return &RoutePlanRepo{
		Arango: arng,
		Redis:  red,
	}
}

type RoutePlanRepoI interface {
	Create(plan RoutePlan) error
	FindById(id string) (RoutePlan, error)
}

type RoutePlanRepo struct {
	Arango *arango.Client
	Redis  *redis.Client
}

func (repo *RoutePlanRepo) Create(plan RoutePlan) error {
	if _, err := repo.Arango.CreateDoc("RoutePlan", plan); err != nil {
		return err
	}
	return nil
}

func (repo *RoutePlanRepo) FindById(id string) (rp RoutePlan, err error) {
	if err1 := repo.Redis.Get(id, &rp); err1 == nil && rp.ID != "" {
		return
	}
	if err1 := repo.Arango.FindById("RoutePlan", id, &rp); err1 != nil {
		err = err1
		return
	}
	return
}

// data
var m = map[string]RoutePlan{
	"1": RoutePlan{ID: "1"},
	"2": RoutePlan{ID: "2"},
	"3": RoutePlan{ID: "3"},
	"4": RoutePlan{ID: "4"},
	"5": RoutePlan{ID: "5"},
}
