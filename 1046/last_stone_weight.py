# 1046. Last Stone Weight
# https://leetcode.com/problems/last-stone-weight/
from typing import List

def main():
  solution = Solution()
  print("=========================================")
  print(solution.lastStoneWeight([2, 7, 4, 1, 8, 1]))
  print(1)
  print("=========================================")
  # print(solution.)
  # print()
  # print("=========================================")

class Solution:
  def lastStoneWeight(self, stones: List[int]) -> int:
    while len(stones) > 1:
      largest, larger = 0, 0
      left = []
      for stone in stones:
        if stone > largest:
          largest, stone = stone, largest
        if stone > larger:
          larger, stone = stone, larger
        if stone != 0:
          left.append(stone)
      if largest != larger:
        left.append(largest-larger)
      stones = left
    if len(stones) == 0:
      return 0
    return stones[0]

if __name__ == '__main__':
    main()
