package sync

import "fmt"
import "sync"
import "time"

var SharedResource = make(map[string]interface{})

func WaitGroup() {

	var waitgroup sync.WaitGroup

	waitgroup.Add(2)

	mutex := sync.Mutex{}
	condition := sync.NewCond(&mutex)

	go func() {

		condition.L.Lock()

		for _, ok := SharedResource["resource1"]; ok == false; {
			condition.Wait()
		}

		fmt.Println(SharedResource["resource1"])

		condition.L.Unlock()
		waitgroup.Done()
		fmt.Println("Done with Task 1")

	}()

	go func() {

		condition.L.Lock()

		for _, ok := SharedResource["resource2"]; ok == false; {
			condition.Wait()
		}

		fmt.Println(SharedResource["resource2"])

		condition.L.Unlock()
		waitgroup.Done()
		fmt.Println("Done with Task 2")

	}()


	condition.L.Lock()

	SharedResource["resource1"] = "foo"
	time.Sleep(3 * time.Second)
	SharedResource["resource2"] = "bar"

	condition.Broadcast()
	condition.L.Unlock()

	waitgroup.Wait()
	fmt.Println("Done with all tasks")

}
