/* 
 * EclipseCon Europe Che Challenge
 * This program prints the numbers from 1 to 100, with every multiple of 3
 * replaced by "Fizz", and every multiple of 5 replaced by "Buzz". Numbers
 * divisible by both are replaced by "FizzBuzz". Otherwise, the program
 * prints the number.
 *
 * Your mission, if you choose to accept it, is to change this program so that,
 * in addition to the above, it prints "Fizz" for any number which includes a 3,
 * and prints "Buzz" for any number including a 5. After modification, for
 * example, the number 53 should be replaced by "FizzBuzz".
 */

package main
 
import "fmt"
 
func main() {
    for i := 1; i <= 100; i++ {
    	var fizz, buzz bool = false, false;
    	
		if (i%3 == 0) {
			fizz = true;
		}
		
		if (i%5 == 0) {
			buzz = true;
		}


		if (fizz && buzz) {
            fmt.Println("FizzBuzz")
		} else if fizz {
            fmt.Println("Fizz")
		} else if (buzz) {
            fmt.Println("Buzz")
		} else {
            fmt.Println(i)
		}
    }
}