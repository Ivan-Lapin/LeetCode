class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        rev = list(str(x))
        rev.reverse()
        if int(''.join(rev)) == x:
            return True
        else:
            return False