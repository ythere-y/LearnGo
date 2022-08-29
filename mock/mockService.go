package mock

type MockService struct {
}

func GetMockService() ServiceInter {
	return &MockService{}
}
func (c *MockService) Get(id int) (name string, err error) {
	return "mock service", nil
}
