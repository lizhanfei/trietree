package trietree

type TreeNode struct {
	Content rune
	IsEnd bool
	Child map[rune]*TreeNode
}
