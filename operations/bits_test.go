package operations

import (
	//"fmt"
	//"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestRightMostOperations(t *testing.T) {
	for _, test := range []struct{
		x, r uint
	}{
		{0, 0},
		{0x58, 0x50},     // 01011000 , 01010000
	}{
		//fmt.Println(strconv.FormatUint(uint64(test.x), 2))
		//fmt.Println(strconv.FormatUint(uint64(test.r), 2))
		assert.Equal(t , TurnOffRightMostOne(test.x) , test.r )
	}



}
