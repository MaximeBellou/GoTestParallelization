package tests

import (
	"fmt"
	"testing"
	"time"
)

func TestSerialization2(t *testing.T) {
	fmt.Println(fmt.Sprintf("Serial Test #2 started at : %s", time.Now().String()))
	time.Sleep(5 * time.Second)
	fmt.Println(fmt.Sprintf("Serial Test #2 ended at : %s", time.Now().String()))
}
