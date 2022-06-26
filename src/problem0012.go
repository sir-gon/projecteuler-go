/**
 * Highly divisible triangular number
 *
 * https://projecteuler.net/problem=12
 *
 * The sequence of triangle numbers is generated by adding the natural numbers.
 * So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28.
 * The first ten terms would be:
 *
 * 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
 *
 * var us list the factors of the first seven triangle numbers:
 *
 *  1: 1
 *  3: 1,3
 *  6: 1,2,3,6
 * 10: 1,2,5,10
 * 15: 1,3,5,15
 * 21: 1,3,7,21
 * 28: 1,2,4,7,14,28
 * We can see that 28 is the first triangle number to have over five divisors.
 *
 * What is the value of the first triangle number to have over five hundred divisors?
 */

package projecteuler

import (
	"gon.cl/projecteuler.net/src/helpers"
	log "gon.cl/projecteuler.net/src/lib"
)

func Problem0012() int {
	var answer = 0

	const top = 500
	var triangular = 0
	var i = 1

	for answer < top {
		triangular += i
		var listOfDivisors = helpers.Divisors(triangular)
		var amountOfDivisors = len(listOfDivisors)

		log.Debug("Triangular number: %d has %d divisors", triangular, amountOfDivisors)

		answer = helpers.IntMax(answer, amountOfDivisors)

		i += 1
	}

	log.Info("Problem0012 answer => %d", answer)

	return answer
}
