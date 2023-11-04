class Solution:
    def isAcronym(self, words, s):
        if s.__len__() != words.__len__():
            return False
        for i in range(words.__len__()):
            if words[i][0] != s[i]:
                return False
        return True

s = Solution()
print(s.isAcronym(["never","gonna","give","up","on","you"], "ngguoy"))
