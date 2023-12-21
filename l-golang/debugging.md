debug flag

```
-gcflags="all=-N -l"

go build -gcflags="all=-N -l" -o myApp
go build -tags dynamic -> to have confluent-kafka-go link librdkafka dynamically on m1

arch -arm64 /opt/homebrew/bin/brew install librdkafka -> to install arm64 version
```

command, _ := http2curl.GetCurlCommand(req)
fmt.Println()
fmt.Println("curl_request", command)
fmt.Println()

go get moul.io/http2curl



debug
--------

running process attach to process
* VScode
* Goland


executable 
* VSCode

```
{
    "configurations": [
    {
        "name": "Launch file",
        "type": "go",
        "request": "launch",
        "mode": "exec",
        "program": "${workspaceFolder}/out/transport-pricing-service",
        "args": ["migrate"],
        "env": {
            "BCS_ENABLED": "false"
        },
        "cwd": "${workspaceFolder}",
        
    }
    ]
}
```
* Goland
