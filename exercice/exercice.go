package exercice

import (
	"fmt"
	"time"
)
import "sort"

var x, y int
var (
	a int
	b bool
)

func main() {
	defer func() {
		fmt.Println("end program")
	}()
	us := make(map[string]*User)

	u := User{
		ID:   123,
		name: "bo",
		age:  20,
	}

	var i int
	var f float32
	var b bool
	var s string

	us[u.name] = &u
	fmt.Println("Hello World")
	fmt.Println(u)
	fmt.Println(us)
	fmt.Printf("variable : %v %v %v %q \n", i, f, b, s)
	fmt.Printf("variable : %v %v %v %v \n", a, b, x, y)
	fmt.Printf("address of i is %x \n", &i)


	a := map[int]int{0:1}
	for i := 0; i < len(a); i++  {
		fmt.Printf("value of map is %v \n", a[i])
	}

	m1 := map[int]int{0:1}
	var keys []int
	for k := range m1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// Exercice : les dictionnaires map
	// 1.
	// m := map[string]int{}
	m := make(map[string]int)	// allocation du mémoire
	// 2.
	var letter rune = 'a'
	// 3.
	for i := 0; i < 26; i++  {
		m[string(letter)] = i
		// auto incrémentez la valeur
		letter++
	}
	fmt.Printf("value of rune is %v \n ",  m["w"])
	// fmt.Printf("value of auto increment rune is %v \n ",  r++)


	fmt.Println(Books{"Go lang", "Bo", "Go formation", 6495407})
	fmt.Println(Books{title: "Go lang", author: "bo", subject: "Go formation", bookId: 6495407})

	var Book1 Books        /* 声明 Book1 为 Books 类型 */
	var Book2 Books        /* 声明 Book2 为 Books 类型 */

	Book1.title = "Go"
	Book1.author = "bo"
	Book1.subject = "Go formation"
	Book1.bookId = 6495407

	Book2.title = "Python"
	Book2.author = "bo"
	Book2.subject = "Python formation"
	Book2.bookId = 6495700
	printBook(Book1)


	// tableau dynamique
	dynamicTable := make([]int, 6)
	//var newDynamicElement = []
	fmt.Println(dynamicTable)

	fmt.Println("hello", 1, 2.43, 'a')

	aa := []int {1, 2, 3, 4, 5}
	bb := aa[:]
	cc := aa[1:3]
	dd := aa[1:3:4]

	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)

	// exercice : Les tableaux dynamiques
	aaa := []string {"A", "B", "C", "D"}
	fmt.Println(aaa)

	bbb := make([]string, len(aaa))

	copy(bbb, aaa)

	aaa[0] = "Z"

	fmt.Println(aaa)
	fmt.Println(bbb)
	//ssss := [2]int{1, 2, 3}

	// exercice les tableaux
	var numbers = make([]string, 110)
	t := [...]string{9: "a"}
	numbers[9] = "a"

	fmt.Println(numbers[9])
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))
	fmt.Println(t)

	// exercice : type et méthode

	fmt.Println(time.Now())

}

type User struct {
	ID   int
	name string
	age  int
}

func (u *User) String() string {
	return u.name
}


type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func printBook( book Books ) {
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book subject : %s\n", book.subject)
	fmt.Printf( "Book book_id : %d\n", book.bookId)
}

func printSlice (mySlice []string) {
	fmt.Printf(" slice value : ", mySlice)
}

//func (s string) String (i int) myInt {
//
//}

type myInt int32

func (mi myInt) DivideDeProf (n int) myInt {
	return mi / myInt(n)
}

func Divide (x, y myInt) myInt {
	return x/y
}

func Add (x, y myInt) myInt {
	return x+y
}

func Sub (x, y myInt) myInt {
	return x-y
}

func Multiply (x, y myInt) myInt {
	return x*y
}
