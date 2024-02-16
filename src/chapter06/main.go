//package main
//
//import "fmt"
//
//func main() {
//	Function1()
//}
//
//func Function1() {
//	fmt.Printf("In Function1 at the top\n")
//	defer Function2()
//	fmt.Printf("In Function1 at the bottom!\n")
//}
//
//func Function2() {
//	fmt.Printf("Function2: Deferred until the end of the calling function!")
//}

// defer_dbconn.go
//package main
//
//import "fmt"
//
//func main() {
//	doDBOperations()
//}
//
//func connectToDB() {
//	fmt.Println("ok, connected to db")
//}
//
//func disconnectFromDB() {
//	fmt.Println("ok, disconnected from db")
//}
//
//func doDBOperations() {
//	connectToDB()
//	fmt.Println("Defering the database disconnect.")
//	defer disconnectFromDB() //function called here with defer
//	fmt.Println("Doing some DB operations ...")
//	fmt.Println("Oops! some crash or network error ...")
//	fmt.Println("Returning from function here!")
//	return //terminate the program
//	// deferred function executed here just before actually returning, even if there is a return or abnormal termination before
//}

// defer_logvalues.go
//package main
//
//import (
//	"io"
//	"log"
//)
//
//func func1(s string) (n int, err error) {
//	defer func() {
//		log.Printf("func1(%q) = %d, %v", s, n, err)
//	}()
//	return 7, io.EOF
//}
//
//func main() {
//	func1("Go")
//}

//package main
//
//import "fmt"
//
//func trace(s string)   { fmt.Println("entering:", s) }
//func untrace(s string) { fmt.Println("leaving:", s) }
//
//func a() {
//	trace("a")
//	defer untrace("a")
//	fmt.Println("in a")
//}
//
//func b() {
//	trace("b")
//	defer untrace("b")
//	fmt.Println("in b")
//	a()
//}
//func main() {
//	b()
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	result := 0
//	start := time.Now()
//	for i := 0; i <= 25; i++ {
//		result = fibonacci(i)
//		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
//	}
//	end := time.Now()
//	delta := end.Sub(start)
//	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
//}
//
//func fibonacci(n int) (res int) {
//	if n <= 1 {
//		res = 1
//	} else {
//		res = fibonacci(n-1) + fibonacci(n-2)
//	}
//	return
//}

//
///* Output:
//fibonacci(0) is: 1
//fibonacci(1) is: 1
//fibonacci(2) is: 2
//fibonacci(3) is: 3
//fibonacci(4) is: 5
//fibonacci(5) is: 8
//fibonacci(6) is: 13
//fibonacci(7) is: 21
//fibonacci(8) is: 34
//fibonacci(9) is: 55
//fibonacci(10) is: 89
//fibonacci(11) is: 144
//fibonacci(12) is: 233
//fibonacci(13) is: 377
//fibonacci(14) is: 610
//fibonacci(15) is: 987
//fibonacci(16) is: 1597
//fibonacci(17) is: 2584
//fibonacci(18) is: 4181
//fibonacci(19) is: 6765
//fibonacci(20) is: 10946
//fibonacci(21) is: 17711
//fibonacci(22) is: 28657
//fibonacci(23) is: 46368
//fibonacci(24) is: 75025
//fibonacci(25) is: 121393
//longCalculation took this amount of time: 3.0001ms
//*/

// fibonacci_memoization.go
//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//const LIM = 41
//
//var fibs [LIM]uint64
//
//func main() {
//	var result uint64 = 0
//	start := time.Now()
//	for i := 0; i < LIM; i++ {
//		result = fibonacci(i)
//		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
//	}
//	end := time.Now()
//	delta := end.Sub(start)
//	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
//}
//
//func fibonacci(n int) (res uint64) {
//	// memoization: check if fibonacci(n) is already known in array:
//	if fibs[n] != 0 {
//		res = fibs[n]
//		return
//	}
//	if n <= 1 {
//		res = 1
//	} else {
//		res = fibonacci(n-1) + fibonacci(n-2)
//	}
//	fibs[n] = res
//	return
//}

//
///*
//Output: LIM=40:
//normal (fibonacci.go): the calculation took this amount of time: 4.730270 s
//     with memoization: the calculation took this amount of time: 0.001000 s
//*/

// filter_factory.go
//package main
//
//import "fmt"
//
//type flt func(int) bool
//type sliceSplit func([]int) ([]int, []int)
//
//func isOdd(integer int) bool {
//	if integer%2 == 0 {
//		return false
//	}
//	return true
//}
//
//func isBiggerThan4(integer int) bool {
//	if integer > 4 {
//		return true
//	}
//	return false
//}
//
//func filterFactory(f flt) sliceSplit {
//	return func(s []int) (yes, no []int) {
//		for _, val := range s {
//			if f(val) {
//				yes = append(yes, val)
//			} else {
//				no = append(no, val)
//			}
//		}
//		return
//	}
//}
//
//func main() {
//	s := []int{1, 2, 3, 4, 5, 7}
//	fmt.Println("s = ", s)
//	oddEvenFunction := filterFactory(isOdd)
//	odd, even := oddEvenFunction(s)
//	fmt.Println("odd = ", odd)
//	fmt.Println("even = ", even)
//	//separate those that are bigger than 4 and those that are not.
//	bigger, smaller := filterFactory(isBiggerThan4)(s)
//	fmt.Println("Bigger than 4: ", bigger)
//	fmt.Println("Smaller than or equal to 4: ", smaller)
//}

/*
s =  [1 2 3 4 5 7]
odd =  [1 3 5 7]
even =  [2 4]
Bigger than 4:  [5 7]
Smaller than or equal to 4:  [1 2 3 4]
*/

//package main
//
//import "fmt"
//
//func main() {
//	var f = Adder()
//	fmt.Print(f(1), " - ")
//	fmt.Print(f(20), " - ")
//	fmt.Print(f(300))
//}
//
//func Adder() func(int) int {
//	var x int
//	return func(delta int) int {
//		x += delta
//		return x
//	}
//}
//package main
//
//import "fmt"
//
//func main() {
//	f()
//}
//
//func f() {
//	for i := 0; i < 4; i++ {
//		g := func(i int) { fmt.Printf("%d ", i) }
//		g(i)
//		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
//	}
//
//}

// function_parameter.go
//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	callback(1, Add)
//}
//
//func Add(a, b int) {
//	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
//}
//
//func callback(y int, f func(int, int)) {
//	f(y, 2) // this becomes Add(1, 2)
//}

//package main
//
//import "fmt"
//
//func main() {
//	// make an Add2 function, give it a name p2, and call it:
//	p2 := Add2()
//	fmt.Printf("Call Add2 for 2 gives: %v\n", p2(2))
//	// make a special Adder function, a gets value 2:
//	TwoAdder := Adder(4, 3)
//	fmt.Printf("The result is: %v\n", TwoAdder(3))
//}
//
//func Add2() func(b int) int {
//	return func(b int) int {
//		return b + 2
//	}
//}
//
//func Adder(a, c int) func(b int) int {
//	return func(b int) int {
//		return a + b + c
//	}
//}

/* Output:
Call Add2 for 2 gives: 4
The result is: 4
*/

//package main
//
//import "fmt"
//
//func main() {
//	var min, max int
//	min, max = MinMax(78, 65)
//	fmt.Printf("Minimum is: %d, Maximum is: %d\n", min, max)
//}
//
//func MinMax(a int, b int) (min int, max int) {
//	if a < b {
//		min = a
//		max = b
//	} else { // a = b or a < b
//		min = b
//		max = a
//	}
//	// return min, max
//	return
//}

// mut_recurs.go
//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
//	fmt.Printf("%d is odd: is %t\n", 17, odd(17))   // 17 is odd: is true
//	fmt.Printf("%d is odd: is %t\n", 18, odd(18))   // 18 is odd: is false
//}
//
//func even(nr int) bool {
//	if nr == 0 {
//		return true
//	}
//	return odd(RevSign(nr) - 1)
//}
//
//func odd(nr int) bool {
//	if nr == 0 {
//		return false
//	}
//	return even(RevSign(nr) - 1)
//}
//
//func RevSign(nr int) int {
//	if nr < 0 {
//		return -nr
//	}
//	return nr
//}

//package main
//
//import (
//	"fmt"
//)
//
//func Multiply(a, b int, reply *int) {
//	*reply = a * b
//}
//
//func main() {
//	n := 0
//	reply := &n
//	Multiply(10, 5, reply)
//	fmt.Println("Multiply:", *reply) // Multiply: 50
//	fmt.Println("Multiply:", n)      // Multiply: 50
//}

//package main
//
//import (
//	"fmt"
//)
//
//func f() (ret int) {
//	defer func() {
//		ret++
//	}()
//	return 1
//}
//
//func main() {
//	fmt.Println(f())
//}

// Output: 2

//package main
//
//import "fmt"
//
//func main() {
//	var i1 int = MultiPly3Nums(2, 5, 6)
//	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", i1)
//	// fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
//}
//
//func MultiPly3Nums(a int, b int, c int) int {
//	// var product int = a * b * c
//	// return product
//	return a * b * c
//}

//package main
//
//import "fmt"
//
//func main() {
//	x := Min(1, 3, 2, 0)
//	fmt.Printf("The minimum is: %d\n", x)
//	slice := []int{7, 9, 3, 5, 1}
//	x = Min(slice...)
//	fmt.Printf("The minimum in the slice is: %d", x)
//}
//
//func Min(s ...int) int {
//	if len(s) == 0 {
//		return 0
//	}
//	min := s[0]
//	for _, v := range s {
//		if v < min {
//			min = v
//		}
//	}
//	return min
//}

/*
The minimum is: 0
The minimum in the slice is: 1
*/
//package main
//
//import "fmt"
//
//func main() {
//	printRecursive(1)
//}
//
//func printRecursive(i int) {
//	if i > 10 {
//		return
//	}
//	printRecursive(i + 1)
//	fmt.Printf("%d ", i)
//}

// compose.go
//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func Compose(f, g func(x float64) float64) func(x float64) float64 {
//	return func(x float64) float64 { // closure
//		return f(g(x))
//	}
//}
//
//func main() {
//	fmt.Print(Compose(math.Sin, math.Cos)(0.5)) // output: 0.7691963548410085
//}

// error_returnval.go
//package main
//
//import (
//	"errors"
//	"fmt"
//	"math"
//)
//
//func main() {
//	fmt.Print("First example with -1: ")
//	ret1, err1 := MySqrt(-1)
//	if err1 != nil {
//		fmt.Println("Error! Return values are: ", ret1, err1)
//	} else {
//		fmt.Println("It's ok! Return values are: ", ret1, err1)
//	}
//
//	fmt.Print("Second example with 5: ")
//	//you could also write it like this
//	if ret2, err2 := MySqrt(5); err2 != nil {
//		fmt.Println("Error! Return values are: ", ret2, err2)
//	} else {
//		fmt.Println("It's ok! Return values are: ", ret2, err2)
//	}
//	// named return variables:
//	fmt.Println(MySqrt2(5))
//}
//
//func MySqrt(f float64) (float64, error) {
//	//return an error as second parameter if invalid input
//	if f < 0 {
//		return float64(math.NaN()), errors.New("i won't be able to do a sqrt of negative number")
//	}
//	//otherwise use default square root function
//	return math.Sqrt(f), nil
//}
//
//// MySqrt2 name the return variables - by default it will have 'zero-ed' values i.e. numbers are 0, string is empty, etc.
//func MySqrt2(f float64) (ret float64, err error) {
//	if f < 0 {
//		//then you can use those variables in code
//		ret = float64(math.NaN())
//		err = errors.New("i won't be able to do a sqrt of negative number")
//	} else {
//		ret = math.Sqrt(f)
//		//err is not assigned, so it gets default value nil
//	}
//	//automatically return the named return variables ret and err
//	return
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	for i := uint64(0); i < uint64(30); i++ {
//		fmt.Printf("Factorial of %d is %d\n", i, Factorial(i))
//	}
//}
//
///* unnamed return variables:
//func Factorial(n uint64) uint64 {
//	if n > 0 {
//		return n * Factorial(n-1)
//	}
//	return 1
//}
//*/
//
//// Factorial named return variables:
//func Factorial(n uint64) (fac uint64) {
//	fac = 1
//	if n > 0 {
//		fac = n * Factorial(n-1)
//		return
//	}
//	return
//}

package main

import (
	"fmt"
)

func main() {
	pos := 4
	result, pos := fibonacci(pos)
	fmt.Printf("the %d-th fibonacci number is: %d\n", pos, result)
	pos = 10
	result, pos = fibonacci(pos)
	fmt.Printf("the %d-th fibonacci number is: %d\n", pos, result)
}

func fibonacci(n int) (val, pos int) {
	if n <= 1 {
		val = 1
	} else {
		v1, _ := fibonacci(n - 1)
		v2, _ := fibonacci(n - 2)
		val = v1 + v2
	}
	pos = n
	return
}
