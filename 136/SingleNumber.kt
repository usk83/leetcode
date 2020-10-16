// 136. Single Number
// https://leetcode.com/problems/single-number/
fun main(args: Array<String>) {
    val solution = Solution()
    println("=========================================")
    println(solution.singleNumber(intArrayOf(2, 2, 1)))
    println(1)
    println("=========================================")
    println(solution.singleNumber(intArrayOf(4, 1, 2, 1, 2)))
    println(4)
    println("=========================================")
}

class Solution {
    fun singleNumber(nums: IntArray): Int {
        nums.sort()
        for (i in 0 until nums.size-1 step 2) {
            if (nums.get(i) != nums.get(i+1)) {
                return nums.get(i)
            }
        }
        return nums.get(nums.size-1)
    }
}
