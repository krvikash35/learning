
* Each segment of the WAL log size is `16 MB`. if` wal_keep_segments=15` then 15*16=240MB. by default it wal keep seg is zero.
* wal_level determines how much information is written to the WAL: 
  * `replica`: which writes enough data to support WAL archiving and replication, including running read-only queries on a standby server
  * `minimal`:  removes all logging except the information required to recover from a crash or immediate shutdown
  * `logical` : adds information necessary to support logical decoding
* `confirmed_flush_lsn`: The address (LSN) up to which the logical slot's consumer has confirmed receiving data. Data older than this is not available anymore. NULL for physical slots.
* `restart_lsn`: The address (LSN) of oldest WAL which still might be required by the consumer of this slot and thus won't be automatically removed during checkpoints unless this LSN gets behind more than max_slot_wal_keep_size from the current LSN. NULL if the LSN of this slot has never been reserved.


```
\l
\c dbname
\du
\d tablename
\d
\d+ tablename
\dx -> extension list
\df -> function list
```


**Partman**
```
SELECT "partman".run_maintenance(p_parent_table:='table', p_analyze := false, p_jobmon := true);

SELECT "partman".run_maintenance( p_analyze := false, p_jobmon := true);

call "partman".partition_data_proc(p_parent_table := 'table')
```


```
pg_dump -h host -p port -U username -d db -f dump.sql
```

hba file
```
show hba_file;
vi /etc/postgresql/11/main/pg_hba.conf
host    all             all             10.0.0.0/8            md5
sudo systemctl reload postgresql@11-main

select table_name, pg_relation_size(quote_ident(table_name)) from information_schema.tables where table_schema = 'public' order by 2;
pg_is_in_recovery();

select * from pg_replication_slots ;
select * from pg_settings where name like '%autovacuum%';
select sum(size) from pg_ls_waldir();
SELECT sum(numbackends) FROM pg_stat_database;

EXPLAIN ANALYZE SELECT * FROM users ORDER BY username;
```

```
SELECT
  pid,
  user,
  pg_stat_activity.query_start,
  now() - pg_stat_activity.query_start AS query_time,
  query,
  state,
  wait_event_type,
  wait_event
FROM pg_stat_activity
WHERE (now() - pg_stat_activity.query_start) > interval '5 minutes';

pg_cancel_backend(pid);
pg_terminate_backend(pid);
```

```
select pid, 
       usename, 
       pg_blocking_pids(pid) as blocked_by, 
       query as blocked_query
from pg_stat_activity
where cardinality(pg_blocking_pids(pid)) > 0;
```

```
CREATE TABLE users (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username text NOT NULL);
INSERT INTO users (username)
SELECT 'person' || n
FROM generate_series(1, 1000) AS n;
ANALYZE users;
```



* dd
* psql -h localhost -p 5432 -U rolename dbname connect to db using psql client
* \conninfo get current DB and user info
* \du list all the role/user
* \l list databases
* \c db_name connect to db
* \dt list table of current schema which is by default public. equivalent to \dt public.*
* \dt *.* list table of all schema
* \dt inventory.* list tables of schema whose name start with inventory
* \dn list schema
* \o out.tx toggle output file by \o
* SET search_path TO public, inventory  it help to avoid putting schema name prefix
* show search_path; see current value (list of schema) in search path


crud
```
select * from users limit 1;
select id, name from users where age > 7 and name like 'vik%' order by age asc;

operator: =, <, <=, >, >=, between, in, like 

SELECT array_to_json(array_agg(row_to_json(t))) FROM 
    (select id, order_number from example_table limit 2) t
```

```
update users set name='vikash' and age='20' where id=12
```

```
CREATE DATABASE bird_encyclopedia;
\c bird_encyclopedia

CREATE TABLE birds (
  id SERIAL PRIMARY KEY,
  bird VARCHAR(256),
  description VARCHAR(1024)
);

INSERT INTO birds (bird , description) VALUES ('pigeon', 'common in cities'), ('eagle', 'bird of prey');
```

* Neither TIMESTAMP variants store a time zone (or an offset), despite what the names suggest.
* TIMESTAMP WITHOUT TIME ZONE stores local date-time
* TIMESTAMP WITH TIME ZONE stores a point on the UTC time line


if we put condition in where clause then filter happens after join but if we put in on clause then filter happen before join
```
SELECT t1.col1, t1.col2, t2.col3
  FROM table1 t1
  LEFT JOIN table2 t2
    ON t1.col1 = t2.col1
    AND t1.col2 != 'some value'
    
SELECT t1.col1, t1.col2, t2.col3
  FROM table1 t1
  LEFT JOIN table2 t2
    ON t1.col1 = t2.col1
  WHERE t1.col2 != 'some value'    
```
autovaccum & dead tupple

* autovacuum running frequent is good thing unless it take itself long time
* autovacuum avoid tabel growwing unnecessarily. If a table does not have autovacuum running, it will have to keep growing to store new data somewhere because if data is being deleted then it can not be reused unless autovacuum runs.
* system_load1/5/15 is avg system cpu load over given period of 1, 5 & 15 minute.

```
23:16:49 up  10:49,  5 user,  load average: 1.00, 0.40, 3.35

1cpu core: 
    cpu was 100% utilized. 1 process were running on cpu over last 1 minute.
    cpu was idle 60% on averge. no process were waiting for cpu over last 5 minute.
    cpu was overloaded by 235% on average. 2.35 processes were waiting over last 15m
2cpu core: 
    1cpu 100% idle. no cpu were waiting for cpu time over last 1min
    idle 60% on avg. no process wait for cpu time over last 5m
    overload by 135% on avg. 1.35 process wait for cpu time over last 15m
```


```
create database <dbname>
create user <username> SUPERUSER
alter user <username> WITH PASSWORD <password>
```

## links
* [failover](https://severalnines.com/blog/failover-postgresql-replication-101)
* [replication](https://severalnines.com/blog/postgresql-streaming-replication-deep-dive)
* [logical replication](https://severalnines.com/blog/overview-logical-replication-postgresql)