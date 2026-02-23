# 76. Minimum Window Substring

**Link:** [LeetCode 76](https://leetcode.com/problems/minimum-window-substring/)
**Difficulty:** 🔴 Hard
**Pattern:** Sliding Window (Variable Size) + Frequency Map

## 💡 The "Aha!" Moment (The Ultimate Hack)
Instead of looping through a 256-size array every single time to check if our window is valid, **track a `matched` counter.**
1. Count the unique characters needed from string `t`.
2. Expand the `right` pointer. If we find a character we need, and its frequency in our window matches the target frequency, increment `matched`.
3. **The Trigger:** When `matched == unique_chars_needed`, our window is perfectly valid!
4. Shrink the `left` pointer to find the minimum size, updating `minLen`, and decrement `matched` if we throw away a crucial character.

## 🧠 Step-by-Step Logic
1. **Setup:** Create a `targetMap` (or array of size 256) to store the frequencies of characters in `t`.
2. **Expand (Right):** Loop `right` from `0` to `len(s)`. Add `s[right]` to our `windowMap`.
    - *Check:* Does `windowMap[s[right]] == targetMap[s[right]]`? If yes, `matched++`.
3. **Shrink (Left):** `while (matched == required_unique_chars)`:
    - *Record:* Is `right - left + 1 < minLen`? Update `minLen` and `minStart`.
    - *Discard:* Remove `s[left]` from `windowMap`.
    - *Check:* Did we just break our valid window? If `windowMap[s[left]] < targetMap[s[left]]`, then `matched--`.
    - *Move:* `left++`.

## ⏱️ Time & Space Complexity
- **Time:** $O(N + M)$ - Where $N$ is the length of `s` and $M$ is the length of `t`. Both `left` and `right` pointers only move forward, meaning we visit each character at most twice. (The $O(1)$ 256-loop is completely eliminated).
- **Space:** $O(1)$ - The maps/arrays will never exceed 256 characters (ASCII limit), which is a constant space regardless of input size.

## ⚠️ Edge Cases & Gotchas
- **Strings are empty:** If `s` or `t` are empty, or `len(s) < len(t)`, return `""` immediately.
- **Initialization Trap:** Initialize `minLen` to `Integer.MAX_VALUE` (or `len(s) + 1`), **NOT** `0`, otherwise it will never update.
- **Off-by-one:** When slicing the final substring, remember to use `[minStart : minStart + minLen]`.

## 💻 Core Code Skeleton (Mental Revision)
```go
// 1. Populate target map and count required UNIQUE chars
required := len(targetMap) 
matched := 0

for right < len(s) {
    // 2. Add to window
    charRight := s[right]
    windowMap[charRight]++
    if windowMap[charRight] == targetMap[charRight] {
        matched++
    }

    // 3. Shrink if valid
    for matched == required {
        // Update minLen here...
        
        charLeft := s[left]
        windowMap[charLeft]--
        if windowMap[charLeft] < targetMap[charLeft] {
            matched-- // Window is no longer valid!
        }
        left++
    }
    right++
}