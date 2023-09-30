
## security

```
security add-generic-password -a my-sample-account -p "my-sample-password" -s my-sample-service
security find-generic-password -a my-sample-account -w 2>/dev/null | tr -d '\n'
security delete-generic-password -a my-sample-account
```


## networksetup
```

```

## scutil
```
scutil --nc list
scutil --nc start My_L2TP_VPN
```

apple script to connect to L2TP VPN
```applescript
set MyVPNName to "My_L2TP_VPN"

tell application "System Events"
	tell current location of network preferences
		set myConnection to the service MyVPNName
		if myConnection is not null then
			if current configuration of myConnection is not connected then
				connect myConnection
			end if
		end if
	end tell
end tell
```

## VPN connection
* To increase productivity, we want to programatically control the vpn connection for auto-connect/re-connect use case.
* `scutil` is used to connect to VPN services but it does not support `IKEv2` VPN, hence many resort to apple script. Read more [here](https://blog.timac.org/2018/0719-vpnstatus/).

