package main

import "fmt"

//_ "github.com/learn/init_order/pkg1" // Adjust the import path as necessary

var m map[string]string = map[string]string{
	"pkg1": "pkg1",
	"pkg2": "pkg2",
	"pkg3": "pkg3",
}

var z1 *int
var s1 *string

var a, b, c int = 1, 2, 3
var e, f, g string

func method1() {
	a, b, c := 4, 5, 6
	e, f, g := "e", "f", "g"
	fmt.Println(a, b, c, e, f, g)

	a = 7
}

var nodeName = struct {
	field1 int
	field2 string
	field3 bool
}{
	field1: 1,
	field2: "2",
	field3: true,
}

func method2(a string, b string) (string, string) {
	a = "a"
	b = "b"

	fmt.Println(a, b, c)
	return "a", "b"
}

type Person struct {
	Name string `json:"name" xml:"name"`
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}
func (p *Person) SetName(name string) {
	p.Name = name
}
func (p *Person) SetAge(age int) {
	p.Age = age
}
func (p Person) GetName() string {
	return p.Name
}
func (p Person) GetAge() int {
	return p.Age
}
func (p Person) GetInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}
func (p Person) GetInfo2() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}
func (p Person) GetInfo3() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func main() {
	p := Person{Name: "John", Age: 30}
	fmt.Println(p.GetInfo())
	p.SetName("Doe")
	fmt.Println(p.GetInfo())
	p.SetAge(35)
	fmt.Println(p.GetInfo())
	fmt.Println("Node Name:", nodeName.field1, nodeName.field2, nodeName.field3)
	fmt.Println("Map:", m)
	fmt.Println("Method1 Output:")

}
