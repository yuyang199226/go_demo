package main

import (
	"fmt"
	"sync"
)

func main() {
	tpidTotrackids := map[string][]string {"t1": {"k1", "k2"}, "t2": {"k4","k8"}, "t3": {"k11","k3","k5"}}
	fmt.Println("xxx")
	wg := &sync.WaitGroup{}
	for tpid, trackids := range tpidTotrackids {
		wg.Add(1)
		go func(tpid string, trackids []string) {
			defer wg.Done()
			fmt.Println(tpid, trackids)
		}(tpid, trackids)


	}
	wg.Wait()
}
