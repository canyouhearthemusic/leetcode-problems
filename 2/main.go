package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	var carry int

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		node := &ListNode{Val: sum % 10}

		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}

	return head
}

func main() {
	//l1 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: nil,
	//		},
	//	},
	//}
	//
	//l2 := &ListNode{
	//	Val: 3,
	//	Next: &ListNode{
	//		Val:  4,
	//		Next: nil,
	//	},
	//}
	//
	//result := addTwoNumbers(l1, l2)
	//
	//fmt.Println(result)

	nums := []int{1, 2, 3, 4, 5, 6, 1, 2, 1}

	result := findDuplicates(nums)

	fmt.Println(result)
}

func findDuplicates(nums []int) []int {
	result := make([]int, 0)
	seen := make(map[int]bool)

	for _, num := range nums {
		if added, exists := seen[num]; exists {
			if !added {
				seen[num] = true
				result = append(result, num)
			}
		} else {
			seen[num] = false
		}
	}

	return result
}
