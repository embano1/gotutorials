# Run
Example:
- Limit container by 200mb of memory (`-m`)
- Limit container by 0.1 CPU (`--cpus`)
- Disable swap (`--memory-swappiness`)
- 5s delay (`-d`) between each memory acquisition (100mb chunks)
- also burn one CPU (-b)

```bash
$ docker run -it -m 200m --cpus 0.1 --memory-swappiness=0 embano1/memhogger -d 5s -b
```

# Usage

```bash
$ docker run embano1/memhogger:latest -h
Usage of memhogger:
  -b    Also burn CPU (default: off)
  -d duration
        Delay between allocations (default 1s)
  -i int
        How many 100MB chunks to allocate (default 5)
```

# Build

Will build a binary and container using Docker multi-stage builds.  

Requires git (to pull) and Docker CLI (build binary and image) installed. Replace Docker image tags with your values.

```bash
$ git clone https://github.com/embano1/gotutorials/
$ cd gotutorials/memhogger
$ COMMIT=$(git rev-parse --short HEAD)
$ docker build -t embano1/memhogger:${COMMIT} .
$ docker tag embano1/memhogger:${COMMIT} embano1/memhogger:latest
```