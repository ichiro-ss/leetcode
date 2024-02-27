class Solution:
    def convert(self, s: str, numRows: int) -> str:
        zigzag = [[None] * len(s) for _ in range(numRows)]
        to_bottom = True
        i, j = 0, 0
        for _ in range(len(s)):
            zigzag[i][j] = s[_]
            if to_bottom and i == numRows - 1:
                to_bottom = False
            if not to_bottom and i == 0:
                to_bottom = True
            if numRows > 1:
                if to_bottom:
                    i += 1
                else:
                    i -= 1
                    j += 1
            else:
                j += 1

        res = ""
        for i in range(numRows):
            for j in range(len(s)):
                if zigzag[i][j] is not None:
                    res += zigzag[i][j]

        return res

s = "PAYPALISHIRING"
numRows = 4
print(Solution().convert(s, numRows))
