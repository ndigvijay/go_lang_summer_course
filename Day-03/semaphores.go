package main

import (
	"fmt"
	"go-lang.org/x/sync/semaphore"
)
	

func main(){
	fmt.Println("semaphores")
	// mutex and semaphores are used in synchronization 
	// primitives,used to control access to the shared
	// resource
	// mutex(mutual exclusion) 
	// mutex allows only one  thread at a time 
	// one thread only can mutex at any given time
	// ownership: A mutex is lock which is aquired by one thread and must be 
	// unlock by same thread

	// use case:typically used to protect citical section where a shared resource is accessed or modified
	// mutexes are binary: locked and unlocked
	// semaphores:
		// example:it maintains a counter which allows multiple thread to access a limited number 
	// of instances of the resource
	// counter of all resources
	// semaphores are also binary-unlock unlock but with certain conditions without ownership contraints
	// use case: specific number of 

	// summary:
	// mutex: one thread-one mutexconst

	// samphores:count of available resources
	// mulitple thread can access a single thread  a limited amount of resources simultanously
		//include a package called go-lang.org
		//importing it

}