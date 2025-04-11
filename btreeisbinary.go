package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := true
	if root.Left != nil {
		if root.Left.Data >= root.Data {
			left = false
		} else {
			left = BTreeIsBinary(root.Left)
		}
	}

	right := true

	if root.Right != nil {
		if root.Right.Data < root.Data {
			right = false
		} else {
			right = BTreeIsBinary(root.Right)
		}
	}

	return right && left
}
