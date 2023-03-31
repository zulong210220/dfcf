package pkg

import (
	"fmt"
	"testing"

	. "github.com/zulong210220/dfcf/types"
)

func TestCli(t *testing.T) {
	cms := GetAllHkCodes()
	fmt.Println(len(cms))
}

func TestF10(t *testing.T) {
	cm := &CodeMarket{
		Market: 1,
		Code:   "603063",
	}
	cm = &CodeMarket{
		Market: 0,
		Code:   "300999",
	}
	cm = &CodeMarket{
		Market: 0,
		Code:   "002456",
	}
	GetF10(cm)
}
