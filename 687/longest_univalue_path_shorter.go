// longestUnivaluePathV3 共通処理を関数にまとめて極力短くしたバージョン
func longestUnivaluePathV3(root *TreeNode) int {
	var longestLength int
	setIfLongest := func(currentLength int) {
		if currentLength > longestLength {
			longestLength = currentLength
		}
	}

	var trace func(node *TreeNode) int
	var checkLength func(cur, next *TreeNode) int
	checkLength = func(cur, next *TreeNode) int {
		if cur == nil || next == nil {
			return 0
		}
		currentLength := trace(next)
		if cur.Val != next.Val {
			setIfLongest(currentLength)
			return 0
		}
		return currentLength + 1
	}
	trace = func(node *TreeNode) int {
		if node == nil || (node.Left == nil && node.Right == nil) {
			return 0
		}

		leftLength := checkLength(node, node.Left)
		rightLength := checkLength(node, node.Right)

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
