# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        dummyHead = ListNode()
        p = l1; q = l2; curr = dummyHead
        carry = 0
        while(p is not None or q is not None):
            x = p.val if p is not None else 0
            y = q.val if q is not None else 0
            sum = carry + x + y
            carry = sum // 10
            curr.next = ListNode(sum % 10)
            curr = curr.next
            if p is not None:
                p = p.next
            if q is not None:
                q = q.next
        if carry > 0:
            curr.next = ListNode(carry)
        return dummyHead.next

"""exp.
[2,4,3]
[5,6,4]
[0]
[0]
[9,9,9,9,9,9,9]
[9,9,9,9]
Output:
[7,0,8]
[0]
[8,9,9,9,0,0,0,1]
"""