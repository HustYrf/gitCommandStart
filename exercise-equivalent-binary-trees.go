package main

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

func Walk(t *Tree,ch chan int){
	if t == nil{
		return
	}
	Walk(t.Left,ch)
	ch <- t.Value
	Walk(t.Right,ch)
}


func Same(t1,t2 *Tree) bool{
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1,ch1)
	go Walk(t2,ch2)
	for i:=0;i<10;i++{
		x1,y1 := <-ch1,<-ch2
		if x1!=y1{
			return false
		}
	}
	return true
}

func main(){

}