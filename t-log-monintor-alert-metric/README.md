load
* load1: avg no of process waiting in last 1 min
* load5
* load15

```
23:16:49 up  10:49,  5 user,  load average: 1.00, 0.40, 3.35

ON 1cpu core: 
    cpu was 100% utilized. 1 process were running on cpu over last 1 minute.
    cpu was idle 60% on averge. no process were waiting for cpu over last 5 minute.
    cpu was overloaded by 235% on average. 2.35 processes were waiting over last 15m
ON 2cpu core: 
    1cpu 100% idle. no cpu were waiting for cpu time over last 1min
    idle 60% on avg. no process wait for cpu time over last 5m
    overload by 135% on avg. 1.35 process wait for cpu time over last 15m
```
usages: cpu time being used by various process
* system: kernel/system space program.
* user: user space code like app. take most of the time
* nice: subset of user state. process with low priority.
* idle: doing nothing
* iowait: waiting for disk read/write. if spike problem outside cpu


## disk
* sda, sda1, sda2
* linux auto assign the number to hard disk. these are block device other are char
* if single hard disk contain partion then those partition will also be assigned. one disk have max 15 part

```
$ sudo fdisk -l
Disk /dev/sda: 465.76 GiB, 500107862016 bytes, 976773168 sectors
Device     Boot     Start       End   Sectors   Size Id Type
/dev/sda1  *         2048    104447    102400    50M  7 HPFS/NTFS/exFAT
/dev/sda2          104448 484152681 484048234 230.8G  7 HPFS/NTFS/exFAT
/dev/sda3       484155392 485249023   1093632   534M 27 Hidden NTFS WinRE
/dev/sda4       485251070 976771071 491520002 234.4G  5 Extended
/dev/sda5       485251072 974772223 489521152 233.4G 83 Linux
/dev/sda6       974774272 976771071   1996800   975M 82 Linux swap / Solaris

$ lsblk
NAME   MAJ:MIN RM   SIZE RO TYPE MOUNTPOINTS
sda      8:0    0 465.8G  0 disk 
├─sda1   8:1    0    50M  0 part 
├─sda2   8:2    0 230.8G  0 part 
├─sda3   8:3    0   534M  0 part 
├─sda4   8:4    0     1K  0 part 
├─sda5   8:5    0 233.4G  0 part /
└─sda6   8:6    0   975M  0 part [SWAP]
sr0     11:0    1  1024M  0 rom 
```

disk io request per sec
* diskio_reads 
* diskio_writes 

* diskio_write_bytes 
* diskio_read_bytes

* diskio_io_time(percent io util): 
* diskio_weighted_io_time(avg que depth): here
    * The disk queue length reports on the number of outstanding operations to a particular volume
    * A good rule of thumb is that there should never be more than half the number of spindles in the queue length. If you have a 10-disk RAID volume, the queue length should be less than 5

* diskio_read_time
* diskio_write_time
* disk_total/disk_used/disk_used_percent


a portion of the filesystem is dedicated to inodes. An inode is a data structure that describes a file or a folder.It includes things like the owner, the group, permissions, file size, created etc. To check how many inodes are in use and free, use df -i
* disk_inodes_free
* disk_inodes_used

## memory
usages: 
* total: Your total (physical) RAM (excluding a small bit that the kernel permanently reserves for itself at startup); that's why it shows ca. 11.7 GiB , and not 12 GiB, which you probably have.
* used: memory in use by the OS
* available: total - used
* free: memory not in use.
* shared, buffered, cache: This shows memory usage for specific purposes, these values are included in the value for used

## network

## tuples and vacuum
* autovacuum avoid tabel growwing unnecessarily. If a table does not have autovacuum running, it will have to keep growing to store new data somewhere because if data is being deleted then it can not be reused unless autovacuum runs.
* autovacuum running frequent is good thing unless it take itself long time
* dead row/tupple: on delete mark the row as dead, on update: insert new row & mark old row as dead
* live row/tupple: current version of rows that are avilable for txn/query


## golang
* goroutines count: On a machine with 4 GB of memory installed, this limits the maximum number of goroutines to slightly less than 1 million.
* GC Pause time in ms: stop the world. happens when a region of memory is full and the JVM requires space to continue
* GC Pause frequency: calls per minute
* GC Pauses as % of wall-clock time
* cpu % utilization(ueser & system)
* memory allocation


## newrelic
web vs non-web txn
* When you instrument or wrap a transaction that has an HTTP request and response writer, New Relic treats it as a web transaction.
* When you instrument or wrap a transaction that does not have HTTP data, New Relic treats it as a non-web transaction.
golang: monitor txn: [here](https://docs.newrelic.com/docs/apm/agents/go-agent/instrumentation/instrument-go-transactions/)
```
txn := app.StartTransaction("transaction_name")
defer txn.End()

To monitor a web transaction, call the Transaction.SetWebRequest and optionally the Transaction.SetWebResponse APIs:
// req is a *http.Request, this marks the transaction as a web transaction
txn.SetWebRequestHTTP(req)

// writer is a http.ResponseWriter, use the returned writer in place of the original
writer = txn.SetWebResponse(writer)
writer.WriteHeader(500)
```