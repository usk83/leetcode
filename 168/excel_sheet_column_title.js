// 168. Excel Sheet Column Title
// https://leetcode.com/problems/excel-sheet-column-title/
const main = () => {
  console.log("=========================================");
  console.log("%s - A", convertToTitle(1));
  console.log("%s - Z", convertToTitle(26));
  console.log("%s - AA", convertToTitle(27));
  console.log("%s - AZ", convertToTitle(52));
  console.log("%s - BA", convertToTitle(53));
  console.log("%s - BZ", convertToTitle(78));
  console.log("%s - AAA", convertToTitle(703));
  console.log("=========================================");
  console.log("%s - A", convertToTitle(1));
  console.log("%s - AA", convertToTitle(27));
  console.log("%s - AAA", convertToTitle(703));
  console.log("%s - AAAA", convertToTitle(18279));
  console.log("%s - AAAAA", convertToTitle(475255));
  console.log("%s - BAAAA", convertToTitle(932231));
  console.log("%s - ZAAAA", convertToTitle(11899655));
  console.log("%s - AAAAAA", convertToTitle(12356631));
  console.log("=========================================");
};

const convertToTitle = n => {
  return n === 0 ? '' : convertToTitle(Math.floor((n-1)/26)) + String.fromCharCode(65+(n-1)%26);
};

main();
