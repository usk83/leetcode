/**
 * // This is the CustomFunction's API interface.
 * // You should not implement it, or speculate about its implementation
 * function CustomFunction() {
 *
 *     @param {integer, integer} x, y
 *     @return {integer}
 *     this.f = function(x, y) {
 *         ...
 *     };
 *
 * };
 */
/**
 * @param {CustomFunction} customfunction
 * @param {integer} z
 * @return {integer[][]}
 */
var findSolution = function(customfunction, z) {
  let x = 1;
  let y = 1;
  let result = [];
  for (let x = 1; x < 1000; x++) {
    let ret = customfunction.f(x, 1)
    if (ret > z) {
      break;
    }
    for (let y = 1; y < 1000; y++) {
      let ret = customfunction.f(x, y)
      if (ret == z) {
        result.push([x, y]);
        break;
      }
      if (ret > z) {
        break;
      }
    }
  }
  return result;
};

// 1 <= function_id <= 9
// 1 <= z <= 100
// It's guaranteed that the solutions of f(x, y) == z will be on the range 1 <= x, y <= 1000
// It's also guaranteed that f(x, y) will fit in 32 bit signed integer if 1 <= x, y <= 1000

console.debug('=========================================');
console.log(findSolution(1, 5));
console.log([[1,4],[2,3],[3,2],[4,1]]);
console.debug('=========================================');
console.log(findSolution(2, 5));
console.log([[1,5],[5,1]]);
console.debug('=========================================');
