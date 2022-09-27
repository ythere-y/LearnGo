package concurrency

import (
	"fmt"
	"os"
	"time"
)

// 练习参考网站
// https://zengweigang.gitbooks.io/core-go/content/eBook/14.4.html

func Main() {
	//selectTimeOut()
	//panicWhileChannelEmpty()
	//closeChannelToAvoidAllSleep()
	//quitWithSelect()
	FetchData()
}

func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
}
func panicWhileChannelEmpty() {
	// fatal error: all goroutines are asleep - deadlock!
	ch := make(chan int)
	go tel(ch)
	for {
		i := <-ch
		fmt.Printf("ok is %t and the counter is at %d\n", i)
	}
}
func telWithClose(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	close(ch)
}
func closeChannelToAvoidAllSleep() {
	var ok = true
	var i int
	ch := make(chan int)
	go telWithClose(ch)
	for ok {
		if i, ok = <-ch; ok {
			fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
		}
	}
}
func telWithQuit(ch chan int, quit chan bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	quit <- true
}

func quitWithSelect() {
	ch := make(chan int)
	quit := make(chan bool)

	go telWithQuit(ch, quit)
	for {
		select {
		case i := <-ch:
			fmt.Printf("The counter is at %d\n", i)
		case <-quit:
			// break ; will fail, break will break from select
			os.Exit(0)
		}
	}

}

func selectTimeOut() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e3)
		timeout <- true
	}()
	ch := make(chan int, 1)
	go func() {
		time.Sleep(1e9)
		ch <- 101
	}()

	select {
	case u := <-ch:
		fmt.Println("get from ch = ", u)
	case <-timeout:
		fmt.Println("time out !")
	}
}
