// https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75
package productofarrayexceptself

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type IntList []int

func IntListToString(list []int) string {
	strList := make([]string, len(list))
	for i, val := range list {
		strList[i] = fmt.Sprintf("%d", val)
	}
	return fmt.Sprintf("[%s]", strings.Join(strList, ","))
}

func Test(t *testing.T) {

	assert.Equal(t, IntListToString(productExceptSelf([]int{1, 2, 3, 4})), "[24,12,8,6]")
	assert.Equal(t, IntListToString(productExceptSelf([]int{-1, 1, 0, -3, 3})), "[0,0,9,0,0]")
	assert.Equal(t, IntListToString(productExceptSelf([]int{1, 2, 3, 4, 5})), "[120,60,40,30,24]")

}
