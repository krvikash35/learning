
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

disk
```
df -ha  -> complete disk usage in human format
df -T  -> disk usage by filesystem type

du -ha /opt -> all file
du -hs /opt -> specific file/dir

fdisk -l
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

```
1702716596                  Timestamp 
1702716596441               Timestamp in milliseconds 
2023-12-16T08:49:56Z        ISO 8601
2023-12-16T14:19:56+05:30   ISO 8601 with IST tz
16 Dec 2023, 08:49:56       Date Time(UTC)
16 Dec 2023, 14:19:56       Date Time(IST)

```

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
Space bar, Page Down, f       One page forward.
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

## journalctl

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
journalctl -b               -> since last boot
journalctl -b -1            -> since 2nd last boot

journalctl --since "1 hour ago"
journalctl --since "2 days ago"
$journalctl --since "2015-06-26 23:15:00" --until "2015-06-26 23:20:00"

journalctl -u nginx.service         -> by unit
journalctl -f                       -> like tail command, live log
journalctl -n 50                    -> latest 50 entries
journalctl -r                       -> reverse, sort by timestamp
journalctl -o json-pretty|verbose|cat
```



## Networking tool 
* nc - netcat
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
 nc -l 1234                         # start tcp server on localhost:1234
nc localhost 1234                   # connect to tcp server on localhost:1234

nc -zv localhost 20-30              # port scan b/w 20 & 30
nc -zv localhost 80 20 22           # port scan on 80, 20, 22

echo "hello" | nc host port
echo -n "testkey:500|g" | nc -w 1 -u -4 localhost 8125
```


```
netstat -a -> show all port 
netstat -l -> list only listening port
netstat -i -> list interface
ip link show
```

```
lsof -i:6831
```

```
telnet host
telnet host/ip port
```

```
ngrep -d lo0 port 3001
ngrep -d any
ngrep host localhost port 3001

ngrep -W byline -q 'http' | grep '/driver_found'  -C 50
ngrep -W byline -q -d any -n 2000 'int-api.marketplace.golabs.io' port 80 or port 443
ngrep -W byline -q -d any port 80 or port 443 | grep -C 100 'int-api.marketplace.golabs.io'
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
lsof -Pan -p 12574 -i
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

tar
```
tar -xvzf jruby-dist-9.4.3.0-bin.tar.g
```

tty vs stty vs psedu-ttyps vs terminal/console vs shell

* `tty` -> teletype. kind of device file that act like a terminal. real tty connect to physical terminal.
* `pty` -> psedo teletype. act like terminal to process reading/writing. pty connect to program eg xterm, cmd window or shell window. befor pty, pipe was ther but no control, so pty solve this.
* some tty are provided by kernel on behalf of hardware device. i.e input coming from keyboar & output going to text mode.
* `pseduo-ttys`` are provided by programs called terminal emulator such a xterm
`shell` -> i.e bash, zsh a shell interpreter, control tty to offer user experience


```
ssh vikash.kumar@dbhost << EOF
sudo su
psql -U postgres
\c dbname
SELECT array_to_json(array_agg(row_to_json(t))) FROM
    (select id, order_number from tablename limit 2) t;
EOF
```

Pipe & input direction
```
echo 'This is a test.\n line2' | wc -w

wc -w <<EOF
This is a test.
Apple juice.
100% fruit juice and no added sugar, colour or preservative.
EOF
```

apt-get install -y postgresql-client -> on ubuntu machine
apk add postgresql-client            -> on pod with alpine


## top

```
top 
shift + m -> sort by mem
shitf + p -> sort by cpu
```


`<<` is known as heredoc
```
cat
```

write to file
* `echo hello > out.txt` - write to file
* `cat > out.txt` - write to file from terminal, press ctrl + c once done
* `cat <<EOF > out.txt` -write to file from terminal, enter EOF once done