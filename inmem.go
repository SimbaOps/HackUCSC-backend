package sccommute

import (
	"encoding/json"
	"fmt"
	"sync"
)

func init() {
	//I want the indicies to start at 1 so I fill the 0th spot
	//with a zeroed id
	data = append(data, LocPing{})
}

var idx map[string]int = make(map[string]int)
var data []LocPing = make([]LocPing, 0)
var lock *sync.Mutex = &sync.Mutex{}

func postIM(loc LocPing) {
	lock.Lock()
	defer lock.Unlock()
	if i := idx[loc.BusId]; i == 0 {
		idx[loc.BusId] = len(data)
		data = append(data, loc)
	} else {
		data[i] = loc
	}
}

func getIM(id string) LocPing {
	lock.Lock()
	defer lock.Unlock()
	return data[idx[id]]
}

func getAllJSON() string {
	bytes, _ := json.Marshal(data[1:])
	return string(bytes)
}

func printData() {
	fmt.Println(data)
	fmt.Println(idx)
}
