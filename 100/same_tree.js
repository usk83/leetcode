// 100. Same Tree
// https://leetcode.com/problems/same-tree/
const main = () => {};

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {boolean}
 */
const isSameTree = (p, q) => !p || !q ? p === q : p.val === q.val && isSameTree(p.left, q.left) && isSameTree(p.right, q.right);

main();
