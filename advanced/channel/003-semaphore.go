package main

type Empty interface{}

type Semaphore chan Empty

// P acquire n resources
func (sem Semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		sem <- e
	}
}

// V release n resources
func (sem Semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-sem
	}
}

// Lock mutex
func (sem Semaphore) Lock() {
	sem.P(1)
}

// Unlock mutex
func (sem Semaphore) Unlock() {
	sem.V(1)
}

/*
signal-wait model
*/
func (sem Semaphore) Wait(n int) {
	sem.P(n)
}

func (sem Semaphore) Signal() {
	sem.V(1)
}

func main() {

}
