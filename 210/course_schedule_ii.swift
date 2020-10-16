// 210. Course Schedule II
// https://leetcode.com/problems/course-schedule-ii/
// import Foundation

func main() {
  let solution = Solution()
  print("=========================================")
  print(solution.findOrder(2, [[1, 0]]))
  print([0, 1])
  print("=========================================")
  print(solution.findOrder(4, [[1, 0], [2, 0], [3, 1], [3, 2]]))
  print("\([0, 1, 2, 3]) or \([0, 2, 1, 3])")
  print("=========================================")
  print(solution.findOrder(2, [[1, 0],[0, 1]]))
  print([])
  print("=========================================")
}

class Solution {
  func findOrder(_ n: Int, _ edges: [[Int]]) -> [Int] {
    var inDegrees = [Int](repeating: 0, count: n)
    var graph = [[Int]](repeating: [Int](), count: n)
    for edge in edges {
      inDegrees[edge[0]] += 1
      graph[edge[1]].append(edge[0])
    }

    var initialVs = [Int]()
    for v in 0..<n {
      if inDegrees[v] == 0 {
        initialVs.append(v)
      }
    }

    var orderedCourse = [Int]()
    orderedCourse.reserveCapacity(n)

    var dfs: ((Int) -> Void)!
    dfs = {(v) in
      orderedCourse.append(v)
      for u in graph[v] {
        inDegrees[u] -= 1
        if inDegrees[u] == 0 {
          dfs(u)
        }
      }
    }

    for v in initialVs {
      dfs(v)
    }

    return orderedCourse.count == n ? orderedCourse : [];
  }
}

main()
