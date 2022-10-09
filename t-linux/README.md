
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


Edit
```
vi
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

less
* `Down Arrow, Enter, e, j`    One line forward.
* `Up Arrow, y, k`                        One line backward.
* `Space bar, Page Down`           One page forward.
* `Page Up, b`                                 One page backward.
* `Right Arrow`                               Scroll right.
* `Left Arrow`	                                Scroll left.
* `/[string]`                                   search forward for string
* `?/[string]`                                 search backward for string
* `n`                                                     next match
* `N`                                                     previous match
* `q`                                                     quite

```bash
less log.txt
less -N log.tx          # show line no
less +F log.txt         # real time monitoring 
less -M log.txt         # statistic line no, % 
```