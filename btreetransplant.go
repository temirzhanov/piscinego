package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root == node {
		return rplc
	}

	if root.Left != nil {
		root.Left = BTreeTransplant(root.Left, node, rplc)
		root.Left.Parent = root
	}

	if root.Right != nil {
		root.Right = BTreeTransplant(root.Right, node, rplc)
		root.Right.Parent = root
	}

	return root
}
