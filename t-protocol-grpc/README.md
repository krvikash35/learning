* typical client server model, uses tcp, more efficient/performant compare to `http`. Infact it uses http2.
* It is best suited to be used with `protobuf` data format, which is highly efficient, lightweight data format.
* Though we can use json with grpc but we will only be able to utilize benefit of `grpc`, not the benefit of protobuf as well.
* we can use `gcurl` simmilar to `curl` but grpc is not that portable/debugable like we do http using curl.


* service can't return primitive(`bool`, `string`), it has to be either `message` or `void`type
