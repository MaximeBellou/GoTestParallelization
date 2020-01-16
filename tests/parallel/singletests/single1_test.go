package tests

import (
	"fmt"
	"testing"
	"time"
)

func TestParallelization1(t *testing.T) {
	t.Parallel()
	fmt.Println(fmt.Sprintf("Test in Parallel #1 started at : %s", time.Now().String()))
	time.Sleep(5 * time.Second)
	fmt.Println(fmt.Sprintf("Test in Parallel #1 ended at : %s", time.Now().String()))
}
