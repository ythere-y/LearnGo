package mock

type ServiceInter interface {
	Get(id int) (name string, err error)
}

func GetService() ServiceInter {
	return &RealService{}
}

type RealService struct {
}

func (c *RealService) Get(id int) (name string, err error) {
	return "real service", nil
}
