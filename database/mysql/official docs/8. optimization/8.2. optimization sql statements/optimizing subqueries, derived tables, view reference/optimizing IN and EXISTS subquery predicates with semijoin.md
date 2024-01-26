# Optimizing IN and EXISTS Subquery Predicates with Semijoin Transformations

A semijoin is a preparation-time transformation that enables multiple execution strategies such as table pullout, duplicate weedout, first match, loose scan, and materialization. The optimizer uses semijoin strategies to improve subquery execution, as described in this section.

For an inner join between two tables, the join returns a row from one table as many times as there are matches in the other table. But for some questions, the only information that matters is whether there is a match, not the number of matches.

```sql
SELECT class.class_num, class.class_name
FROM class INNER JOIN roster
WHERE class.class_num = roster.class_num;
```

However, the result lists each class once for enrolled student. For the question being asked, this is unnecessary duplication of information.

Assuming that class_num is a primary key in the class table, duplicate suppression is possible by using SELECT DISTINCT, but it is inefficient to generate all matching rows first only to eliminate duplicates later. 

The same duplicate-free result can be obtained by using a subquery:

```sql
CREATE TABLE class (id int AUTO_INCREMENT PRIMARY KEY, name varchar(255));

CREATE TABLE roster (id int AUTO_INCREMENT PRIMARY KEY, class_num int);

INSERT INTO class (name) VALUES ('math'), ('bio'), ('literature'), ('geo');

INSERT INTO roster (class_num) VALUES (1), (1), (2), (2);

SELECT class.id, class.name
    FROM class
    INNER JOIN roster ON id
    WHERE class.id = roster.class_num;
+----+------+
| id | name |
+----+------+
| 1  | math |
| 1  | math |
| 2  | bio  |
| 2  | bio  |
+----+------+
4 rows in set

SELECT DISTINCT class.id, class.name
    FROM class
    INNER JOIN roster on id
    WHERE class.id = roster.class_num;

+----+------+
| id | name |
+----+------+
| 1  | math |
| 2  | bio  |
+----+------+
```

However, the result lists each class once for each enrolled student. For the question being asked, this is unnecessary duplication of information.

Assuming that class_num is a primary key in the class table, duplicate suppression is possible by using SELECT DISTINCT, but it is inefficient to generate all matching rows first only to eliminate duplicates later.

The same duplicate-free result can be obtained by using a subquery:

```sql
SELECT class.id, class.name
    FROM class
    WHERE id IN
        (SELECT class_num FROM roster);
```

Here, the optimizer can recognize that the IN clause requires the subquery to return only one instance of each class number from the roster table. In this case, the query can use a semijoin; that is, an operation that returns only one instance of each row in class that is matched by rows in roster.


The following statement, which contains an EXISTS subquery predicate, is equivalent to the previous statement containing an IN subquery predicate:

```sql
SELECT id, name
FROM class 
WHERE EXISTS 
    (SELECT * FROM roster WHERE class.id = roster.class_num)
```

In MySQL 8.0.16 and later, any statement with an EXISTS subquery predicate is subject to the same semijoin transforms as a statement with an equivalent IN subquery predicate.
