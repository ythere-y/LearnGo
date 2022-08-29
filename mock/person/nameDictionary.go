package person

var names = [...]string{"Jack", "Queen", "King"}

func GetName(idx int) string {
	return names[idx]
}
