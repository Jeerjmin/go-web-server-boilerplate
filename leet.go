package leetcode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {

	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}

	val := this.serialize(root.Left)
	if val != "null" {
		return val
	}

	val = serialize(root.Right)
	if val != "null" {
		return val
	}
}

func (this *Codec) deserialize(data string) *TreeNode {

	fmt.Println("des", data)

	return &TreeNode{}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
