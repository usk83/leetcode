# 791. Custom Sort String
# https://leetcode.com/problems/custom-sort-string/
def main():
  solution = Solution()
  print(solution.customSortString("cba", "abcd"))
  print(solution.customSortString("cbafg", "abcd"))

class Solution:
  def customSortString(self, S: str, T: str) -> str:
    dict = {}
    for s in S:
      dict[s] = 0

    sorted = ""
    for t in T:
      if t in dict:
        dict[t] += 1
      else:
        sorted += t

    for k, v in dict.items():
      sorted += k * v

    return sorted

if __name__ == '__main__':
    main()
