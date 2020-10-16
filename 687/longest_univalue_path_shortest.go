// longestUnivaluePathV4 最適化バージョン
func longestUnivaluePathV4(root *TreeNode) int {
	return max(trace(root))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func trace(node *TreeNode) (int, int) {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return 0, 0
	}

	leftLength, currentLongest := traceNext(node.Left, node.Val, 0)
	rightLength, currentLongest := traceNext(node.Right, node.Val, currentLongest)

	if leftLength != 0 && rightLength != 0 {
		currentLongest = max(currentLongest, leftLength+rightLength)
	}

	return max(leftLength, rightLength), currentLongest
}

func traceNext(node *TreeNode, val, currentLongest int) (int, int) {
	if node == nil {
		return 0, currentLongest
	}
	if val != node.Val {
		return 0, max(currentLongest, max(trace(node)))
	}
	length, prevLongest := trace(node)
	return length + 1, max(currentLongest, prevLongest)
}
