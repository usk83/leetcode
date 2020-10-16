// 139. Word Break
// https://leetcode.com/problems/word-break/
const main = () => {
  console.log("=========================================");
  console.log(wordBreak("leetcode", ["leet", "code"]));
  console.log(true);
  console.log("=========================================");
  console.log(wordBreak("applepenapple", ["apple", "pen"]));
  console.log(true);
  console.log("=========================================");
  console.log(wordBreak("catsandog", ["cats", "dog", "sand", "and", "cat"]));
  console.log(false);
  console.log("=========================================");
  console.log(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", ["a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"]));
  console.log(false);
  console.log("=========================================");
};

const wordBreak = (s, wordDict) => {
  const dp = new Array(s.length + 1).fill(false);
  dp[0] = true;
  for (let i = 0; i <= s.length; i++) {
    if (!dp[i]) continue; // その文字まで到達できてない場合はその先の可能性はない
    for (const word of wordDict) {
      if (i + word.length > s.length || dp[i + word.length]) { // wordを検証する必要があるか確認
        continue
      }
      if (word === s.slice(i, i + word.length)) { // これまで到達した文字列に今回のwordを繋げられる場合
        dp[i + word.length] = true               // つなげた先のindexを到達可能にする
                                                 // ここで末尾まで到達したときに即座に結果を返すことも可能
      }
    }
  }
  return dp[s.length];
};

/*
 * Junya's code
 */
// const wordBreak = (s, wordDict) => {
//   const arr = new Array(s.length + 1).fill(false);
//   arr[0] = true;
//   for (let i = 0; i < s.length + 1; i++) {
//     for (let j = 0; j < i; j++) {
//       if (!arr[j]) continue;
//       if (wordDict.includes(s.slice(j, i))) {
//         arr[i] = true;
//         break;
//       }
//     }
//   }
//   return arr[s.length];
// };

main();

/*
 * Review
 */
// 解説聞いたあとに実際に実装してみるの、めっちゃ大事なのでいいと思います！
//
//
// せっかくなのでコードみてみて思ったことを
// 以下は個人的な意見なので参考までに
//
// ロジックが複雑になってて分かりにくく感じました
// どういう操作をしているか完全に理解できていればいいですが、そうでなかったら少し見直せるかも
//
// 解読するの苦戦したので合っているか誤りがあるかもしれませんが、やっていることは
//
// 1. 先頭から1文字ずつ各文字まで到達可能であるかを調べていく O(N)
// 2. 特定の文字番目までの部分文字列を切り出す O(N)
// 3. 到達可能性がある部分文字列について、辞書から検索する O(M)
//
// こうなっているかと思います
//
// 実際には早期に弾かれるケースがあるので、ワーストケースが O(N^2*M) と言っていいかは怪しいですが、
// なんでロジックが分かりにくいのかなと考えた結果
//
// 操作の元と対象が人間の思考と逆になっている
//
// と感じました
//
// 1. 先頭から1文字ずつ各文字まで到達可能であるかを調べていく O(N)
// というロジックを変えずに次の処理について考えると、操作の元と対象は
//
// 部分文字列→辞書（部分文字列が辞書に存在するかを調べる）
//
// のではなく
//
// 辞書→部分文字列（辞書の文字が部分文字列と一致するか）
//
// と考えるほうが自然にかなと
// その場合、コード全体の構成は
//
// 1. 先頭から1文字ずつ各文字まで到達可能であるかを調べていく O(N)
// 2. 辞書内の単語1つ1つについて、その単語が末尾にマッチした場合の到達可能性があるのであれば照合する O(M)
//
// となるかなと思います
//
// 基本的にこの問題、人間的に考えると、
//
// 全体文字列の部分が辞書の単語で構成されるか
// というより
// 辞書の単語で全体文字列を構成できるか
//
// と考えたほうが自然かなと思うんですよね
//
// ちなみにそう考えると、
// ①各文字ごとにその文字まで到達可能であるか
// と考えるより
// ②到達可能である文字からどの文字まで到達することができるか
// と考えたほうが素直かと思います
//
// ということで②のバージョンのコードをjsで書いてみました
//
// ```
// const wordBreak = (s, wordDict) => {
//   const dp = new Array(s.length + 1).fill(false);
//   dp[0] = true;
//   for (let i = 0; i <= s.length; i++) {
//     if (!dp[i]) continue; // その文字まで到達できてない場合はその先の可能性はない
//     for (const word of wordDict) {
//       if (i + word.length > s.length || dp[i + word.length]) { // wordを検証する必要があるか確認
//         continue
//       }
//       if (word === s.slice(i, i + word.length)) { // これまで到達した文字列に今回のwordを繋げられる場合
//         dp[i + word.length] = true               // つなげた先のindexを到達可能にする
//                                                  // ここで末尾まで到達したときに即座に結果を返すことも可能
//       }
//     }
//   }
//   return dp[s.length];
// };
// ```
//
// 元のコードより長くなってしまったので、ちょっとどっちのほうがいいか不安になってきましたが参考までに
//
// あと、indexを使ってのforループは可読性の面では不利なので、index自体を意図して必要とする場合以外は`for-of`とかの適切なイテレーション方法を取れるといいですね
//
// 参考: https://frog-events.slack.com/archives/CF1Q3KZAT/p1581487951049300?thread_ts=1581433249.047600&cid=CF1Q3KZAT
