package epan_test

import (
	"fmt"
	"testing"

	"github.com/tmc/epan"
)

func TestBasics(t *testing.T) {
	fmt.Println(epan.Version())
}
