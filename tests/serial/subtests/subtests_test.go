package tests

import (
	"fmt"
	"testing"
	"time"
)

func TestSerialization3(t *testing.T) {
	tests := []struct {
		name  string
		value int
	}{
		{name: "Serial SubTest #1", value: 5},
		{name: "Serial SubTest #2", value: 10},
	}
	fmt.Println(fmt.Sprintf("Serial Test #3 started at : %s", time.Now().String()))

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println(fmt.Sprintf("%s started at : %s", tc.name, time.Now().String()))
			time.Sleep(time.Duration(tc.value) * time.Second)
			fmt.Println(fmt.Sprintf("%s ended at : %s", tc.name, time.Now().String()))
		})
	}

	fmt.Println(fmt.Sprintf("Serial Test #3 ended at : %s", time.Now().String()))
}
