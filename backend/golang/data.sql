select
  scci.app_name as app_name,
  scci.table_name as impala_table,
  FROM_UNIXTIME (checkpoint_time / 1000) as max_impala_ts_ms,
  binlog_monitor.last_ts_ms as max_source_ts_ms,
  TIMESTAMPDIFF (
    SECOND,
    FROM_UNIXTIME (checkpoint_time / 1000),
    binlog_monitor.last_ts_ms
  ) as delay_seconds
from
  bigdata.snapshot_cdc_config_ingest scci
  JOIN bigdata_monitor.binlog_monitor binlog_monitor ON scci.topic_name = binlog_monitor.topic
where
  scci.status = 1
  and binlog_monitor.last_ts_ms is not null
  and TIMESTAMPDIFF (
    SECOND,
    FROM_UNIXTIME (checkpoint_time / 1000),
    binlog_monitor.last_ts_ms
  ) >= 21600
  and binlog_monitor.last_ts_ms >= ?
  and checkpoint_time >= ?;

select
  scci.app_name as app_name,
  scci.table_name as impala_table,
  FROM_UNIXTIME (checkpoint_time / 1000) as max_impala_ts_ms,
  binlog_monitor.last_ts_ms as max_source_ts_ms,
  TIMESTAMPDIFF (
    SECOND,
    FROM_UNIXTIME (checkpoint_time / 1000),
    binlog_monitor.last_ts_ms
  ) as delay_seconds
from
  bigdata.snapshot_cdc_config_ingest scci
  JOIN bigdata_monitor.binlog_monitor binlog_monitor ON scci.topic_name = binlog_monitor.topic
where
  scci.status = 1
  and binlog_monitor.last_ts_ms is not null
  and TIMESTAMPDIFF (
    SECOND,
    FROM_UNIXTIME (checkpoint_time / 1000),
    binlog_monitor.last_ts_ms
  ) >= 21600;

select
  scci.app_name as app_name,
  scci.table_name as impala_table,
  FROM_UNIXTIME (checkpoint_time / 1000) as max_impala_ts_ms,
  binlog_monitor.last_ts_ms as max_source_ts_ms,
  TIMESTAMPDIFF (
    SECOND,
    FROM_UNIXTIME (checkpoint_time / 1000),
    binlog_monitor.last_ts_ms
  ) as delay_seconds
from
  (
    select
      *
    from
      bigdata.snapshot_cdc_config_ingest scci
    where
      checkpoint_time <= UNIX_TIMESTAMP () - 21600
  ) scci
  JOIN bigdata_monitor.binlog_monitor binlog_monitor ON scci.topic_name = binlog_monitor.topic
where
  scci.status = 1
  and binlog_monitor.last_ts_ms is not null
  and TIMESTAMPDIFF (
    SECOND,
    FROM_UNIXTIME (checkpoint_time / 1000),
    binlog_monitor.last_ts_ms
  ) >= 21600
