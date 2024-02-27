class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        i = 0
        length = 0; max_length = 0
        max_position = {}
        for j in range(len(s)):
            if s[j] in max_position:
                i = max(i, max_position[s[j]])
            max_position[s[j]] = j
            max_length = max(max_length, j - i + 1)
        return max_length

x = Solution()
print(x.lengthOfLongestSubstring("abcabcbb"))