package letcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var list3, ans *ListNode
	if list1.Val < list2.Val {
		list3 = list1
		list1 = list1.Next
	} else {
		list3 = list2
		list2 = list2.Next
	}
	ans = list3

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			if list1.Val < list2.Val {
				list3.Next = list1
				list1 = list1.Next
			} else {
				list3.Next = list2
				list2 = list2.Next
			}
			list3 = list3.Next
		}
	}
	for list1 != nil {
		list3.Next = list1
		list1 = list1.Next
		list3 = list3.Next
	}
	for list2 != nil {
		list3.Next = list2
		list2 = list2.Next
		list3 = list3.Next
	}
	return ans
}
