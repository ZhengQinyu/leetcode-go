package _202001

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	num := head.Val
	for head.Next != nil {
		head = head.Next
		num = num<<1 + head.Val
	}
	return num
}
