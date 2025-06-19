package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

/*
编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值
*/
func testIntPointer(i *int) {
	*i += 10

}

/*
实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
*/
func testIntsPointer(is *[]int) {
	for i := range *is {
		(*is)[i] *= 2
	}
}

/*
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
*/
func taskScheduler(tasks []func()) {
	ch := make(chan struct{})
	for _, task := range tasks {
		go func(t func()) {
			start := time.Now()
			t()
			fmt.Printf("Task completed in %v\n", time.Since(start))
			ch <- struct{}{}
		}(task)
	}
	for range tasks {
		<-ch
	}
}

/*
定义一个接口Shape，包含方法Area和Perimeter，
然后定义两个结构体Rectangle和Circle实现该接口，最后在主函数中创建这两个形状的实例并计算它们的面积和周长。
*/
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

/*
使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
*/
type Persion struct {
	Name string
	Age  int
}
type Employee struct {
	EmployeeID int
	p          Persion
}

func (e Employee) PrintInfo() {
	fmt.Printf("EmployeeID is %d ,Name is %s,Age is %d\n", e.EmployeeID, e.p.Name, e.p.Age)
}

type Increase struct {
	sync.Mutex
	count int
}

func (i *Increase) Increase() {
	i.Lock()
	i.count++
	i.Unlock()
}

func main() {
	i := 1
	testIntPointer(&i)
	fmt.Println("testIntPointer:", i)

	is := []int{1, 2, 3, 4, 5}
	testIntsPointer(&is)
	fmt.Println("testIntsPointer:", is)

	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				fmt.Println("0-10的基数", i)
			}
		}
	}()

	go func() {
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("0-10的偶数:", i)
			}
		}
	}()

	time.Sleep(1 * time.Second) // 等待协程执行完毕

	tasks := []func(){
		func() {
			fmt.Println("Task 1")
		},
		func() {
			fmt.Println("Task 2")
		},
		func() {
			fmt.Println("Task 3")
		},
	}
	taskScheduler(tasks)
	fmt.Println("All tasks completed")

	r := Rectangle{Width: 5, Height: 3}
	c := Circle{Radius: 2}

	fmt.Printf("Rectangle Area: %.2f, Perimeter: %.2f\n", r.Area(), r.Perimeter())
	fmt.Printf("Circle Area: %.2f, Perimeter: %.2f\n", c.Area(), c.Perimeter())
	fmt.Println("Program completed successfully")

	p := Persion{Name: "Tom", Age: 20}
	e := Employee{EmployeeID: 1, p: p}
	e.PrintInfo()

	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("ch:", <-ch)
		}
	}()

	time.Sleep(1 * time.Second)

	var wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range ch {
			fmt.Println(i)
		}
	}()
	wg.Wait()
	fmt.Println("Channel processing completed")

	inc := Increase{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				inc.Increase()
			}
		}()

	}
	wg.Wait()
	fmt.Println(inc.count)

}
