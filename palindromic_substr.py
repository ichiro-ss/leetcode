class Solution:
    def longestPalindrome(self, s: str) -> str:
        n = len(s)
        P = [[False for _ in range(n)] for _ in range(n)]
        max_str = ""
        for i in range(n-1, -1, -1):
            for j in range(n):
                if i == j:
                    P[i][j] = True
                    max_str = s[i]
        for i in range(n-2, -1, -1):
            if s[i] == s[i+1]:
                P[i][i+1] = True
                max_str = s[i:i+2]

        for width in range(2, n):
            for i in range(n - width -1, -1, -1):
                if s[i] == s[i+width] and \
                    0 <= i+1 < n and 0 <= i+width-1 < n and \
                        P[i+1][i+width-1]:
                    P[i][i+width] = True
                    # print("!", end="")
                    max_str = s[i:i+width+1]
                # print(s[i], ",", s[i+width])

        return max_str

substr = str(input())
print(Solution().longestPalindrome(substr))
