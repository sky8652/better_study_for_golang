/**
实现链表的逆序
方法二：递归法
*/
package main

import "fmt"

//链表定义
type LNode struct {
	Data interface{} // 数据域
	Next *LNode      //  对下一个节点的作用
}

//创建链表
func CreateNode(node *LNode, max int) {
	cur := node // 不用这个变量试试下面打印  这里有个go函数调用与栈的关系
	for i := 1; i < max; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}

func main() {
	head := &LNode{}
	CreateNode(head, 8)
	PrintNode("逆序前", head)
	Reverse(head)
	PrintNode("逆序后", head)
}

//打印链表的方法
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

func Reverse(node *LNode) {
	firstNode := node.Next
	//递归调用
	newHead := ReverseChild(firstNode)
	node.Next = newHead
}

func ReverseChild(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := ReverseChild(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}
