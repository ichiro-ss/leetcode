class Solution:
    def twoSum(self, nums: list, target: int) -> list:
        hashmap = {}
        for i in range(len(nums)):
            complement = target - nums[i]
            if complement in hashmap:
                return [hashmap[complement], i]
            hashmap[nums[i]] = i
            print(hashmap)

x = Solution()
print(x.twoSum([0, 3, 0], 0))

"""exp1. twoSum([2,7,11,15], 17)
{2: 0}
{2: 0, 7: 1}
{2: 0, 7: 1, 11: 2}
[0, 3]

exp2. twoSum([3, 3, 3], 6)
{3: 0}
[0, 1]

exp3. twoSum([0, 3, 0], 0)
{0: 0}
{0: 0, 3: 1}
[0, 2]
"""