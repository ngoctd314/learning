SELECT
  query_database.db_name AS db_name,
  query_schema.database_id AS database_id,
  query_user.name AS username,
  slow_query.query_pattern_id AS query_pattern_id,
  query_database.rate_config_id AS rate_config_id,
  query_user.group_id AS user_group_id,
  query_database.group_id AS database_group_id,
  dsa_group.name AS group_name,
  COUNT(*) AS total
FROM
  slow_query
  JOIN query_pattern ON slow_query.query_pattern_id = query_pattern.id
  JOIN query_schema ON query_schema.id = query_pattern.schema_id
  JOIN query_database ON query_schema.database_id = query_database.id
  JOIN query_user ON slow_query.user_id = query_user.id
  JOIN dsa_group ON query_user.group_id = dsa_group.id
WHERE
  slow_query.created >= ?
  AND slow_query.created <= ?
GROUP BY
  slow_query.query_pattern_id;
