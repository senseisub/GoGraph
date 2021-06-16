package main

import (
	"fmt"
	"log"
)

//bi-directional graph
type Graph struct {
	internalMap map[int] []int
}

func (graph *Graph) AddEdge(first, second int) {
	if graph.internalMap == nil {
		graph.internalMap = make(map[int] [] int)
	}
	_, containsFirst := graph.internalMap[first]
	_, containsSecond := graph.internalMap[second]
	if !containsFirst {
		graph.internalMap[first] = [] int {second}
	} else{
		firstSlice := graph.internalMap[first]
		graph.internalMap[first] = append(firstSlice, second)
	}
	if !containsSecond {
		graph.internalMap[second] = [] int {first}
	} else{
		secondSlice := graph.internalMap[second]
		graph.internalMap[second] = append(secondSlice, first)
	}
}

func (graph *Graph) BreadthFirst(origin int) {
	visited := make(map[int] bool)
	for key, _ := range graph.internalMap {
        visited[key] = false
    }

	queue := &LinkedList{}
	visited[origin] = true
	queue.Insert(origin);

	for queue.size != 0 {
		origin = queue.PopFront()
		fmt.Print(origin, " ")
		for _, element := range graph.internalMap[origin] {
			if !visited[element]{
				visited[element] = true
				queue.Insert(element)
			}
		}

	}
	fmt.Println()
	
}



func main() {
	newGraph := Graph{}
	newGraph.AddEdge(1,2)
	newGraph.AddEdge(1,3)
	newGraph.AddEdge(1,4)
	newGraph.AddEdge(5,6)
	newGraph.AddEdge(5,7)
	newGraph.AddEdge(5,8)
	newGraph.BreadthFirst(5)
}


//LinkedList Implementation

type ListNode struct {
	data int
	next *ListNode
	prev *ListNode
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

func (list *LinkedList) Insert(x int) {
	ptr := &ListNode{data : x}
	if list.head == nil {
		list.head, list.tail = ptr, ptr
	} else{
		list.tail.next = ptr
		ptr.prev = list.tail
		list.tail = ptr
	}
	list.size++
}

func (list *LinkedList) Get(index int) int{
	place := 0
	curr := list.head
	if curr == nil{
		log.Fatal("ERROR OUT OF BOUNDS")
	}
	if index < 0 {
		return list.head.data
	}
	for curr != nil {
		if place == index{
			return curr.data
		}
		place++
		curr = curr.next
	}
	return list.tail.data
}

func (list *LinkedList) PopEnd() int{
	if list.size == 0 {
		log.Fatal("No more elements")
	}
	list.size--
	if list.head == list.tail {
		temp := list.tail
		list.tail, list.head = nil, nil
		return temp.data
	} else{
		temp := list.tail
		list.tail = list.tail.prev
		list.tail.next = nil
		return temp.data
	}
	return 0
}

func (list *LinkedList) PopFront() int{
	if list.size == 0 {
		log.Fatal("No more elements")
	}
	list.size--
	if list.head == list.tail {
		temp := list.tail
		list.tail, list.head = nil, nil
		return temp.data
	} else{
		temp := list.head
		list.head = list.head.next
		list.head.prev = nil
		return temp.data
	}
	return 0
}