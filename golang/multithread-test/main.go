package main

import "example/workers"
import "fmt"
import "time"

func main() {

	scheduler := workers.NewScheduler()

	worker1 := scheduler.CreateWorker("generators/archlinux", 100)
	worker2 := scheduler.CreateWorker("generators/opensuse", 200)
	worker3 := scheduler.CreateWorker("generators/debian", 300)

	scheduler.AddWorker(worker1)
	scheduler.AddWorker(worker2)
	scheduler.AddWorker(worker3)

	go scheduler.Start()

	go func() {
		for {
			select {
			case <- scheduler.Context.Done():
				return
			default:
				fmt.Print("\u001b[2J\u001b[0f")
				fmt.Print("\u001b[3J")
				fmt.Println("Status Loop")
				scheduler.Print()
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	scheduler.Wait()

	fmt.Println("Program finished.")

}
