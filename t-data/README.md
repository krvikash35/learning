
* postgres
* patroni postgres
* mongo
* cassandra
* spanner
* redis
* rabbitmq
* kafka
* bigquery
* bigtable
* hadoop
* spark


## todo
* ACID (atomic, consistency, isolation, durability)
* CAP (consistency, availability, partition)
* DBas
* sql v non-sql
* txn vs non-txn
* scaling
* replication
* sharding
* partitioning
* federation
* clustering
* usecase(read, write, process, analytic, hugeData, streaming, caching, event)

sql order
```
(8)  SELECT (9) DISTINCT (11) TOP <top_specification> <select_list>
(1)  FROM <left_table>
(3)       <join_type> JOIN <right_table>
(2)       ON <join_condition>
(4)  WHERE <where_condition>
(5)  GROUP BY <group_by_list>
(6)  WITH {CUBE | ROLLUP}
(7)  HAVING <having_condition>
(10) ORDER BY <order_by_list>
```