package huawei

import (
	"fmt"
	"strings"
	"unicode"
)

func Main() {
	main2()
}

func main1() {

	//r := bufio.NewReader(os.Stdin)
	line := "Hello World."
	words := "Good Hello LOOP"
	//line, _ := r.ReadString('\n')
	//words, _ := r.ReadString('\n')
	//fmt.Printf("line = %v, words = %v\n", line, words)
	strarray := strings.Fields(strings.TrimSpace(strings.ToLower(words)))

	//fmt.Printf("get %v\n", strarray)
	mp := make(map[string]int)
	for i := 0; i < len(strarray); i++ {
		mp[strarray[i]] = i
	}
	res := strings.Builder{}
	var l = len(line)
	var yinFlag = false
	for i := 0; i < l; i++ {
		//fmt.Printf("idx = %d\n", i)
		if unicode.IsLetter(rune(line[i])) == false {
			res.WriteString(string(line[i]))
			if line[i] == '"' {
				yinFlag = !yinFlag
			}
			continue
		}
		if yinFlag == true {
			res.WriteString(string(line[i]))
		} else {

			wd := strings.Builder{}
			back := i
			for back = i; back < l; back++ {
				if unicode.IsLetter(rune(line[back])) == true {
					wd.WriteString(string(line[back]))
					continue
				} else {
					break
				}
			}
			//fmt.Printf("wd = %v\n", wd.String())
			copWd := strings.ToLower(wd.String())
			if _, ok := mp[copWd]; ok {
				//fmt.Printf("get replace idx = %d\n", mp[copWd])
				rpIdx := mp[copWd]
				res.WriteString(fmt.Sprintf("%d", rpIdx))
			} else {
				res.WriteString(wd.String())
			}
			i = back - 1
		}
	}
	fmt.Printf("%s", res.String())
}
