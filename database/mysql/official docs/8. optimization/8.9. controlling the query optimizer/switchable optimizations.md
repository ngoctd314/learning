# Switchable Optimizations

The optimizer_switch system variable enables control over optimizer behavior.

```sql
  SELECT @@optimizer_switch\G
***************************[ 1. row ]***************************
@@optimizer_switch | index_merge=on,index_merge_union=on,index_merge_sort_union=on,index_merge_intersection=on,engine_condition_pushdown=on,index_condition_pushdown=on,mrr=on,mrr_cost_based=on,block_nested_loop=on,batched_key_access=off,materialization=on,semijoin=on,loosescan=on,firstmatch=on,duplicateweedout=on,subquery_materialization_cost_based=on,use_index_extensions=on,condition_fanout_filter=on,derived_merge=on,use_invisible_indexes=off,skip_scan=on,hash_join=on,subquery_to_derived=off,prefer_ordering_index=on,hypergraph_optimizer=off,derived_condition_pushdown=on
```

To change the value of optimizer_switch, assign a value consisting of a comma-separated list of one or more commands:

```sql
SET [GLOBAL|SESSION] optimizer_switch='command[,command]...'
```

|Comamnd Syntax|Meaning|
|-|-|
|default|Reset every optimization to its default value|
|opt_name=default|Set the named optimization to its default value|
|opt_name=off|Disable the named optimization|
|opt_name=on|Enable the named optimization|


index_merge=on,index_merge_union=on,index_merge_sort_union=on,index_merge_intersection=on,engine_condition_pushdown=on,index_condition_pushdown=on,mrr=on,mrr_cost_based=on,block_nested_loop=on,batched_key_access=off,materialization=on,semijoin=on,loosescan=on,firstmatch=on,duplicateweedout=on,subquery_materialization_cost_based=on,use_index_extensions=on,condition_fanout_filter=on,derived_merge=on,prefer_ordering_index=on,favor_range_scan=off
