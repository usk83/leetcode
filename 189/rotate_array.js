// 189. Rotate Array
// https://leetcode.com/problems/rotate-array/
const main = () => {
  let nums;
  console.log("=========================================");
  nums = [1, 2, 3, 4, 5, 6, 7]
  rotate(nums, 3)
  console.log(nums);
  console.log([5, 6, 7, 1, 2, 3, 4]);
  console.log("=========================================");
  nums = [-1, -100, 3, 99]
  rotate(nums, 2)
  console.log(nums);
  console.log([3, 99, -1, -100]);
  console.log("=========================================");
  nums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]
  rotate(nums, 3)
  console.log(nums);
  console.log([10, 11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9]);
  console.log("=========================================");
  nums = [1, 2, 3, 4, 5, 6]
  rotate(nums, 4)
  console.log(nums);
  console.log([3, 4, 5, 6, 1, 2]);
  console.log("=========================================");
};

function rotate(nums, k) {
  k = k % nums.length;
  const times = nums.length - k;
  const orgLength = nums.length;
  for (let i = 0; i < times; i++) {
    nums[orgLength + i] = nums[i];
  }
  nums.splice(0, times);
}

main();
