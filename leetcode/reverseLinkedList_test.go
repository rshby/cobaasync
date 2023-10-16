package leetcode

import (
	"testing"
)

type NodeList struct {
	Val  int
	Next *NodeList
}

func TestReverseLinkedList(t *testing.T) {

}

func ReverseList(head *NodeList) *NodeList {
	if head == nil || head.Next == nil {
		return head
	}

	reverseListHead := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return reverseListHead
}
