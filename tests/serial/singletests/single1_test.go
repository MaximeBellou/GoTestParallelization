package tests

import (
	"fmt"
	"testing"
	"time"
)

func TestSerialization1(t *testing.T) {
	fmt.Println(fmt.Sprintf("Serial Test #1 started at : %s", time.Now().String()))
	time.Sleep(5 * time.Second)
	fmt.Println(fmt.Sprintf("Serial Test #1 ended at : %s", time.Now().String()))
}
