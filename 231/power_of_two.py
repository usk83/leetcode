# 231. Power of Two
# https://leetcode.com/problems/power-of-two/
def main():
  solution = Solution()
  print("=========================================")
  print(solution.isPowerOfTwo(1))
  print(True)
  print("=========================================")
  print(solution.isPowerOfTwo(16))
  print(True)
  print("=========================================")
  print(solution.isPowerOfTwo(218))
  print(False)
  print("=========================================")
  print(solution.isPowerOfTwo(0))
  print(False)
  print("=========================================")

class Solution:
  def isPowerOfTwo(self, n: int) -> bool:
    return n > 0 and n & n - 1 == 0

  # def isPowerOfTwo(self, n: int) -> bool:
  #   while (n > 1 and n&1 != 1): n >>= 1
  #   return n == 1

if __name__ == '__main__':
    main()
