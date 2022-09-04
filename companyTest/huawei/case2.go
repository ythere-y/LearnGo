package huawei

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func move(mp [][]int, x int, y int, row int, col int, step int) int {
	//fmt.Printf("arrive [%d,%d],map = %v, step = %d\n", x, y, mp, step)
	dst := make([][]int, len(mp))
	for i, _ := range dst {
		dst[i] = make([]int, len(mp[0]))
		copy(dst[i], mp[i])
	}
	switch mp[x][y] {
	case -1:
		return math.MaxInt / 2
	case 0:
		dst[x][y] = -1
		return dfs(dst, x, y, row, col, step+1)
	case 1:
		return math.MaxInt / 2
	case 6:
		dst[x][y] = -1
		if x-1 >= 0 && dst[x-1][y] == 1 {
			dst[x-1][y] = 0
		}
		if x+1 < row && dst[x+1][y] == 1 {
			dst[x+1][y] = 0
		}
		if y-1 >= 0 && dst[x][y-1] == 1 {
			dst[x][y-1] = 0
		}
		if y+1 < col && dst[x][y+1] == 1 {
			dst[x][y+1] = 0
		}
		return dfs(dst, x, y, row, col, step+1)
	case 4:
		dst[x][y] = -1
		return dfs(dst, x, y, row, col, step+3)
	default:
		return step
	}
}
func dfs(mp [][]int, x int, y int, row int, col int, step int) int {
	fmt.Printf("arrive [%d,%d], step = %d\n", x, y, step)
	var res = step
	if x-1 >= 0 {
		res = move(mp, x-1, y, row, col, step)
	}
	if x+1 < row {
		res = int(math.Min(float64(res), float64(move(mp, x+1, y, row, col, step))))
	}
	if y-1 >= 0 {
		res = int(math.Min(float64(res), float64(move(mp, x, y-1, row, col, step))))
	}
	if y+1 < col {
		res = int(math.Min(float64(res), float64(move(mp, x, y+1, row, col, step))))
	}

	return res

}

func main2() {

	//r := bufio.NewReader(os.Stdin)
	//strrow, _ := r.ReadString(' ')
	//strcol, _ := r.ReadString('\n')
	strrow := " 4"
	strcol := " 4"

	strrow = strings.Trim(strrow, " ")
	strcol = strings.TrimSpace(strcol)
	fmt.Printf("cow = %s, col = %s\n", strrow, strcol)
	Norow, _ := strconv.ParseInt(strrow, 10, 64)
	Nocol, _ := strconv.ParseInt(strcol, 10, 64)
	row := int(Norow)
	col := int(Nocol)
	fmt.Printf("cow = %v,strcol = %v\n", row, col)

	var ar [][]int

	ar = append(ar, []int{1, 1, 1, 1})
	ar = append(ar, []int{1, 6, 2, 1})
	ar = append(ar, []int{1, 1, 0, 1})
	ar = append(ar, []int{1, 3, 1, 1})
	//for i := 0; i < row; i++ {
	//	var line []int
	//	for j := 0; j < col; j++ {
	//		var strCur string
	//		if j == col-1 {
	//			strCur, _ = r.ReadString('\n')
	//		} else {
	//			strCur, _ = r.ReadString(' ')
	//		}
	//		NoCur, _ := strconv.ParseInt(strings.TrimSpace(strCur), 10, 64)
	//		cur := int(NoCur)
	//		line = append(line, cur)
	//	}
	//	ar = append(ar, line)
	//}
	//fmt.Printf("ar = \n%v\n", ar)

	startx := 0
	starty := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if ar[i][j] == 2 {
				startx = i
				starty = j
				ar[i][j] = -1
			}
		}
	}

	step := dfs(ar, startx, starty, row, col, 0)
	fmt.Printf("\n step = %d\n", step)
}
