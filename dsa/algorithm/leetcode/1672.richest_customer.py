class Solution1672(object):
    def maximumWealth(self, accounts):
        """
        :type accounts: List[List[int]]
        :rtype: int
        """
        return max([sum(acc) for acc in accounts])
