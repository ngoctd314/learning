# SQL Processing

## Introduction

We now have some understanding of the structure of tables and indexes; we also understand how these objects relate to buffer pools, disks, and disk servers and how the latter are used to make the data available to the SQL process. We are now in a position to consider the processing of the SQL calls.

**PREDICATES**

A WHERE clause consists of one or more predicates (search arguments).

## OPTIMIZERS AND ACCESS PATHS

One of the long-standing advantages of relational databases has been that data is requested with little or no thought for the way in which the data is to be accessed. This decision is made by a component of the DBMS called the optimizer. These have varied widely across different relational systems and probably always will, but they all try to access the data in the most effective way possible, using statistics stored by the system collected on the data.

Before an SQL statement can be executed, the optimizer must first decide how to access the data; which index, if any, should be used; how the index should be used; should assisted random read be used; and so forth.

**Index Slices and Matching Columns**

Thus a thin slice of the index.

**Index Screening and Screening Columns**

## Filter factors

The filter factor specifies the selectivity of a predicate - what proportion of the source table rows satisfy the condition expressed by the predicate.

**Filter Factors for Compound Predicates**

The filter factor for a compound predicate can be derived from the filter factors of the simple predicates.
