// 1365. How Many Numbers Are Smaller Than the Current Number
// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
const main = () => {
  console.log("=========================================");
  console.log(smallerNumbersThanCurrent([8,1,2,2,3]));
  console.log([4,0,1,1,3]);
  console.log("=========================================");
  console.log(smallerNumbersThanCurrent([6,5,4,8]));
  console.log([2,1,0,3]);
  console.log("=========================================");
  console.log(smallerNumbersThanCurrent([7,7,7,7]));
  console.log([0,0,0,0]);
  console.log("=========================================");
  console.log(smallerNumbersThanCurrent([5, 0, 10, 0, 10, 6]));
  console.log([2, 0, 4, 0, 4, 3]);
  console.log("=========================================");
};

// ref: https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/discuss/525004/JavaScript-Clean-solution-using-array-sort
function smallerNumbersThanCurrent(nums) {
  const sorted = Array.from(nums).sort((n1, n2) => n2 - n1);
  console.debug('=== DEBUG START ======================================');
  console.debug(sorted);
  console.debug(sorted.map((num, index) => [num, nums.length - index - 1]));
  const map = new Map(sorted.map((num, index) => [num, nums.length - index - 1]));
  console.debug('=== DEBUG END ========================================');
  return nums.map(num => map.get(num));
}

main();
