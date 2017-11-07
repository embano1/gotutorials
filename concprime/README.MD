# Built to benchmark cgroups quota (limits) and the effect of GOMAXPROCS.  
To simply run benchmark (defaults to 1000 ints):

```go
go test -bench=. -cpu 1,2,4,8 -benchtime=10s
```
Note: `-cpu` flag in the go testing framework allows to adjust `GOMAXPROCS`. Otherwise GOMAXPROCS defaults to `runtime.NumCPU()`.  

Play with cgroup settings, e.g. `docker run --cpus 2 [...]` and compare the results of the benchmark for the 2CPU benchmark run with/without cgroup quotas. It is important to note that the Docker run parameter `--cpus` will not give you a dedicated CPU. It´s just a convenience function for `--cpu-quota`: 

```
// CPU hardcap limit (in usecs). Allowed cpu time in a given period.
```

Test results on a 4CPU Ubuntu 16.04 VM with Docker v17.x-CE (container image: `1.9.2-alpine3.6`) with container cpu quota set to "2":

```
benchmark                 NO_LIMITS         2 CPU LIMIT         
                          ns/op             new ns/op     
BenchmarkFindPrimes       230144812         234288931    
BenchmarkFindPrimes-2     131764153         130698084    
BenchmarkFindPrimes-4     85272696        ->150131093 (vs 130698084)   
BenchmarkFindPrimes-8     93632039        ->161014741 (vs 130698084)   
```

It´s obvious that BenchmarkFindPrimes-4/8 run slower compared to NO_LIMITS.  

But do note the **increase of ns/op** for the "2 CPU LIMIT" test between the runs **BenchmarkFindPrimes-2** and **4/8** respectively. In the first, GOMAXPROCS is aligned to the container cpu quota (`GOMAXPROCS=2`), in the latter `GOMAXPROCS > cgroup cpu quota`.

The Go runtime, like many other languages, do not yet have full cgroup support. Hinting cgroup quotas (limits) to the Go runtime not only leads to predictable performance (SLOs), but also prevent unexpected runtime behavior like long [GC pauses](https://github.com/golang/go/issues/19378).

