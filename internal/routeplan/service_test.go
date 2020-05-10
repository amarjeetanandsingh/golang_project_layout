package routeplan

import (
	"errors"
	"testing"
)

func TestRoutePlanService_FindRoutePlanById(t *testing.T) {
	rpRepo := newMockRepo()
	rpService := NewRoutePlanService(rpRepo)

	// put data in repo
	rpService.CreateRoutePlan(RoutePlan{ID: "1"})
	rp, err := rpService.FindRoutePlanById("1")
	if err != nil {
		t.Error("todo")
	}
	if rp.ID != "1" {
		t.Error("todo")
	}
}

type mockRepo struct {
	Db map[string]RoutePlan
}

func newMockRepo() *mockRepo {
	return &mockRepo{
		Db: map[string]RoutePlan{},
	}
}

func (r *mockRepo) Create(plan RoutePlan) error {
	if _, ok := r.Db[plan.ID]; ok {
		return errors.New("route plan already exists")
	}

	r.Db[plan.ID] = plan
	return nil
}

func (r *mockRepo) FindById(id string) (RoutePlan, error) {
	plan, ok := r.Db[id]
	if !ok {
		return RoutePlan{}, errors.New("...")
	}
	return plan, nil
}
