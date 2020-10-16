// 404. Sum of Left Leaves
// https://leetcode.com/problems/sum-of-left-leaves/
public class TreeNode {
  public var val: Int
  public var left: TreeNode?
  public var right: TreeNode?
  public init(_ val: Int) {
    self.val = val
    self.left = nil
    self.right = nil
  }
}

class Solution {
  func sumOfLeftLeaves(_ root: TreeNode?) -> Int {
    var dfs: ((TreeNode?, Bool) -> Int)!
    dfs = {(_ node: TreeNode?, _ isLeft: Bool) in
      guard let node = node else { return 0 }
      if node.left == nil && node.right == nil && isLeft {
        return node.val
      }
      return dfs(node.left, true) + dfs(node.right, false)
    }
    return dfs(root, false)
  }
}
