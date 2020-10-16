// 1389. Create Target Array in the Given Order
// https://leetcode.com/problems/create-target-array-in-the-given-order/

/**
 * https://leetcode.com/problems/create-target-array-in-the-given-order/discuss/547440/JavaScript-Update-indices-recursively-without-splice
 */

const main = () => {
  console.log("=========================================");
  console.log(createTargetArray([0, 1, 2, 3, 4], [0, 1, 2, 2, 1]));
  console.log([0, 4, 1, 3, 2]);
  console.log("=========================================");
  console.log(createTargetArray([1, 2, 3, 4, 0], [0, 1, 2, 3, 0]));
  console.log([0, 1, 2, 3, 4]);
  console.log("=========================================");
  console.log(createTargetArray([1], [0]));
  console.log([1]);
  console.log("=========================================");
  console.log(createTargetArray([], []));
  console.log([]);
  console.log("=========================================");
  console.log(createTargetArray([1, 2, 3, 4, 5], [0, 1, 2, 3, 4]));
  console.log([1, 2, 3, 4, 5]);
  console.log("=========================================");
  console.log(createTargetArray([4, 2, 4, 3, 2], [0, 0, 1, 3, 1]));
  console.log([2, 2, 4, 4, 3]);
  console.log("=========================================");
};

/**
 * @param {number[]} nums
 * @param {number[]} index
 * @return {number[]}
 */
function createTargetArray(nums, index) {
  const indices = Array(index.length).fill(null);
  for (const [from, to] of index.entries()) {
    // console.debug('=== DEBUG START ======================================');
    // console.debug("from:", from, "to:", to);
    set(from, to);
    // console.debug(indices);
    // console.debug('=== DEBUG END ========================================');
  }
  return indices.map(index => nums[index]);

  function set(from, to) {
    if (from === null) {
      return;
    }
    set(indices[to], to + 1);
    indices[to] = from;
  }
}

main();
