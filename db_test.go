package sccommute

import (
	"testing"
)

func TestPost(t *testing.T) {
	Post(LocPing{"Simba", 37, -122.0539})
	//	defer close(postChannel)
}
