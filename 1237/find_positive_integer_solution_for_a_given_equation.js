// 1237. Find Positive Integer Solution for a Given Equation
// https://leetcode.com/problems/find-positive-integer-solution-for-a-given-equation/
const main = () => {
  console.log("=========================================")
  console.log("No local tests available.")
  console.log("=========================================")
};

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
// 1 <= function_id <= 9
// 1 <= z <= 100
// It's guaranteed that the solutions of f(x, y) == z will be on the range 1 <= x, y <= 1000
// It's also guaranteed that f(x, y) will fit in 32 bit signed integer if 1 <= x, y <= 1000
// const findSolution = (customfunction, z) => {
//   const result = [];
//   for (let x = 1; x <= 1000; x++) {
//     if (customfunction.f(x, 1) > z) {
//       break;
//     }
//     let start = 1;
//     let end = 1000;
//     while (start <= end) {
//       const y = Math.floor((start + end) / 2);;
//       const ret = customfunction.f(x, y);
//       if (ret === z) {
//         result.push([x, y]);
//         break;
//       }
//       if (ret > z) {
//         end = y - 1;
//       } else {
//         start = y + 1;
//       }
//     }
//   }
//   return result;
// };


const findSolution = (customfunction, z) => {
  const result = [];
  (() => {
    for (let x = 1; x <= 1000; x++) {
      for (let y = 1; y <= 1000; y++) {
        const ret = customfunction.f(x, y);
        if (ret === z) {
          result.push([x, y]);
          break;
        } else if (ret > z) {
          if (y === 1) {
            return;
          }
          break;
        }
      }
    }
  })();
  return result;
};

const findSolution = (customfunction, z) => {
  const result = [];
OUTER_LOOP:
  for (let x = 1; x <= 1000; x++) {
    for (let y = 1; y <= 1000; y++) {
      const ret = customfunction.f(x, y);
      if (ret === z) {
        result.push([x, y]);
        break;
      } else if (ret > z) {
        if (y === 1) {
          break OUTER_LOOP;
        }
        break;
      }
    }
  }
  return result;
};

main();
