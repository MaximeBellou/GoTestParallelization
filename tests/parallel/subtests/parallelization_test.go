package tests

import (
	"fmt"
	"testing"
	"time"
)

func TestParallelization3(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		value int
	}{
		{name: "Parallel SubTest #1", value: 5},
		{name: "Parallel SubTest #2", value: 10},
	}
	fmt.Println(fmt.Sprintf("Parallel Test #3 started at : %s", time.Now().String()))

	for i := range tests {
		tc := tests[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fmt.Println(fmt.Sprintf("%s started at : %s", tc.name, time.Now().String()))
			time.Sleep(time.Duration(tc.value) * time.Second)
			fmt.Println(fmt.Sprintf("%s ended at : %s", tc.name, time.Now().String()))
		})
	}

	fmt.Println(fmt.Sprintf("Parallel Test #3 ended at : %s", time.Now().String()))
}
