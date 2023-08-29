# B-Tree

B-Tree handles massive amounts of data with ease. When it comes to storing and searching large amounts of data, traditional binary search trees can be impractical due to their poor performance and high memory usage.

B-Tree are characterized by the large number of keys that they can store in a single node, which is why they are also known as "large key" trees. Each node in a B-Tree can contain multiple keys, which allows the tree to have a larger branching factor and thus a shallower height. This shallow hight leads to less disk I/O, which results in faster search and insertion operations. B-Tree are particularly well suited for storage sytems that have slow.

B-Tree maintain balance by ensuring that each node has a minimum number of keys, so the tree is always balanced. This balance guarantees that the time complexity for operations such as insertion, deletion, and searching is always O(logn).

## Properties of B-Tree

- All leaves are at the same level
- B-Tree is defined by the term minimum degree 't'. The value of 't' depends upon disk block size.
- Every node except the roto must contain at least t-1 keys. The root may contain a minimum of 1 key.
- All nodes (including root) may contain at most (2*t-1) keys.
- Number of children of a node is equal to the number of keys in it plus 1.
- All keys of a node are sorted in increasing order. The child between two keys k1 and k2 contains all keys in the range from k1 and k2.
- B-Tree grows and shrinks from the root which is unlike Binary Search Tree. Binary Search Trees grow downward and also shrink from download
- Like other balanced Binary Search Trees, the time complexity to search, insert and delete is O(logn)
- Insertion of a Node in B-Tree happens only at Leaf Node

## Time Complexity of B-Tree

|Sr. No.|Algorithm|Time Complexity|
|-|-|-|
|1|Search|O(logn)|
|2|Insert|O(logn)|
|2|Delete|O(logn)|

## Interesting Facts about B-Trees

The minimum height of the B-Tree that exist with n number of nodes and m is the maximum number of children of a node can have is h(min) = [logm(n+1)] - 1

## Traversal in B-Tree

Traversal is also similar to Inorder traversal of Binary Tree. We start from the leftmost child, recursively print the leftmost child, then repeat the same process for the remaining children and keys. In the end, recursively print the rightmost child.

## Search Operation in B-Tree

- Start from the root and recursively traverse down.
- For every visited non-leaf node
    - If the node has the key, we simply return the node
    - Otherwise, we recur down to the appropriate child
- If we reach a leaf node and don't find k in the leaf node, then return NULL

Searching a B-Tree is similar to searching a binary tree. The algorithm is similar and goes with recursion. At each level, the search is optimized as if the key value is not present in the range of the parent then the key is present in another branch.