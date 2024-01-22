# Optimizing Subqueries, Derived Tables, and View References with Semijoin Transformations

The optimizer uses semijoin strategies to improve subquery execution, as described in this section.

For an inner join between two tables, the join returns a row from one table as many times as there are matches in the other table. But for some questions, the only information that matters is whether there is a match, not the number of matches. Suppose that there are tables named class and roster that list classes in a course curriculum and class rosters. To list the classes that actually have students enrolled, you could use this join: 

```sql
SELECT class.class_num, class.class_name
FROM class INNER JOIN roster
WHERE class.class_num = roster.class_num;
```

However, the result lists each class once for each enrolled student. For the question being asked, this is unnecessary duplication of information.

The same duplicate-free result can be obtained by using subquery:

```sql
SELECT class_num, class_name
FROM class
WHERE class_num IN (SELECT class_num FROM roster);
```
