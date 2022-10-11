
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


