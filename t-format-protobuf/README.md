## Syntax
```
syntax = "proto3";

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}

map<string, Project> projects = 3;


service SearchService {
  rpc Search(SearchRequest) returns (SearchResponse);
}
```

## Scalar type
* `double`: generally 8 byte. can hold up to 15 decimal digit for accuracy.
* `float`: generally 4 byte. can hold up to 7 decimal digit for accuracy.
* `int32`: both positive and negative but inefficient for -ve no, so use `sint32`. uses varint(two's encoding).
* `int64`: 
* `uint32`: only positive number.
* `uint64`: 
* `sint32`: both positive and negative. uses zig-zag varint for encoding: +v -> 2 * n, -ve -> 2 * n + 1.
* `sint64`: 
* `bool`
* `string`
* `bytes`


## Non-Scalar type
* enum
* map
* any
* oneof
* message
* service

## Field type
* `singular`
* `optional`
* `repeated`
* `required`
## Options
* option can be package, message, field level.


## Packages

## protobuf and language

## encoding
* bools and enum are encoded as int32. false -> 00 true -> 01.

## proto2 vs proto3
* proto3 remove support of `required` field.
* proto3 support cannonical encoding in json.