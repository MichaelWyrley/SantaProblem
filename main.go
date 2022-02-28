package main

import (
	"fmt"
	"github.com/ChrisGora/semaphore"
	"math/rand"
	"sync"
	"time"
)

const NOELF = 10
const NORAINDEER = 9
const ELFPROBLEM = 0.1
const RAINDEERCOMEHOME = 0.0001
const NOOFPROBLEMS = 3
const NODAYS = 365
const (
	elfenum = iota
	raindeerenum = iota
)

var mutex = sync.Mutex{}

var elfProblem = semaphore.Init(NOOFPROBLEMS, NOOFPROBLEMS)
var elfcondition = sync.NewCond(&mutex)

var rainsemaphore = semaphore.Init(NORAINDEER, 0)
var wakeSanta = make(chan int)
var santaLock = sync.Mutex{}
var santawake = sync.NewCond(&santaLock)


func santa(stopWork []chan bool){
	for {
		santaLock.Lock()
		santawake.Wait()


		reason := <-wakeSanta
		if reason == elfenum {
			for elfProblem.GetValue() != NOOFPROBLEMS {
				elfProblem.Post()
			}
			elfcondition.Broadcast()
		} else if reason == raindeerenum {
			for _, i := range stopWork{
				i <- true
			}
			for elfProblem.GetValue() != NOOFPROBLEMS {
				elfProblem.Post()
			}
			elfcondition.Broadcast()

			fmt.Println("raindeer waking santa")
			break
		}
		santaLock.Unlock()
	}
}


func elf(stopWork chan bool){
	// randomly have a chance each day to have a problem
	stop := false
	for !stop{
		select {
		case <-stopWork:
			stop = true
		default:
			problem := rand.Float32()
			if problem <= ELFPROBLEM{
				// PROBLEM OCCURED
				fmt.Println("elf problem ")

				elfProblem.Wait()
				mutex.Lock()
				if elfProblem.GetValue() == 0{
					santawake.Broadcast()
					wakeSanta<- elfenum
				} else {
					elfcondition.Wait()
				}
				mutex.Unlock()
			}
			// work
			time.Sleep(100 * time.Millisecond)
		}

	}
}

func raindeer(isDecember chan bool){
	// randomly have a chance of coming back which increases every day up to chrispyman
	<-isDecember
	time.Sleep(time.Duration(rand.Intn(1000) )* time.Millisecond)
	fmt.Println("Raindeer woken up")
	rainsemaphore.Post()
	if rainsemaphore.GetValue() == NORAINDEER {
		wakeSanta <- raindeerenum
		santawake.Broadcast()
	}


}

func main(){

	raindeerDay := make([]chan bool, NORAINDEER)
	for i := range raindeerDay {
		raindeerDay[i] = make(chan bool, NODAYS)
		go raindeer(raindeerDay[i])
	}
	stopWorking := make([]chan bool, NOELF)
	for i := 0; i < NOELF; i++ {
		stopWorking[i] = make(chan bool, 1)
		go elf(stopWorking[i])
	}

	go santa(stopWorking)

	time.Sleep(1* time.Second)

	fmt.Println("Its december!")

	for _, e := range raindeerDay {
		e <- true
	}

	time.Sleep(10 * time.Second)

}