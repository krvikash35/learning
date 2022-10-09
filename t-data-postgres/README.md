
pg_dump

```
pg_dump -h host -p port -U username -d db -f dump.sql
```

hba file
```
show hba_file;
vi /etc/postgresql/11/main/pg_hba.conf
host    all             all             10.0.0.0/8            md5
sudo systemctl reload postgresql@11-main
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

INSERT INTO birds (bird , description) VALUES 
('pigeon', 'common in cities')
('eagle', 'bird of prey');
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