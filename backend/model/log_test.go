package model

import (
	"fmt"
	"testing"
)

func TestInt2Quadrant(t *testing.T) {
	fmt.Println(1, Int2Quadrant(1))
	fmt.Println(4, Int2Quadrant(4))
	fmt.Println(5, Int2Quadrant(5))
	fmt.Println(8, Int2Quadrant(8))
	fmt.Println(16, Int2Quadrant(16))
	fmt.Println(20, Int2Quadrant(20))

}
