package routeplan

func NewRoutePlanService(repo RoutePlanRepoI) RoutePlanService {
	return RoutePlanService{Db: repo}
}

type RoutePlanService struct {
	Db RoutePlanRepoI
}

type RoutePlan struct {
	ID          string `json:"_key"`
	Name        string
	StartDate   string
	Source      string
	Destination string
}

func (rp *RoutePlanService) CreateRoutePlan(doc RoutePlan) error {
	if err := rp.Db.Create(doc); err != nil {
		return err
	}
	return nil
}

func (rp *RoutePlanService) FindRoutePlanById(id string) (RoutePlan, error) {
	routePlan, err := rp.Db.FindById(id)
	if err != nil {
		return RoutePlan{}, err
	}
	return routePlan, nil
}
