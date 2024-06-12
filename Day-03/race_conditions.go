//race condition

//in go a race condition arrises when multiple go routines(threads) concurrentky and the same time modify 
// the same piece of data, without any synchornization
// lack of co ordination-unpredictable outcome

// prevention , mutexes and locks
// ensures only one go -routine can acessa shared resource at a time, othe go routines which are attempting to
// access the resource will wait the mutex(lock) is unlock 


// channels, prevents race condition

// atomic operations

// out of syllabus- go race condition , built in tool , helps to indefntify race condition , durirng compilation


// functions used to prevent race conditions
// 1.sync-package - mutex.Lock() applies lock	
// 2.mutex.Unlock() released the lock

// sync.Mutex()- initialize a new sync mutex object
//               synchronize object to share resource

// sync.WaitGroup()


// wg.Add(n)- n= no of go routines that are added to the 
// call Add(n)
// s=w.Done() -signals other go routines that it has finished the task
