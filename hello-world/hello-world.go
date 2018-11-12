package main
import "fmt"
import "time"

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")
    done <- true
}


func main() {
	var array_temp [2]int
	done := make(chan bool, 1)
	go worker(done)
	<- done
	fmt.Println("emp:", array_temp)
	fmt.Println("hello world")
}