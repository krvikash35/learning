
* vi
* ls
* find
* grep
* tail
* head
* awk
* sed
* less
* morez
* nl
* date
* nc


## vi

```
i       insert mode
escp    command mode
:w      save file
:q      quite file
:wq     save file and quit
:q!     quit forcfully without saving
:set nu show line number


W       move forward one word
b       move backward one word
0       move start of line
$       move end of line
H       move to top of screen
M       move to middle of screen
L       move to bottom of screen
G       move to end of doc
8G      move to line no 8

/text   search forward
?text   search backward
n       next match
N       previous match

o       open new line below
O       open new line above

dw      delete word
dl      delete leter
D       delete line         
```

Search, navigate
```

```

find
* File Types: `f` - regular, `d` - directory, `l`-symbLink, `c` - charDevice, `b` - blockDevice
* Options
    * `-type [fileType]` 
    * `-name [pattern]`
    * `-iname [pattern]` : incase sensitive
    * `-size +10MB`
```
find [where] [option] [what]
find /usr/local  -iname  "*elastic*"
find /usr/local -type c -iname  "*elastic*"
```

ls
* options
    * `-h`  human readable
    * `-a` all including hidden file
    * `-l` long detail
    * `-t` sort by modified time
    * `-r` reverse the sorting
    * `-s` no of block in file

```
ls -halt
```

tail
* options
    * 

```bash
tail file.log                # by default show last 10 line
tail -n 20 file.log          # last 20 line
tail -n 20 f1.log f2.log     # last 20 line
tail -n 20 f1.log f2.log     # last 20 line
tail -c 500 file.log         # last 500 byte
tail -20 file.log            # last 20 line
```

process
```

```

networking
```
tailnet host port
nc host port
echo "hello" | nc host port
```

disk
```

```

input, output, pipe
```
ls -lart| grep [text]
ls > out.txt
```

date
* utc
* gmt
* iso 8601: format, human friendly, unlimited. 
    * `2020-05-09T09:17:46+05:30`  offset of +5.30(IST)
    * `2020-05-09T13:49:54+00:00`
    * `2020-05-09T13:49:54Z` 
* unix : format, machine friendly, limited, no of seconds elapsed since 1,Jan,1970(midNight UTC/GMT). 

```bash
date -d"2014-02-14T12:30" +%s   # convert iso8601 to unix
date -r 1234567890              # convert unix to iso8601
```

## less
See file content one screen at a time.
* `less file.txt`

Options
```

```

Navigations
```
Down Arrow, Enter, e, j    One line forward.
Up Arrow, y, k             One line backward.
Space bar, Page Down       One page forward.
Page Up, b                 One page backward.
G                          go to end of file
g                          go to start of file
Right Arrow                Scroll right.
Left Arrow	               Scroll left.
/[string]                  search forward for string
?/[string]                 search backward for string
n                          next match
N                          previous match
q                          quite
-S                         wrap/unwrap line
10g                        go line no 10
-I                         toggle case sensitivity
F                          live mode tail like, Ctrl+C to escape
:[a]                       mark current pos with letter a
'a                         go to pos marked with a
```


Examples

```bash
less log.txt
less -N log.tx          # show line no
PAGER="less" psql       # change the default pager used psql
```

## Terminal Navigations
* Configure Iterm to use **Natural Text Editing** preset in `Preferences->Profiles->KeysMappings->Preset`
```
option + >            move one word forward
option + <            move one word backward
cmd + <               move to start of line
cmd + >               move to end of line
option + del          delete word
```

## grep
```bash
grep "pattern1\|pattern2"
```



```
journalctl
journalctl :all logs
journalctl -r :to reverse new entries first
journalctl -n 2 :to view 2 log entries
journalctl -p warning
journalctl | grep Centaur
journalctl -o verbose
journalctl --list-boots
journalctl -u service-name.service
journalctl -u service-name --no-pager
journalctl --list-boots
```

## systemctl
* `service` is based on old init system /etc/init.d. 
* `systemctl` operate on file based in  /lib/systemd if not fallback to /etc/init.d
```
systemctl reboot
systemctl poweroff
systemctl list-units --type=service
systemctl list-units --type=service --state=active[running,stopped,enabled,disabled,failed]
systemctl list-units --failed
systemctl start|stop|restart|enable|disable {servicename} enable-onboot
systemctl is-enabled|is-active {servicename}
systemctl kill {servicename}
systemctl kill -s 9 {servicename}

service sshd status|start|stop|restart
```

## nc - netcat
* port scan: without sending data to server, it can check if port are open
* dummy chat client-server model
* to test server by sending dummy data 



Options
```
-l          # listen
-u          # use udp instead of tcp
-w  2       # timeout
-z          # only scan port without sending any data
-v          # verbose
```

Examples
```
nc -l 1234                          # start tcp server on localhost:1234
nc localhost 1234                   # connect to tcp server on localhost:1234

nc -zv localhost 20-30              # port scan b/w 20 & 30
nc -zv localhost 80 20 22           # port scan on 80, 20, 22
```

```
netstat --all --program | grep 6831
lsof -i:6831
netstat -l --> list only listening port

telnet host
telnet ip_address port
```

## Other
```
PAGER="less -S" psql
EDITOR=vim knife node edit hostname

ln -s existing_file_path new_file_path
ln -s /usr/local/Cellar/python@3.8/3.8.13_1/bin/python3 /usr/local/bin/python3.8

df -h
lsblk

scp host:hostPath localHostPath

lsof -i:9090
lsof -p 39518
top -pid $(pgrep Python | sed -e ':a' -e 'N' -e '$!ba' -e 's/\n/ -pid /g')

echo dkkd dkls > filename    # write to file, escape using " for special char
echo dkkd dkls >> filename   # append to file
cat > somefile               # write to file using stdin
```


```
echo "vikash \n rohit \n kohli" | fzf
echo "vikash kumar" | cut -d' ' -f1;
cat $PROJECT_LIST_FILE | fzf | awk '{print $2}';

name="vikash"
if [[ $name =~ "kas" ]]; then
  echo "substring match"
  exit 0
fi
```


```
wget https://github.com/tsenart/vegeta/releases/download/v12.11.0/vegeta_12.11.0_linux_amd64.tar.gz
tar xfz vegeta_12.11.0_linux_amd64.tar.gz
echo "GET http://localhost:8081/ping" | vegeta attack -duration=5s -rate=2 | vegeta report
```


```
jq - json cli processor, can filter
fx - json cli viewer
yh - yml cli viewer
```