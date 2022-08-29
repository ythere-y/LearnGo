package person

type Male interface {
	Get(id int) int
}

type RealMale struct {
}

func (m *RealMale) Get(id int) int {
	return -id
}
