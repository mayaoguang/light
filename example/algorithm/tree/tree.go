package tree

import "container/list"

//二叉树
type binaryTree struct {
	root   *BinTreeNode //根节点
	height int
	size   int
}

func NewBinaryTree(root *BinTreeNode) *binaryTree {
	return &binaryTree{root: root}
}

//GetSize 获得二叉树总结点数
func (this *binaryTree) GetSize() int {
	return this.root.size
}

//IsEmpty 判断二叉树是否为空
func (this *binaryTree) IsEmpty() bool {
	return this.root != nil
}

//GetRoot 获得二叉树根节点
func (this *binaryTree) GetRoot() *BinTreeNode {
	return this.root
}

//GetHeight 获得二叉树高度，根节点层为0
func (this *binaryTree) GetHeight() int {
	return this.root.height
}

//Find 获得第一个与数据e相等的节点
func (this *binaryTree) Find(e interface{}) *BinTreeNode {
	if this.root == nil {
		return nil
	}
	p := this.root
	return isEqual(e, p)
}

func isEqual(e interface{}, node *BinTreeNode) *BinTreeNode {
	if e == node.GetData() {
		return node
	}

	if node.HasLChild() {
		lp := isEqual(e, node.GetLChild())
		if lp != nil {
			return lp
		}
	}

	if node.HasRChild() {
		rp := isEqual(e, node.GetRChild())
		if rp != nil {
			return rp
		}

	}

	return nil
}

//PreOrder 先序遍历，并保存在链表里
func (this *binaryTree) PreOrder() *list.List {
	traversal := list.New()
	preOrder(this.root, traversal)
	return traversal
}

func preOrder(rt *BinTreeNode, l *list.List) {
	if rt == nil {
		return
	}
	l.PushBack(rt)
	preOrder(rt.GetLChild(), l)
	preOrder(rt.GetRChild(), l)
}

//InOrder 中序遍历，并保存在链表里
func (this *binaryTree) InOrder() *list.List {
	traversal := list.New()
	inOrder(this.root, traversal)
	return traversal
}

func inOrder(rt *BinTreeNode, l *list.List) {
	if rt == nil {
		return
	}
	inOrder(rt.GetLChild(), l)
	l.PushBack(rt)
	inOrder(rt.GetRChild(), l)
}

//PostOrder 后序遍历，并保存在链表里
func (this *binaryTree) PostOrder() *list.List {
	traversal := list.New()
	postOrder(this.root, traversal)
	return traversal
}

func postOrder(rt *BinTreeNode, l *list.List) {
	if rt == nil {
		return
	}
	postOrder(rt.GetLChild(), l)
	postOrder(rt.GetRChild(), l)
	l.PushBack(rt)
}
