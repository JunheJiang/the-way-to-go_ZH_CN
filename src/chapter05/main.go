//package main
//
//import "fmt"
//
//func main() {
//	bool1 := true
//	if bool1 {
//		fmt.Printf("The value is true\n")
//	} else {
//		fmt.Printf("The value is false\n")
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	for i := 0; i < 5; i++ {
//		fmt.Printf("This is the %d iteration\n", i)
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	var i int = 5
//
//	for i >= 0 {
//		i = i - 1
//		fmt.Printf("The variable i is now: %d\n", i)
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	var i int = 5
//	for {
//		i = i - 1
//		fmt.Printf("The variable i is now: %d\n", i)
//		if i < 0 {
//			break
//		}
//	}
//}

//package main
//
//func main() {
//	for i := 0; i < 3; i++ {
//		for j := 0; j < 10; j++ {
//			if j > 5 {
//				break
//			}
//			print(j)
//		}
//		print("  ")
//	}
//}

//package main
//
//func main() {
//	for i := 0; i < 10; i++ {
//		if i == 5 {
//			continue
//		}
//		print(i)
//		print(" ")
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//LABEL1:
//	for i := 0; i <= 5; i++ {
//		for j := 0; j <= 5; j++ {
//			if j == 4 {
//				continue LABEL1
//			}
//			fmt.Printf("i is: %d, and j is: %d\n", i, j)
//		}
//	}
//
//}

// for_string.go
//package main
//
//import "fmt"
//
//func main() {
//	str := "Go is a beautiful language!"
//	fmt.Printf("The length of str is: %d\n", len(str))
//	for ix := 0; ix < len(str); ix++ {
//		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
//	}
//	str2 := "日本語"
//	fmt.Printf("The length of str2 is: %d\n", len(str2))
//	for ix := 0; ix < len(str2); ix++ {
//		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
//	}
//}

//package main
//
//func main() {
//	i := 0
//HERE:
//	print(i)
//	i++
//	if i == 5 {
//		return
//	}
//	goto HERE
//}

//package main
//
//import "fmt"
//
//func main() {
//	a := 1
//	goto TARGET // compile error
//	b := 9
//TARGET:
//	b += a
//	fmt.Printf("a is %v *** b is %v", a, b)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var first int = 10
//	var cond int
//
//	if first <= 0 {
//		fmt.Printf("first is less than or equal to 0\n")
//	} else if first > 0 && first < 5 {
//		fmt.Printf("first is between 0 and 5\n")
//	} else {
//		fmt.Printf("first is 5 or greater\n")
//	}
//
//	if cond = 5; cond > 10 {
//		fmt.Printf("cond is greater than 10\n")
//	} else {
//		fmt.Printf("cond is not greater than 10\n")
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	str := "Go is a beautiful language!"
//	fmt.Printf("The length of str is: %d\n", len(str))
//	for pos, char := range str {
//		fmt.Printf("Character on position %d is: %c \n", pos, char)
//	}
//	fmt.Println()
//	str2 := "Chinese: 日本語"
//	fmt.Printf("The length of str2 is: %d\n", len(str2))
//	for pos, char := range str2 {
//		fmt.Printf("character %c starts at byte position %d\n", char, pos)
//	}
//	fmt.Println()
//	fmt.Println("index int(rune) rune    char bytes")
//	for index, rune := range str2 {
//		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
//	}
//}

/* Output:
The length of str is: 27
Character on position 0 is: G
Character on position 1 is: o
Character on position 2 is:
Character on position 3 is: i
Character on position 4 is: s
Character on position 5 is:
Character on position 6 is: a
Character on position 7 is:
Character on position 8 is: b
Character on position 9 is: e
Character on position 10 is: a
Character on position 11 is: u
Character on position 12 is: t
Character on position 13 is: i
Character on position 14 is: f
Character on position 15 is: u
Character on position 16 is: l
Character on position 17 is:
Character on position 18 is: l
Character on position 19 is: a
Character on position 20 is: n
Character on position 21 is: g
Character on position 22 is: u
Character on position 23 is: a
Character on position 24 is: g
Character on position 25 is: e
Character on position 26 is: !

The length of str2 is: 18
character C starts at byte position 0
character h starts at byte position 1
character i starts at byte position 2
character n starts at byte position 3
character e starts at byte position 4
character s starts at byte position 5
character e starts at byte position 6
character : starts at byte position 7
character   starts at byte position 8
character 日 starts at byte position 9
character 本 starts at byte position 12
character 語 starts at byte position 15

index int(rune) rune    char bytes
0       67      U+0043 'C' 43
1       104      U+0068 'h' 68
2       105      U+0069 'i' 69
3       110      U+006E 'n' 6E
4       101      U+0065 'e' 65
5       115      U+0073 's' 73
6       101      U+0065 'e' 65
7       58      U+003A ':' 3A
8       32      U+0020 ' ' 20
9       26085      U+65E5 '日' E6 97 A5
12      26412      U+672C '本' E6 9C AC
15      35486      U+8A9E '語' E8 AA 9E
*/

// package main
//
// import (
//
//	"fmt"
//	"strconv"
//
// )
//
//	func main() {
//		var orig string = "ABC"
//		// var an int
//		var newS string
//		// var err error
//
//		fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
//		// anInt, err = strconv.Atoi(origStr)
//		an, err := strconv.Atoi(orig)
//		if err != nil {
//			fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
//			return
//		}
//		fmt.Printf("The integer is %d\n", an)
//		an = an + 5
//		newS = strconv.Itoa(an)
//		fmt.Printf("The new string is: %s\n", newS)
//	}
//
// package main
//
// import "fmt"
//
//	func main() {
//		var num1 int = 100
//		switch num1 {
//		case 98, 99:
//			fmt.Println("It's equal to 98")
//		case 100:
//			fmt.Println("It's equal to 100")
//		default:
//			fmt.Println("It's not equal to 98 or 100")
//		}
//	}
//package main
//
//import "fmt"
//
//func main() {
//	var num1 int = 7
//	switch {
//	case num1 < 0:
//		fmt.Println("Number is negative")
//	case num1 > 0 && num1 < 10:
//		fmt.Println("Number is between 0 and 10")
//	default:
//		fmt.Println("Number is 10 or greater")
//	}
//}

//package main
//
//func main() {
//	// 1 - use 2 nested for loops
//	for i := 1; i <= 25; i++ {
//		for j := 1; j <= i; j++ {
//			print("G")
//		}
//		println()
//	}
//	// 2 -  use only one for loop and string concatenation
//	str := "G"
//	for i := 1; i <= 25; i++ {
//		println(str)
//		str += "G"
//	}
//}

package main

import (
	"fmt"
)

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}
	fmt.Printf("%v\n", i) //<-- compile error:  undefined i
}
