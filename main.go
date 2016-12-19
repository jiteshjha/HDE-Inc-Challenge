/*
 * Author: Jitesh Kumar Jha
 * Email : jiteshjha96@gmail.com
 * Github: jiteshjha
 *
 * **Description**
 *
 * We want to calculate a sum of squares of some integers, excepting negatives
 * The first line of the input will be an integer N (1 <= N <= 100)
 * Each of the following N test cases consists of one line containing an integer X (0 < X <= 100),
 * followed by X integers (Yn, -100 <= Yn <= 100) space-separated on the next line
 * For each test case, calculate the sum of squares of the integers excepting negatives,
 * and print the calculated sum to the output. No blank line between test cases
 * (Take input from standard input, and output to standard output)
 *
 */

package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
)

/*
 * Input: Integer (num)
 * Output: Square of the integer
 *
 * get_square(num) computes the square of integer num
 * if the integer is positive, else returns a 0
 */
func get_square(num int) int {
  if num < 0 {
    return 0
  }
    return num * num
}

/*
 * Input: 1) position in the string array (pos)
 *        2) String array of integers (arr)
 * Output: Resultant sum of squares
 *
 * compute(pos, arr) recursively computes
 * the sum of squares of positive integers
 */

func compute(pos int, arr[] string) int {

  // Convert string to integer value
  i, err := strconv.Atoi(arr[pos])

  // If error occurs, raise the error
  if(err != nil) {
    panic(err)
  }

  if pos == 0 {
    return get_square(i)
  }
    return get_square(i) + compute(pos - 1, arr)
}

/*
 * Input: Number of sum of squares computations (iterations)
 * Gets 1) number of integers (num_integers)
 *         of range (0 < num_integers <= 100)
 *      2) String of integers (input)
 */

func sum_square(iterations int) {
  if(iterations > 0) {
    var num_integers int
    inputreader := bufio.NewReader(os.Stdin)

    // Input num_integers
    fmt.Scanf("%d\n", &num_integers)

    // Input string og integers
    input, _ := inputreader.ReadString('\n')

    // Remove delimiter "\n"
    input = strings.Split(input, "\n")[0]

    /* Split the string with integers
     * separated by " " (space) to array of strings
     */
    s := strings.Split(input, " ")

    // Get the resultant sum of squares
    result := compute(num_integers - 1, s)
    fmt.Print(result)
    fmt.Print("\n")

    // Move to next interation of sum of squares computation
    sum_square(iterations - 1)
  }
}

/*
 * main() calls sum_sqaure() to compute
 * sum of the squares of positive integers
 * Input: An integer n (1 <= n <= 100)
 */

func main() {

  // Range of n: (1 <= n <= 100)
  var n int
  fmt.Scanf("%d", &n)
  sum_square(n)
}
