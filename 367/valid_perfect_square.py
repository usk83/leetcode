# 367. Valid Perfect Square
# https://leetcode.com/problems/valid-perfect-square/
def main():
  solution = Solution()
  print("=========================================")
  print(solution.isPerfectSquare(1))
  print(True)
  print("=========================================")
  print(solution.isPerfectSquare(4))
  print(True)
  print("=========================================")
  print(solution.isPerfectSquare(9))
  print(True)
  print("=========================================")
  print(solution.isPerfectSquare(16))
  print(True)
  print("=========================================")
  print(solution.isPerfectSquare(14))
  print(False)
  print("=========================================")
  print(solution.isPerfectSquare(2147483647))
  print(False)
  print("=========================================")
  print(solution.isPerfectSquare((1<<63) - 1))
  print(False)
  print("=========================================")
  print(solution.isPerfectSquare(3037000499 ** 2))
  print(True)
  print("=========================================")

class Solution:
  def isPerfectSquare(self, num: int) -> bool:
    start, end = 1, num
    while (start <= end):
        mid = start + (end - start >> 1)
        div = num // mid
        if div == mid and num % mid == 0:
            return True
        elif div > mid:
            start = mid + 1
        else:
            end = mid - 1
    return False

  # def isPerfectSquare(self, num: int) -> bool:
  #   start, end = 1, num
  #   while (start <= end):
  #       mid = (start + end) // 2
  #       square = mid * mid
  #       if square == num:
  #           return True
  #       elif square < num:
  #           start = mid + 1
  #       else:
  #           end = mid - 1
  #   return False

if __name__ == '__main__':
    main()
