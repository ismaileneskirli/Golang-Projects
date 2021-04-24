package main

// fmt for printing and other stuff.
import (
	"fmt"
)

// vays of declaring variable, by default they initialized as 0 :
// var n1 int
// var n2 float64
// var (
// 	m1 = 2
// 	m2 = 3
//)

// or simply var number = 200
// strings are mutable in go(meaing you can change them) :

// arrays :
func todo() {
	//var arr[] int
	arr := []int{1, 2, 3, 4}
	arr2 := []string{"salut", "je m'appelle enes", "bonne nuit"}
	fmt.Println(arr)
	arr2 = append(arr2, "au revoir", "funny dudee")
	fmt.Println(arr2)
}

// pointers : what is the use of them ?
// m1 and m2 are pointers
func swap(m1, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}

// structure encapsulation
type Car struct {
	Name    string
	Age     int
	ModelNo int
}

// Funtion to take struct variable as parameter.
func (c Car) Drive() {
	fmt.Println(c.Name + " is on the road !")
}

type Vehicle interface {
	Drive()
	Stop()
}

func main() {
	// assigning like this wont work outside main. : means assign when changing dont use it.
	// m3 := 3
	// m4 := 5
	// fmt.Println(m3 + m4)

	// m5 := "i am mutable"
	// m5 = "muted"
	// m6 := "mutable string"
	// fmt.Println(m5)
	// fmt.Println(strings.Contains(m5, m6)) // checks if first  parameter contains second.
	// fmt.Println(strings.ReplaceAll(m5, "m", "n"))
	// fmt.Println(strings.Split(m6, " "))

	// Arrays
	//todo()

	// pointers
	// p1 := 2
	// ptr := &p1
	// // hexa decimal represantation of adress
	// fmt.Println(ptr)
	// // value that pointer shows
	// fmt.Println(*ptr)
	// m1, m2 := 2, 3
	// swap(&m1, &m2)
	// fmt.Println(m1, m2)

	// Structures, struct func usages.
	// c := Car{
	// 	Name:    "Toyota",
	// 	Age:     1,
	// 	ModelNo: 5,
	// }

	// fmt.Println(c)
	// c.Drive()

	//fmt.Print(quote.Go())
	// name := "enes"
	// number := 123

	// learning the type of variable ...
	// fmt.Printf("%T", c)
	// %v shows variable
	// %9 adds 9 space before variable, % 5 adds 5 space before var ...
	// fmt.Printf("%v is a %T but not %9v", name, name, number)
	// %b shows the binary version of the number
	// fmt.Printf("%b \n", number)

	// Console Input
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Type sth: ")
	// // whatever i write will be string. so need to convert if gets numbers.
	// scanner.Scan()
	// input := scanner.Text()
	// // way of converting string into int
	// number, _ := strconv.ParseInt(input, 10, 64)
	// fmt.Printf("You typed %v \n", number)
}

// go mod init name
// go mod tidy ( for installing dependencies ( packages))
// go run main.go
// go build main.go
