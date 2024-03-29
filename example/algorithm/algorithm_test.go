package algorithm

import (
	"fmt"
	"light/example/algorithm/dp"
	"light/example/algorithm/list"
	"light/example/algorithm/tree"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tmp := make(map[int]int)
	fmt.Println(dp.Fibonacci(tmp, 10))
}

func TestLengthOfLIS(t *testing.T) {
	fmt.Println(dp.LengthOfASC([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
func TestCommonSub(t *testing.T) {
	fmt.Println(dp.LongestCommonSubsequence("abc", "cdeabddc"))
}

func TestTree(t *testing.T) {
	a := tree.NewBinTreeNode(1)
	tree1 := tree.NewBinaryTree(a)
	a.SetLChild(tree.NewBinTreeNode(2))
	a.SetRChild(tree.NewBinTreeNode(5))
	a.GetLChild().SetRChild(tree.NewBinTreeNode(3))
	a.GetLChild().GetRChild().SetLChild(tree.NewBinTreeNode(4))
	a.GetRChild().SetLChild(tree.NewBinTreeNode(6))
	a.GetRChild().SetRChild(tree.NewBinTreeNode(7))
	a.GetRChild().GetLChild().SetRChild(tree.NewBinTreeNode(9))
	a.GetRChild().GetRChild().SetRChild(tree.NewBinTreeNode(8))

	l := tree1.InOrder() //中序遍历
	for e := l.Front(); e != nil; e = e.Next() {
		obj, _ := e.Value.(*tree.BinTreeNode)
		fmt.Printf("data:%v\n", obj.GetData())
	}
	result := tree1.Find(6)
	fmt.Println(result)
}

func TestIsSubsequence(t *testing.T) {
	fmt.Println(isSubsequence("b", "c"))
}

func TestReversePrefix(t *testing.T) {
	fmt.Println(ReversePrefix("abcdefd", 'x'))
}

func TestMaximumGap(t *testing.T) {
	fmt.Println(MaximumGap([]int{6, 6, 10, 1}))
}

func TestAllSub(t *testing.T) {
	fmt.Println(dp.Permutation("abc"))
}

func TestList(t *testing.T) {
	l4 := list.ListNode{Val: 3, Next: nil}
	l3 := list.ListNode{Val: 1, Next: &l4}
	l2 := list.ListNode{Val: 2, Next: &l3}
	l1 := list.ListNode{Val: 4, Next: &l2}
	head := list.ListNode{Next: &l1}
	n := list.SortList(&head)
	node := n.Next
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func TestMaxProfit(t *testing.T) {
	fmt.Println(dp.MaxProfit([]int{6, 1, 3, 2, 4, 5}))
}
