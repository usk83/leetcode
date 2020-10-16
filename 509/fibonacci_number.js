// 509. Fibonacci Number
// https://leetcode.com/problems/fibonacci-number/
const main = () => {
  console.log("=========================================");
  console.log(fib(10));
  console.log(55);
  console.log("=========================================");
  console.log(fib(15));
  console.log(610);
  console.log("=========================================");
  console.log(fib(22));
  console.log(17711);
  console.log("=========================================");
  console.log(fib(30));
  console.log(832040);
  console.log("=========================================");
};

/**
 * DP solution 2
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 48 ms, faster than 91.78%
 * > Memory Usage: 33.8 MB, less than 81.82%
 */
const fib = (n) => {
  const fibs = [0, 1, 1];
  for(let i = 3; i <= n; i++) {
    fibs[i % 3] = fibs[(i - 1) % 3] + fibs[(i - 2) % 3]
  }
  return fibs[n % 3];
};

/**
 * DP solution
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 */
// const fib = (n) => {
//   const fibs = new Array(n + 1);
//   [fibs[0], fibs[1]] = [0, 1];
//   for(let i = 2; i <= n; i++) {
//     fibs[i] = fibs[i - 1] + fibs[i - 2]
//   }
//   return fibs[n];
// };

/**
 * @musashi.sakamoto 's solution'
 * - Time Complexity: O(2^N)
 * - Space Complexity: O(N)
 */
// var fib = function(N) {
//   if (N === 0) return 0;
//   if (N === 1) return 1;
//   return fib(N - 1) + fib(N - 2);
// };

main();
