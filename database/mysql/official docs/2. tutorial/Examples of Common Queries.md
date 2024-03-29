# Examples of Common Queries

## The Row Holding the Maximum of a Certain Column

Find the number, dealer, and price of the most expensive article

```sql
SELECT article, dealer, price
FROM shop
WHERE price=(SELECT MAX(price) FROM shop);
```

```txt
+---------+--------+-------+
| article | dealer | price |
+---------+--------+-------+
|    0004 | D      | 19.95 |
+---------+--------+-------+
```

Another solution is to use a LEFT JOIN, as show here
```sql
SELECT s1.article, s1.dealer, s1.price
FROM shop s1
LEFT JOIN shop s2 ON s1.price < s2.price
WHERE s2.article IS NULL;
```

You can also do this by sorting all rows descending by price and get only the first row using the MySQL-specific LIMIT clause

```sql
SELECT article, dealer, price
FROM shop
ORDER BY price DESC
LIMIT 1;
```
If there were several most expensive articles, each with a price of 19.95, the LIMIT solution would show only one of them.