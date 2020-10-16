// longestUnivaluePathV2 最長の長さの更新を関数にまとめたバージョン
func longestUnivaluePathV2(root *TreeNode) int {
	var longestLength int
	setIfLongest := func(currentLength int) {
		if currentLength > longestLength {
			longestLength = currentLength
		}
	}

	var trace func(node *TreeNode) int
	trace = func(node *TreeNode) int {
		if node == nil || (node.Left == nil && node.Right == nil) {
			return 0
		}

		leftLength := trace(node.Left)
		rightLength := trace(node.Right)

		if node.Left != nil && node.Val == node.Left.Val {
			leftLength++
		} else {
			setIfLongest(leftLength)
			leftLength = 0
		}

		if node.Right != nil && node.Val == node.Right.Val {
			rightLength++
		} else {
			setIfLongest(rightLength)
			rightLength = 0
		}

		if leftLength != 0 && rightLength != 0 {
			setIfLongest(leftLength + rightLength)
		}

		if leftLength > rightLength {
			return leftLength
		}
		return rightLength
	}

	setIfLongest(trace(root))

	return longestLength
}
