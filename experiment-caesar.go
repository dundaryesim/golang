// Decipher the quote from Julius Caesar:
// L fdph, L vdz, L frqtxhuhg.
// —Julius Caesar
//Your program will need to shift uppercase and lowercase letters by –3. Remember that
//'a' becomes 'x', 'b' becomes 'y', and 'c' becomes 'z', and likewise for uppercase letters.

package main

import "fmt"

func main() {
	message := "L fdph, L vdz, L frqtxhuhg."

	for r := 0; r < len(message); r++ {
		character := message[r]
		if character >= 'a' && character <= 'z' {
			character -= 3
		} else if character >= 'A' && character <= 'Z' {
			character -= 3
		}

		fmt.Printf("%c", character)
	}

}
