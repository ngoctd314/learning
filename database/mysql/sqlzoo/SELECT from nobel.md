# SELECT from Nobel Tutorial

https://sqlzoo.net/wiki/SELECT_from_Nobel_Tutorial

**14. Show the 1984 winners and subject ordered by subject and winner name; but list chemistry and physics last.**

```sql
SELECT winner, subject FROM nobel WHERE yr = 1984
ORDER BY subject in ('chemistry', 'physics'), subject, winner
```
