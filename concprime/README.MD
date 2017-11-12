# Benchmark Cgroups Quota (i.e. Limits) and the Effect of GOMAXPROCS.  

## Quick Start
This test requires a working Docker environment with at least 2 cpus for the Docker host to produce meaningful results.  
Compare the output (`time "real"`) of the following runs:

```sh
docker run --rm -it concprime # should be the fastest compared to next runs if your Docker host has >1 cpus
docker run --cpus 1 --rm -it concprime # will be slower, obviously
docker run --cpus 1 --rm -e GOMAXPROCS=1 -it concprime # should be faster than previous run (GOMAXPROCS == --cpus), because of runtime hinting
```

If you have more cpus for your Docker host, play with the `--cpus` and `GOMAXPROCS` flags. It will suprise you :)  


## Run via go Benchmark Tool 
To simply run benchmark via `go test -bench` (defaults to finding the first 1000 primes):

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

## Build from Source (via Docker)

```bash
git clone https://github.com/embano1/gotutorials/
cd gotutorials/concprime
export COMMIT=$(git rev-parse --short --always HEAD)
#change Docker tag ("concprime:$COMMIT") as needed
docker build -t concprime:$COMMIT --build-arg COMMIT=$COMMIT . 
```
