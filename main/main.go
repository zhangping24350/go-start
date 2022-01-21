package main

//import "C"
import (
	json "encoding/json"
	"fmt"
	cal "go-start/07_test"
	"log"
	"net/http"
	"reflect"
)

func main() {

	//C.SayHello(C.CString("Hello, World\n"))
	//basicgrammar.VariableDeclarations("s",19)

	//name := "王五"
	//fmt.Printf(name)
	//fmt.Printf("ss")
	//http.HandleFunc(
	//	"/test",
	//	getHeader,
	//)
	////http.ListenAndServe(":8080", nil)
	//s := new(Student)
	//s.GetStudent("ss")
	//t := new(Teacher)
	//var ss int8 = 7
	//fmt.Printf("%s",ss)
	//t.nsu = 28278278228227811111111
	//t.GetStudent("fdsh")
	//
	//fmt.Printf("\n\n\nstart to test dubbo")
	//
	//c := make(chan string, 2)
	//
	//go func() {
	//	fmt.Printf("\n\n\nstart to test dubbo 22222")
	//
	//	c <- "sjson"
	//
	//	fmt.Printf("\n\n\nstart to test dubbo 333333")
	//}()
	//
	//go func() {
	//	fmt.Printf("\n\n\nstart to test dubbo 222225555")
	//	time.Sleep(100)
	//	c <- "uuuuryysattsggs"
	//
	//	fmt.Printf("\n\n\nstart to test dubbo 3333336666")
	//}()
	//
	//fmt.Printf("\n\n\nstart to test dubbo 111111")
	//
	//fmt.Printf("\n\nffffffff 111%s",<-c)
	//
	//fmt.Printf("\n\nffffffff222%s",<-c)
	//close(c)
	//C.puts(C.CString("cgo"))

	//res := "ss"
	//typeName := reflect.TypeOf(res)
	//defer func() {
	//	fmt.Println("defer")
	//}()
	//fmt.Println(typeName)
	//fmt.Println(reflect.ValueOf(res))
	//go func() {
	//	fmt.Println("go firdt")
	//	defer fmt.Println("go defer")
	//	panic("goruntine error")
	//}()

	//panic("error")
	//fmt.Println(reflect.Append(reflect.ValueOf(res), reflect.ValueOf("sss")))
	//i := 1
	//v := reflect.ValueOf(&i)
	//v.Elem().SetInt(1778)
	//
	//fmt.Println(i)
	//time.Sleep(1 * time.Second)

	//teacher := Teacher{"ss", 16, "skskks", 22}
	//teacher.GetStudent("heel")

	test := `{"artifact":[{"ext":"demoString","total":0,"src":"demoString","name":"demoString","complete":0,"nofile":false}],"standard_parts":[{"ext":"demoString","total":0,"src":"demoString","name":"demoString","complete":0,"nofile":false}]}`
	wo := &Workpiece{}
	if err := json.Unmarshal([]byte(test), wo); err != nil {
		log.Fatal(err)
	}
	t2 := reflect.TypeOf(StudentService.GetStudent)
	i := reflect.New(t2).Interface()
	canInterface := reflect.New(t2).CanInterface()
	fmt.Println(i)
	fmt.Println(canInterface)

	of := reflect.ValueOf(cal.Add)
	fmt.Println(t2)
	fmt.Println(of)
	var s name = "ss"
	fmt.Sprintf(string(s))
	fmt.Println(wo)
}

func getHeader(rq http.ResponseWriter, rs *http.Request) {
	body := rs.Body

	fmt.Printf("requst : %v", body)
}

type Student struct {
	name   string
	age    int16
	mobile string
}
type Teacher struct {
	name   string
	age    int16
	mobile string
	nsu    uintptr
}

type StudentService interface {
	GetStudent(name string)
}

func (t Teacher) GetStudent(name string) {
	fmt.Printf("name:%s age:%d param: %s", t.name, t.age, name)
	panic("ss")
}

func (s *Student) GetStudent(name string) {
	s2 := new(Student)
	s2.name = "fdds"
	s2.mobile = "fd"

	s3 := Student{name: "nsn", age: 11, mobile: "fd"}
	s3.age = 22

	student := Student{
		name:   "zhangsan",
		age:    18,
		mobile: "1672727",
	}
	fmt.Printf("%s", student)
	type1 := reflect.TypeOf(student)
	fmt.Printf("11 %s", type1.Field(1))
}

type Associate struct {
	Total    int    `json:"total"`
	Complete int    `json:"complete"`
	Src      string `json:"src"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
	Nofile   bool   `json:"nofile"`
}

type AKK interface {
	Get(a int)
}

func (t *Associate) Get(a int) {

}

type (
	name string
	age  int
)

type Workpiece struct {
	StandardParts []Associate `json:"standard_parts"`
	Artifact      []Associate `json:"artifact"`
}
