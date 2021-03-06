/**
 * Smallest multiple
 *
 * https://projecteuler.net/problem=5
 *
 *
 *  2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
 *
 *  What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 */

package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0005(t *testing.T) {

	expectedSolution := 232792560
	bottom := 1
	top := 20
	startFrom := expectedSolution - 1000

	testname := fmt.Sprintf("Problem0005(%d, %d) => %v \n", bottom, top, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0005(bottom, top, startFrom)
		assert.Equal(t, expectedSolution, ans)
	})
}
