// 207. Course Schedule
// https://leetcode.com/problems/course-schedule/
// import Foundation

func main() {
  let solution = Solution()
  print("=========================================")
  print(solution.canFinish(2, [[1, 0]]))
  print(true)
  print("=========================================")
  print(solution.canFinish(2, [[1, 0],[0, 1]]))
  print(false)
  print("=========================================")
}

class Solution {
  func canFinish(_ n: Int, _ edges: [[Int]]) -> Bool {
    return canFinishDFS(n, edges)
    // return canFinishBFS(n, edges)
  }

  func canFinishDFS(_ n: Int, _ edges: [[Int]]) -> Bool {
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

    var dfs: ((Int) -> Void)!
    dfs = {(v) in
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

    for inDegree in inDegrees {
      if inDegree != 0 {
        return false
      }
    }

    return true
  }

  // ［WIP デバッグ中］なにかおかしい
  // 8
  // [[1,0],[2,6],[1,7],[6,4],[7,0],[0,5]]
  func canFinishBFS(_ n: Int, _ edges: [[Int]]) -> Bool {
    var inDegrees = [Int](repeating: 0, count: n)
    var graph = [[Int]](repeating: [Int](), count: n)
    for edge in edges {
      inDegrees[edge[1]] += 1
      graph[edge[0]].append(edge[1])
    }

    var q = Queue<Int>()
    for v in 0..<n {
      if inDegrees[v] == 0 {
        q.enqueue(v)
      }
    }

    if q.isEmpty {
      return false
    }

    while !q.isEmpty {
      let v = q.dequeue()!
      for u in graph[v] {
        inDegrees[u] -= 1
        if inDegrees[u] < 0 {
          return false
        }
        q.enqueue(u)
      }
    }

    return true
  }
}

public struct Queue<T> {
  fileprivate var array = [T]()

  public var isEmpty: Bool {
    return array.isEmpty
  }

  public var count: Int {
    return array.count
  }

  public mutating func enqueue(_ element: T) {
    array.append(element)
  }

  public mutating func dequeue() -> T? {
    if isEmpty {
      return nil
    } else {
      return array.removeFirst()
    }
  }

  public var front: T? {
    return array.first
  }
}

main()
