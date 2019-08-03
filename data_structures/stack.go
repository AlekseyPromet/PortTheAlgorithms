package main

import (
	"fmt"
	"os"
)


func main() {
	push(1)
	push(2)
	push(-10)
	push(10)

	fmt.Println(sum(head))
	pop()
	fmt.Println(sum(head))
	pop()
	fmt.Println(sum(head))
	pop()
	pop()
	pop()
	pop()
	pop()
}

type Node struct {
	data int
	next *Node
	pre  *Node
}

var head *Node
var tmp *Node
var count int


func push(x int) {
	if head == nil {
		head = &Node{0, nil, nil}
		head.data = x

	} else {
		tmp = &Node{0, nil, nil}
		tmp.data = x
		tmp.pre = head
		tmp.next = nil
		head = tmp
	}
	count++
}


func pop() (returnData int){
	if head == nil {
		os.Exit(1)
	} else {
		returnData = head.data
		if head.pre == nil {
			head = nil
		} else {
			head = head.pre
		}
	}
	return
}

func sum(last *Node) (total int) {
	if last == nil{
		return
	}
	for last != nil {
		total += last.data
		last = last.pre
	}
	return
}
