# Index Selectivity

Index selectivity is the ratio of the number of distinct indexed values (the cardinality) to the total number of rows in the table (#T), and ranges from 1/#T to 1. A highly selective index is good because it lets MySQL filter out more rows when it looks for matches. A unique index has a selectivity of 1, which is as good as it gets.
