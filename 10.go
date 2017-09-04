package main

import "fmt"
import "time"
/*
func printn(id int, c chan int){
	for i:=0; i < 10; i++{
		fmt.Println(id, ":", i)
	}
	c <- 0
}
*/

func fun1(c chan string){
	for {
		c <- "from fun1"
		time.Sleep(time.Second * 2)
	}
}

func fun2(c chan string){
	for {
		c <- "from fun2"
		time.Sleep(time.Second * 2)
	}
}

func main(){
	c1 := make(chan string)
	c2 := make(chan string)
	
	go fun1(c1)
	go fun2(c2)
	
	for {
	 select {
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(msg2)
	 }
	}
	/*
	go printn(1, c)
	fmt.Println("waiting...")
	fmt.Println("waiting...")
	fmt.Println("waiting...")
	fmt.Println("waiting...")
	fmt.Println("waiting...")
	fmt.Println("waiting...")
	//time.Sleep(time.Second * 3)	
	<- c
	*/
}
