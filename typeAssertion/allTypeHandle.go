package typeAssertion

import "fmt"

func addNumber(arg []int) int {
	sum := 0
	for _, n := range arg {
		sum += n
	}
	return sum
}

func addFloat64(arg []float64) float64 {
	sum := 0.0
	for _, f := range arg {
		sum += f
	}
	return sum
}
func addVar(list interface{}) interface{} {
	var varSum interface{}
	switch list.(type) {
	case []int:
		varSum = addNumber(list.([]int))
	case []float64:
		varSum = addFloat64(list.([]float64))
	default:
		fmt.Printf("switch failed\n")
	}
	return varSum

}
func allTypeSum() {
	//varSum := []int{1, 2, 3}
	varSum := []float64{1.1, 2.2, 3.3}
	res := addVar(varSum)
	fmt.Printf("result = %v\n", res)
}

type Father interface {
	Get()
}

type BoBoReq struct {
	Name string
}

func (c BoBoReq) Get() {

}

type BoBoResp struct {
	Name string
}
type CoCoReq struct {
	Score int
}

func (c CoCoReq) Get() {

}

type CoCoResp struct {
	Score int
}

func checkReqNil(req interface{}) interface{} {
	//var resp interface{}
	switch (req).(type) {
	case BoBoReq:
		return BoBoReq{Name: "bobo"}
	case CoCoReq:
		return CoCoResp{Score: -1}
	case Father:
		return nil
	default:
		return nil
	}
}
func funcTypeToType() {
	var Breq = BoBoReq{Name: "TOM"}
	var Creq = CoCoReq{Score: 123}

	Bresp := checkReqNil(Breq)
	fmt.Printf("B resp = %#v\n", Bresp)
	Cresp := checkReqNil(Creq)
	fmt.Printf("C resp = %#v\n", Cresp)

}
