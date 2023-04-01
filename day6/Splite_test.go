package day6

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T)  {
	t.Log("test fun")
	if ok := reflect.DeepEqual("a","b");ok {
		fmt.Println("== == ==")
	}
}
func TestMain(m *testing.M)  {

}