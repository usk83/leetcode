// 367. Valid Perfect Square
// https://leetcode.com/problems/valid-perfect-square/
fun main(args: Array<String>) {
    val solution = Solution()
    println("=========================================")
    println(solution.isPerfectSquare(1))
    println(true)
    println("=========================================")
    println(solution.isPerfectSquare(4))
    println(true)
    println("=========================================")
    println(solution.isPerfectSquare(9))
    println(true)
    println("=========================================")
    println(solution.isPerfectSquare(16))
    println(true)
    println("=========================================")
    println(solution.isPerfectSquare(14))
    println(false)
    println("=========================================")
    println(solution.isPerfectSquare(2147483647))
    println(false)
    println("=========================================")
    // println(solution.isPerfectSquare((1 shl 63) - 1))
    // println(false)
    // println("=========================================")
    // println(solution.isPerfectSquare(3037000499 * 3037000499))
    // println(true)
    // println("=========================================")
}

class Solution {
    fun isPerfectSquare(num: Int): Boolean {
        var min = 1
        var max = num
        while (min <= max) {
            val mid = min + (max - min) / 2
            val div = num / mid
            if (div == mid && num % mid == 0) return true
            if (div > mid) {
                min = mid + 1
            } else {
                max = mid - 1
            }
        }
        return false
    }
}
