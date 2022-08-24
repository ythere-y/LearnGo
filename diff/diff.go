package diff

import "fmt"
import "github.com/sergi/go-diff/diffmatchpatch"

func Main() {
	fmt.Printf("diff main started\n")
	fmt.Printf("text1 -> \n%+v\n", text1)
	fmt.Printf("text2 -> \n%+v\n", text2)
	fmt.Printf("****** show ******\n")
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(text1, text2, false)
	fmt.Println(diffs)
	fmt.Printf("[pretty girle]\n")
	fmt.Printf("%+v", dmp.DiffPrettyText(diffs))
	fmt.Printf("[pretty html]\n")
	fmt.Printf("%+v", dmp.DiffPrettyHtml(diffs))

}
