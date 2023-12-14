--9. Find the continents where all countries have a population <= 25000000. Then find the names of the countries associated with these continents. Show name, continent and population.
SELECT
  name,
  continent,
  population
FROM
  world x
WHERE
  continent IN (
    SELECT
      continent
    FROM
      world
    GROUP BY
      continent
    HAVING
      MAX(population) <= 25000000
  );

-- 10. Some countries have populations more than three times that of all of their neighbours (in the same continent). Give the countries and continents.
SELECT
  name,
  continent
FROM
  world x
WHERE
  population > 3 * (
    SELECT
      MAX(population)
    FROM
      world y
    WHERE
      y.continent = x.continent
      and y.name <> x.name
  );
