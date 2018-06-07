// current package name,must first line
package main

// import other package
import "fmt"

// import joker "fmt"  change package name
/*  . -> is omitted package name,don't use, if need many packages*/
import (
	"fmt"
	"os"
	"string"
	"time"
)

// define constatn
const PI = 3.14

// golobal varialbe define
var name = "gopher"

// normal type statment
type newType int

// structure statment
type gopher struct{}

// inter face statment
type golang interface{}

// define main function,it's a start processing,
// and have package main must
func main() {
	Println("Hello World!")
	fmt.Println("Hello World") //use improt package, crash if not use some packages very important
}
