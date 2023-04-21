package tree

import (
	"math"
)

//BinTreeNode 二叉树节点
type BinTreeNode struct {
	data   interface{}  // 数据域
	parent *BinTreeNode // 父节点
	lChild *BinTreeNode // 左孩子
	rChild *BinTreeNode // 右孩子
	height int          // 以该结点为根的子树的高度
	size   int          // 该结点子孙数(包括结点本身)
}

func NewBinTreeNode(e interface{}) *BinTreeNode {
	return &BinTreeNode{data: e, size: 1}
}

//GetData 获得节点数据
func (slf *BinTreeNode) GetData() interface{} {
	if slf == nil {
		return nil
	}
	return slf.data
}

//SetData 设置节点数据
func (slf *BinTreeNode) SetData(e interface{}) {
	slf.data = e
}

//HasParent 判断是否有父亲
func (slf *BinTreeNode) HasParent() bool {
	return slf.parent != nil
}

//GetParent 获得父亲节点
func (slf *BinTreeNode) GetParent() *BinTreeNode {
	if !slf.HasParent() {
		return nil
	}
	return slf.parent
}

//setParent 设置父亲节点
func (slf *BinTreeNode) setParent(p *BinTreeNode) {
	slf.parent = p
	// slf.parent.SetHeight() //更新父结点及其祖先高度
	// slf.parent.SetSize()   //更新父结点及其祖先规模
}

//CutOffParent 断开与父亲的关系
func (slf *BinTreeNode) CutOffParent() {
	if !slf.HasParent() {
		return
	}
	if slf.IsLChild() {
		slf.parent.lChild = nil //断开该节点与父节点的连接
	} else {
		slf.parent.rChild = nil //断开该节点与父节点的连接
	}

	slf.parent = nil       //断开该节点与父节点的连接
	slf.parent.SetHeight() //更新父结点及其祖先高度
	slf.parent.SetSize()   //更新父结点及其祖先规模
}

//HasLChild 判断是否有左孩子
func (slf *BinTreeNode) HasLChild() bool {
	return slf.lChild != nil
}

//GetLChild 获得左孩子节点
func (slf *BinTreeNode) GetLChild() *BinTreeNode {
	if !slf.HasLChild() {
		return nil
	}
	return slf.lChild
}

//SetLChild 设置当前结点的左孩子,返回原左孩子
func (slf *BinTreeNode) SetLChild(lc *BinTreeNode) *BinTreeNode {
	oldLC := slf.lChild
	if slf.HasLChild() {
		slf.lChild.CutOffParent() //断开当前左孩子与结点的关系
	}
	if lc != nil {
		lc.CutOffParent() //断开lc与其父结点的关系
		slf.lChild = lc   //确定父子关系
		lc.setParent(slf)
		slf.SetHeight() //更新当前结点及其祖先高度
		slf.SetSize()   //更新当前结点及其祖先规模
	}
	return oldLC
}

//HasRChild 判断是否有右孩子
func (slf *BinTreeNode) HasRChild() bool {
	return slf.rChild != nil
}

//GetRChild 获得右孩子节点
func (slf *BinTreeNode) GetRChild() *BinTreeNode {
	if !slf.HasRChild() {
		return nil
	}
	return slf.rChild
}

//SetRChild 设置当前结点的右孩子,返回原右孩子
func (slf *BinTreeNode) SetRChild(rc *BinTreeNode) *BinTreeNode {
	oldRC := slf.rChild
	if slf.HasRChild() {
		slf.rChild.CutOffParent() //断开当前左孩子与结点的关系
	}
	if rc != nil {
		rc.CutOffParent() //断开rc与其父结点的关系
		slf.rChild = rc   //确定父子关系
		rc.setParent(slf)
		slf.SetHeight() //更新当前结点及其祖先高度
		slf.SetSize()   //更新当前结点及其祖先规模
	}
	return oldRC
}

//IsLeaf 判断是否为叶子结点
func (slf *BinTreeNode) IsLeaf() bool {
	return !slf.HasLChild() && !slf.HasRChild()
}

//IsLChild 判断是否为某结点的左孩子
func (slf *BinTreeNode) IsLChild() bool {
	return slf.HasParent() && slf == slf.parent.lChild
}

//IsRChild 判断是否为某结点的右孩子
func (slf *BinTreeNode) IsRChild() bool {
	return slf.HasParent() && slf == slf.parent.rChild
}

//GetHeight 取结点的高度,即以该结点为根的树的高度
func (slf *BinTreeNode) GetHeight() int {
	return slf.height
}

//SetHeight 更新当前结点及其祖先的高度
func (slf *BinTreeNode) SetHeight() {
	newH := 0 //新高度初始化为0,高度等于左右子树高度加1中的大者
	if slf.HasLChild() {
		newH = int(math.Max(float64(newH), float64(1+slf.GetLChild().GetHeight())))
	}
	if slf.HasRChild() {
		newH = int(math.Max(float64(newH), float64(1+slf.GetRChild().GetHeight())))
	}
	if newH == slf.height {
		//高度没有发生变化则直接返回
		return
	}

	slf.height = newH //否则更新高度
	if slf.HasParent() {
		slf.GetParent().SetHeight() //递归更新祖先的高度
	}
}

//GetSize 取以该结点为根的树的结点数
func (slf *BinTreeNode) GetSize() int {
	return slf.size
}

//SetSize 更新当前结点及其祖先的子孙数
func (slf *BinTreeNode) SetSize() {
	slf.size = 1 //初始化为1,结点本身
	if slf.HasLChild() {
		slf.size += slf.GetLChild().GetSize() //加上左子树规模
	}
	if slf.HasRChild() {
		slf.size += slf.GetRChild().GetSize() //加上右子树规模
	}

	if slf.HasParent() {
		slf.parent.SetSize() //递归更新祖先的规模
	}

}
