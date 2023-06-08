// https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75
package kidswiththegreatestnumberofcandies

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Join(elems []bool, sep string) string {
	var output []string
	for i := 0; i < len(elems); i++ {
		output = append(output, fmt.Sprintf("%t", elems[i]))
	}

	return fmt.Sprintf("[%s]", strings.Join(output, sep))
}

func Test(t *testing.T) {

	assert.Equal(t, Join(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3), ","), "[true,true,true,false,true]")
	assert.Equal(t, Join(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1), ","), "[true,false,false,false,false]")
	assert.Equal(t, Join(kidsWithCandies([]int{12, 1, 12}, 10), ","), "[true,false,true]")

}
