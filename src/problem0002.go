/**
 * Even Fibonacci numbers
 *
 * https://projecteuler.net/problem=2
 *
 * Each new term in the Fibonacci sequence is generated by adding the previous two terms.
 * By starting with 1 and 2, the first 10 terms will be:
 *
 * 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
 *
 * By considering the terms in the Fibonacci sequence whose values do not exceed four million,
 * find the sum of the even-valued terms.
 */

package projecteuler

import utils "gon.cl/projecteuler.net/src/utils"

func Problem0002(_top int) int {

	last1 := 1
	last2 := 0
	evenSum := 0

	i := 0
	fibo := 0

	for ok := true; ok; ok = fibo < _top {
		fibo = last2 + last1

		utils.Debug("Fibonacci (%d) = %d", i, fibo)

		// acumulate sum of event terms
		if fibo%2 == 0 {
			evenSum += fibo
		}

		// next keys in loop:
		last2 = last1
		last1 = fibo
		i += 1
	}

	utils.Info("Problem0002 result: %d", evenSum)

	return evenSum
}
