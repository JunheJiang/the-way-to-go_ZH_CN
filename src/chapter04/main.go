//alias
//package main
//
//import fm "fmt" // alias
//
//func main() {
//	fm.Println("hello, world")
//}

//package main
//
//import "fmt"

//casting
//import "fmt"
//
//func main() {
//	var n int16 = 34
//	var m int32
//
//	// compiler error: cannot use n (type int16) as type int32 in assignment
//	//m = n
//	m = int32(n)
//
//	fmt.Printf("32 bit int is:  %d\n", m)
//	fmt.Printf("16 bit int is:  %d\n", n)
//}

// char.go
//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var ch int = '\u0041'
//	var ch2 int = '\u03B2'
//	var ch3 int = '\U00101234'
//	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3)
//	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3)
//	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3)
//	fmt.Printf("%U - %U - %U", ch, ch2, ch3)
//}
//
///* Ouput:
//65 - 946 - 1053236
//A - β - 􁈴
//41 - 3B2 - 101234
//U+0041 - U+03B2 - U+101234
//*/
//
///* Output:
//32 bit int is:  34
//16 bit int is:  34
//*/

//package main

//import (
//	"fmt"
//	"strings"
//)
//
//func main() {
//	var str string = "Hello, how is it going, Hugo?"
//	var manyG = "gggggggggg"
//
//	fmt.Printf("Number of H's in %s is: ", str)
//	fmt.Printf("%d\n", strings.Count(str, "H"))
//
//	fmt.Printf("Number of double g's in %s is: ", manyG)
//	fmt.Printf("%d\n", strings.Count(manyG, "gg"))
//}

// callback
//package main
//
//var a string
//
//func main() {
//	a = "G"
//	print(a)
//	f1()
//}
//
//func f1() {
//	a := "O"
//	print(a)
//	f2()
//}
//
//func f2() {
//	print(a)
//}

// scope
//package main
//
//var a = "G"
//
//func main() {
//	n()
//	m()
//	n()
//}
//func n() {
//	print(a)
//}
//func m() {
//	a = "O"
//	print(a)
//}

//package main
//
//import (
//	"fmt"
//	"os"
//	"runtime"
//)
//
//func main() {
//	var goos string = runtime.GOOS
//	fmt.Printf("The operating system is: %s\n", goos)
//	path := os.Getenv("PATH")
//	fmt.Printf("Path is %s\n", path)
//}

// local scope
//package main
//
//var a = "G"
//
//func main() {
//	n()
//	m()
//	n()
//}
//func n() { print(a) }
//func m() {
//	a := "O"
//	print(a)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var i1 = 5
//	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
//
//	var intP *int
//	intP = &i1
//	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
//}

//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//func main() {
//	var str string = "This is an example of a string"
//
//	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
//	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
//}
//
//// Output: T/F? Does the string "This is an example of a string" have prefix Th? true

// random
//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//func main() {
//	for i := 0; i < 10; i++ {
//		a := rand.Int()
//		fmt.Printf("%d / ", a)
//	}
//	for i := 0; i < 5; i++ {
//		r := rand.Intn(8)
//		fmt.Printf("%d / ", r)
//	}
//	fmt.Println()
//	times := int64(time.Now().Nanosecond())
//	//
//	rand.Seed(times)
//
//	for i := 0; i < 10; i++ {
//		fmt.Printf("%2.2f / ", 100*rand.Float32())
//	}
//}

/* Output:
134020434 / 1597969999 / 1721070109 / 2068675587 / 1237770961 / 220031192 / 2031484958 / 583324308 / 958990240 / 413002649 / 6 / 7 / 2 / 1 / 0 /
22.84 / 10.12 / 44.32 / 58.58 / 15.49 / 12.23 / 30.16 / 88.48 / 34.26 / 27.18 /
*/

//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//func main() {
//	var origS string = "Hi there! "
//	var newS string
//
//	newS = strings.Repeat(origS, 3)
//	fmt.Printf("The new repeated string is: %s\n", newS)
//}

// str conversion
//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//func main() {
//	var orig string = "666"
//	var an int
//	var newS string
//
//	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
//
//	an, _ = strconv.Atoi(orig)
//	fmt.Printf("The integer is: %d\n", an)
//	an = an + 5
//	newS = strconv.Itoa(an)
//	fmt.Printf("The new string is: %s\n", newS)
//}

//package main
//
//import "fmt"
//
//func main() {
//	s := "good bye"
//	var p *string = &s
//	*p = "ciao"
//
//	//fmt.Printf("Here is the pointer p: %p\n", s)
//	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
//	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
//	fmt.Printf("Here is the string  s: %s\n", s)  // prints same string
//}

//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//func main() {
//	str := "The quick brown fox jumps over the lazy dog"
//	//to slice
//	sl := strings.Fields(str)
//	fmt.Printf("Splitted in slice: %v\n", sl)
//	for _, val := range sl {
//		fmt.Printf("%s - ", val)
//	}
//	fmt.Println()
//	str2 := "GO1|The ABC of Go|25"
//	sl2 := strings.Split(str2, "|")
//	fmt.Printf("Splitted in slice: %v\n", sl2)
//	for _, val := range sl2 {
//		fmt.Printf("%s - ", val)
//	}
//	fmt.Println()
//	str3 := strings.Join(sl2, ";")
//	fmt.Printf("sl2 joined by ;: %s\n", str3)
//}

// testcrash.go
//// compiles , but crashes
//package main
//
//func main() {
//	var p *int = nil
//	*p = 0
//
//}
//
//// in Windows: stops only with: <exit code="-1073741819" msg="process crashed"/>
//// runtime error: invalid memory address or nil pointer dereference
//package main
//
//func main() {
//	var p *int = nil
//	*p = 0
//
//}
//
//// in Windows: stops only with: <exit code="-1073741819" msg="process crashed"/>
//// runtime error: invalid memory address or nil pointer dereference

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//var week time.Duration
//
//func main() {
//	t := time.Now()
//	fmt.Println(t)                                              // Wed Dec 21 09:52:14 +0100 RST 2011
//	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year()) // 21.12.2011
//	t = time.Now().UTC()
//	fmt.Println(t)          // Wed Dec 21 08:52:14 +0000 UTC 2011
//	fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
//	// calculating times:
//	week = 60 * 60 * 24 * 7 * 1e9 // must be in nano sec
//	weekFromNow := t.Add(week)
//	fmt.Println(weekFromNow) // Wed Dec 28 08:52:14 +0000 UTC 2011
//	// formatting times:
//	fmt.Println(t.Format(time.RFC822))         // 21 Dec 11 0852 UTC
//	fmt.Println(t.Format(time.ANSIC))          // Wed Dec 21 08:56:34 2011
//	fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
//	s := t.Format("20060102")
//	fmt.Println(t, "=>", s) // Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221
//}

// upper lower letter
//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//func main() {
//	var orig string = "Hey, how are you George?"
//	var lower string
//	var upper string
//
//	fmt.Printf("The original string is: %s\n", orig)
//	lower = strings.ToLower(orig)
//	fmt.Printf("The lowercase string is: %s\n", lower)
//	upper = strings.ToUpper(orig)
//	fmt.Printf("The uppercase string is: %s\n", upper)
//}

//package main
//
//import "fmt"
//
//type TZ int
//
//func main() {
//	var a, b TZ = 3, 4
//	c := a + b
//	fmt.Printf("c has the value: %d", c)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var a int
//	var b int32
//	a = 15
//	b = int32(a + a) // compiler error
//	fmt.Println(b)
//	b = b + 5 // ok: 5 is a constant
//	fmt.Println(b)
//}

//package main
//
//import (
//	"fmt"
//	"unicode/utf8"
//)
//
//func main() {
//	// count number of characters:
//	str1 := "asSASA ddd dsjkdsjs dk"
//	fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))
//	fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))
//	str2 := "asSASA ddd dsjkdsjsこん dk"
//	fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
//	fmt.Printf("The number of characters in string str2 is %d", utf8.RuneCountInString(str2))
//}

package main

import (
	"awesomeProject/src/chapter04/trans"
	"fmt"
)

var twoPi = trans.PI // decl computes twoPi

func main() {
	fmt.Printf("2*Pi = %g\n", twoPi) // 2*Pi = 6.283185307179586
}
