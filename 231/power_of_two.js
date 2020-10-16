// 231. Power of Two
// https://leetcode.com/problems/power-of-two/
const main = () => {
  console.log("=========================================");
  console.log(isPowerOfTwo(1));
  console.log(true);
  console.log("=========================================");
  console.log(isPowerOfTwo(16));
  console.log(true);
  console.log("=========================================");
  console.log(isPowerOfTwo(218));
  console.log(false);
  console.log("=========================================");
  console.log(isPowerOfTwo(0));
  console.log(false);
  console.log("=========================================");
};

/**
 * @param {number} n
 * @return {boolean}
 */
const isPowerOfTwo = n => {
  return n > 0 && (n & n - 1) === 0
};

// const isPowerOfTwo = n => {
//   for (; n > 1 && (n & 1) != 1; n >>= 1) {}
//   return n === 1
// };

main();
