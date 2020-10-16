// 506. Relative Ranks
// https://leetcode.com/problems/relative-ranks/
import java.util.*;

class Solution {
  public static void main(String[] args) {
    Solution solution = new Solution();
    System.out.println("=========================================");
    System.out.println(Arrays.toString(solution.findRelativeRanks(new int[]{10, 3, 8, 9, 4})));
    System.out.println(Arrays.toString(new String[]{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}));
    System.out.println("=========================================");
  }

  public String[] findRelativeRanks(int[] nums) {
    Integer[] index = new Integer[nums.length];

    for (int i = 0; i < nums.length; i++) {
      index[i] = i;
    }

    Arrays.sort(index, (a, b) -> (nums[b] - nums[a]));

    String[] result = new String[nums.length];

    for (int i = 0; i < nums.length; i++) {
      if (i == 0) {
        result[index[i]] = "Gold Medal";
      }
      else if (i == 1) {
        result[index[i]] = "Silver Medal";
      }
      else if (i == 2) {
        result[index[i]] = "Bronze Medal";
      }
      else {
        result[index[i]] = (i + 1) + "";
      }
    }

    return result;
  }
}
