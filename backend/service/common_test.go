package service

import (
	"fmt"
	"testing"
)

func TestGetModel(t *testing.T) {
	m, e := GetModel("fourquadrantlog", "log")
	fmt.Println(m, e)
}

func TestCheckStringLen(t *testing.T) {
	fmt.Println(CheckStringLen("abcdefg"))
}
