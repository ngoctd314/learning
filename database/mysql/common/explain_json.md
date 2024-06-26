# EXPLAIN FORMAT=JSON

The `EXPLAIN FORMAT=JSON` statement in MySQL is used to obtain the query execution plan in JSON format. This can be useful for more detailed analysis or if you want to process the execution plan programmatically. 

`rows_examined_per_scan` field might appear in the JSON output. In this example, the `rows_examined_per_scan` is set to 20, indicating that MySQL estimates examining approximately 20 rows per scan during the execution of the query.

This value helps you understand the efficiency of the chosen access method (such as index usage) and can valuable for query optimization. Lower values for `rows_examined_per_scan` generally indicate more efficient query execution.

```json
{
  "query_block": {
    "select_id": 1,
    "cost_info": {
      "query_cost": "1.00"
    },
    "nested_loop": [
      {
        "table": {
          "table_name": "tbl1",
          "access_type": "ALL",
          "rows_examined_per_scan": 1,
          "rows_produced_per_join": 1,
          "filtered": "100.00",
          "cost_info": {
            "read_cost": "0.25",
            "eval_cost": "0.10",
            "prefix_cost": "0.35",
            "data_read_per_join": "1K"
          },
          "used_columns": [
            "id",
            "name",
            "age"
          ]
        }
      },
      {
        "table": {
          "table_name": "tbl2",
          "access_type": "ALL",
          "rows_examined_per_scan": 4,
          "rows_produced_per_join": 1,
          "filtered": "25.00",
          "using_join_buffer": "hash join",
          "cost_info": {
            "read_cost": "0.25",
            "eval_cost": "0.10",
            "prefix_cost": "1.00",
            "data_read_per_join": "1K"
          },
          "used_columns": [
            "id",
            "name",
            "age"
          ],
          "attached_condition": "(`learn_explain`.`tbl2`.`name` = `learn_explain`.`tbl1`.`name`)"
        }
      }
    ]
  }
}
```
