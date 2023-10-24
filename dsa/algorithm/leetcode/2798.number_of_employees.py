class Solution2798(object):
    def numberOfEmployeesWhoMetTarget(self, hours, target):
        """
        :type hours: List[int]
        :type target: int
        :rtype: int
        """
        rs = 0
        for hour in hours:
            if hour >= target:
                rs += 1
        return rs
