package sccommute

import (
	"fmt"
	"testing"
)

func TestInMem(t *testing.T) {
	postIM(LocPing{"simba", 12, -122.03})
	postIM(LocPing{"elaine", 14, -122.03})
	postIM(LocPing{"bugglez", 19, -122.03})
	postIM(LocPing{"elaine", 15, -122.03})
	printData()
	fmt.Println(getAllJSON())
}
