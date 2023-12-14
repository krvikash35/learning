
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