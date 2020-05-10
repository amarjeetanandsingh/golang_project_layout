package trip

import (
	"fmt"
	"git.comp.com/foo/foo/internal/db/arango"
	"git.comp.com/foo/foo/internal/db/redis"
)

func NewTripRepo(arng *arango.Client, red *redis.Client) TripRepoI {
	return &TripRepo{
		Arango: arng,
		Redis:  red,
	}
}

type TripRepoI interface {
	Create(plan Trip) error
	FindById(id string) (Trip, error)
}

type TripRepo struct {
	Arango *arango.Client
	Redis  *redis.Client
}

func (repo *TripRepo) Create(plan Trip) error {
	if _, err := repo.Arango.CreateDoc("Trip", plan); err != nil {
		return err
	}
	return nil
}

func (repo *TripRepo) FindById(id string) (rp Trip, err error) {
	if err1 := repo.Redis.Get(id, &rp); err1 == nil && rp.ID != "" {
		return
	}
	//if err1 := repo.Arango.FindById("Trip", id, &rp); err1 != nil {
	//	err = err1
	//	return
	//}
	rp, ok := m[id]
	if !ok {
		return Trip{}, fmt.Errorf("not found")
	}
	return
}

// data
var m = map[string]Trip{
	"1": Trip{ID: "1"},
	"2": Trip{ID: "2"},
	"3": Trip{ID: "3"},
	"4": Trip{ID: "4"},
	"5": Trip{ID: "5"},
}
