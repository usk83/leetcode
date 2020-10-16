// 231. Power of Two
// https://leetcode.com/problems/power-of-two/
class Solution {
    fun isPowerOfTwo(n: Int): Boolean {
        return n > 0 && n and (n - 1) == 0
    }

    // fun isPowerOfTwo(n: Int): Boolean {
    //     var num = n
    //     while (num > 1 && (num and 1) != 1) { num = num.shr(1) }
    //     return num == 1
    // }
}

fun main(args: Array<String>) {
    val solution = Solution()
    println("=========================================")
    println(solution.isPowerOfTwo(1))
    println(true)
    println("=========================================")
    println(solution.isPowerOfTwo(16))
    println(true)
    println("=========================================")
    println(solution.isPowerOfTwo(218))
    println(false)
    println("=========================================")
    println(solution.isPowerOfTwo(0))
    println(false)
    println("=========================================")
}
