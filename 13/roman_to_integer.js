// 13. Roman to Integer
// https://leetcode.com/problems/roman-to-integer/
const main = () => {
  console.log("=========================================");
  console.log(romanToInt("III"));
  console.log(3);
  console.log("=========================================");
  console.log(romanToInt("IV"));
  console.log(4);
  console.log("=========================================");
  console.log(romanToInt("IX"));
  console.log(9);
  console.log("=========================================");
  console.log(romanToInt("LVIII"));
  console.log(58);
  console.log("=========================================");
  console.log(romanToInt("MCMXCIV"));
  console.log(1994);
  console.log("=========================================");
};

const romanNumerals = { 'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000 };
const romanToInt = s => {
  let sum = val = 0;
  for (let i = s.length - 1; i >= 0; i--) {
    sum += val <= (val = romanNumerals[s[i]]) ? val : -val;
  }
  return sum;
};

main();
