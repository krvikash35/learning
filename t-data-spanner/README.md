
```
gcloud config configurations  list
gcloud config configurations activate emulator
gcloud config configurations create emulator
gcloud config set auth/disable_credentials true
gcloud config set project your-project-id
gcloud config set api_endpoint_overrides/spanner http://localhost:9020/

gcloud spanner instances list
gcloud spanner instances create test-instance --config=emulator-config --description="Test Instance" --nodes=1



gcloud spanner databases list --instance=test-instance
gcloud spanner databases create testdb --instance=test-instance
gcloud spanner databases execute-sql testdb  --instance=test-instance --sql='SELECT * FROM mytable'
```


```
var p Person

stmt := spanner.Statement{
		SQL: `SELECT * from person where name=@c1`,
		Params: map[string]interface{}{
			"name":     "vikash",
		},
	}
	iter := s.spannerClient.Single().Query(context.Background(), stmt)
	defer iter.Stop()

	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			break
		}
		err = row.ToStruct(&p)
		if err != nil {
		}
		fmt.Printf("Person %+v", p)
	}
```

```
mutations, err := spanner.InsertStruct("person", person)
client.Apply(ctx, mutations)
```

```
dsn:= projects/p1/instances/inst1/databases/db1
spanner.newclient(ctx, dsn, credentialJson)
```