# Binary search

Binary search is a basic algorithm in computer science.

The basic idea of binary search algorithm is continuous divide search spacing into two partitions, and then remove invalid partition. 

```c
int binary_search(int A[], int sizeA, int target) {
    int lo = 1, hi = sizeA;
    while (lo <= hi) {
        int mid = (lo + hi) / 2;
        if (A[mid] == target)
            return mid;
        else if (A[mid] < target):
            lo = mid + 1;
        else 
            hi = mid - 1;
    }

    return -1
}
```

**Complexity**

In each step, search space reduce 1/2, so time complexity is O(logn), space complexity O(1) 

**Generality**

We need generality binary search for broader class of problems. Binary search can apply for any monotonic function receive input is an integer. 

Main theorem of binary search

In search space S include candidates for result of problems. The define a search function P: S -> true, false is a function receive a candidate x ∈ S and return true/false give x is valid or not. P is a function checking for a condition.

∀x, y ∈ S, y > S ∧ P(x) = true => P(y) = true

