package main

import (
	"fmt"
	"sync"
	"time"
)

// Suppose we are you know an image processing processing an image
// it takes the image or the url of the image then it process it then we re
// we return it

func worker(url string) string {
	time.Sleep(time.Millisecond * 60)
	fmt.Println("image processed ", url)
	return  url
}
func main() {
	// a variable name waitgroup and type is sync
	var wg sync.WaitGroup
	// fmt.Println("Hello concurency")
	// how much time the worker is taking or my above function is taking
	startime := time.Now()
	wg.Add(2)
	go worker("image_12.png")
	go worker("image_13.png")
	// fmt.Println("result", result)
	// fmt.Println("result1", result1)
	// wg maints a counter 
	wg.Wait()
	fmt.Printf("it took %s ms", time.Since(startime))

}

// my exp with c and cpp tells me first the main functions runs
// then it goes to the result the control is passed then
// it goes to result 2 then my you know result 2 is you
// in nutshel li cnat do anyting until image 12 stops after it stops it goes to you know
// here i  basiclly my main function fires off and executes basicly my worker doesnt get a chance to execute 
// we need to syncronize or need to wait we need to use waitgroups 
