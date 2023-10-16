# Database Normalization

- What problems arise due to Data redundancy

- How Normalization makes the Data more meaningful and useable

**Function Dependency**

```txt
FD X -> Y: is trivial if y ⊂ X 
X = {A, B, C, D}, Y = {C, D}
Y ⊂ X => X -> Y is trivial
```

```txt
FD X -> Y: is non-trivial if y ⊄ X 
X = {A, B, C, D}, Y = {E, F} 
```

```txt
FD X -> Y: is multivalent if Yi -> Yj (Yi, Yj ∈ Y)

A -> {C, D}
But not exist
C -> D OR D -> C
```

```txt
FD X -> Y: is bridge if
X -> Y, Y -> Z
=> X -> Z
```

**Superkey**

Superkey exists only one in a table

{superkey} -> unique record

**Candidate key**

Candidate key as same as a key, is a minimal superkey

K is a candidate key <=>
+ K is unique
+ K is minimal. ∉ K' : K' ⊂ K, K' is unique

Primary Key is choosen from {CandidateKey}

**Subkey**

Subkey ⊂ CandidateKey

Candidate Key: {A, B}

Sub Key: {A}, {B}, {A, B}, {}

**Partial Dependency - PD**

FD A -> B is PD when A -> B and Ai -> B (A is a candidate key)

1) Identify Candidate keys
2) Differentiate prime and non-prime attributes
3) Find a non-prime attribute that is determined by a part or proper subset of candidate key

https://www.youtube.com/watch?v=pYASiat3eQg

## 1NF

- Scalable Table design which can be easily extended.
- If your table is not even in 1st Normal Form, its considered poor DB design.
- Every Table in your Database should at least follow the 1st Normal Form, always or Stop using Database!
- There are 4 basic rules that a table should follow to be in 1st Normal Form.

Rule 1:

- Each column should contain atomic values. Entries like X, Y and W, X violate this rule.

Rule 2:

- A column should contain values that are of the same type. Do not inter-mix different types of values in any column.

Rule 3:

- Each column should have a unique name. Same names leads to confusion at the time of data retrieval.

Rule 4:

- Order in which data is saved doesn't matter. Using SQL query, you can easily fetch data in any order from a table.

R(A_, B, {C}, D) => R(A_, B, D) AND R(A_, C)

## 2NF

- It should be in 1st Normal Form.
- And . It should not have any Partial Dependencies.

R((A, B), C, D)

{A, B} -> C
{A} -> D

R((A, B), C)
R((A), D)

FD X -> Y

X Superkey

Y Subkey

X ! subkey

A B C D
1 1 2 1
1 2 2 2
