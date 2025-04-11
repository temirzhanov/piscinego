package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := BTreeLevelCount(root.Left) + 1
	right := BTreeLevelCount(root.Right) + 1

	if left > right {
		return left
	}
	return right
}
