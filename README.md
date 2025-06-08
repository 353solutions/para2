# Go Workshop

Miki Tebeka
📬 [miki@353solutions.com.com](mailto:miki@353solutions.com), 𝕏 [@tebeka](https://twitter.com/tebeka), 👨 [mikitebeka](https://www.linkedin.com/in/mikitebeka/), ✒️[blog](https://www.ardanlabs.com/blog/)

#### Shameless Plugs

- [LinkedIn Learning Classes](https://www.linkedin.com/learning/instructors/miki-tebeka)
- [Books](https://pragprog.com/search/?q=miki+tebeka)

[Syllabus](_extra/syllabus.pdf)



---

## Day 1: RPC

### Agenda

- Advanced JSON
    - Custom serialization
    - Missing vs empty values
    - Streaming JSON
- HTTP clients
    - Request body
    - Streaming
    - Authentication
- HTTP servers
    - Dependency injection
    - Writing middleware
    - Streaming responses
    - Routing


### Code



- [value.go](value/value.go) - Custom serialization
- [vm.go](vm/vm.go) - Missing vs empty
- [logs.go](logs/logs.go) - Steaming JSON
- [client.go](events/client.go) - HTTP Clients

### Links

- [Installing Multiple Versions of Go](https://go.dev/doc/manage-install)
    - You can set `GOROOT` to `~/sdk/go1.24.3`, see GoLand instructions [here](https://www.jetbrains.com/help/go/create-a-project-with-go-modules-integration.html)
- [Year 2038 Problem](https://en.wikipedia.org/wiki/Year_2038_problem)
- [Method Sets](https://www.youtube.com/watch?v=Z5cvLOrWlLM)
- [mapstructure](https://pkg.go.dev/github.com/mitchellh/mapstructure#example-Decode) - `map[string]any` -> struct
- [Server Side Events in Go](https://www.freecodecamp.org/news/how-to-implement-server-sent-events-in-go/)
- [Chunked Transfer Encoding](https://en.wikipedia.org/wiki/Chunked_transfer_encoding)
- [JSON Lines](https://jsonlines.org/)
- [HTTP Status Cats](https://http.cat/)
- [Fixing For Loops in Go 1.22](https://go.dev/blog/loopvar-preview)
- [JSON - The Fine Print](https://www.ardanlabs.com/blog/2024/10/json-the-fine-print-part-1.html)

### Data & Other

- [Syllabus](_extra/syllabus.pdf)


---

## Day 2: Going Faster

### Agenda

- Benchmarking & profiling
    - tokenizer
- Performance tips & tricks
- Optimizing memory

### Code

TBD

### Links

- [The Architecture of Open Source Applications](https://aosabook.org/en/) - Including a book on performance
- [Plain Text](https://www.youtube.com/watch?v=4mRxIgu9R70) - Fun talk about Unicode
- [Regular Expression Matching Can Be Simple And Fast](https://swtch.com/~rsc/regexp/regexp1.html)
- [Locality of Behaviour](https://htmx.org/essays/locality-of-behaviour/)
- [Rules of Optimization Club](https://wiki.c2.com/?RulesOfOptimizationClub)
- [Computer Latency at Human Scale](https://twitter.com/jordancurve/status/1108475342468120576)
- [So you wanna go fast](https://www.slideshare.net/TylerTreat/so-you-wanna-go-fast-80300458)
- [High Performance Go](https://dave.cheney.net/high-performance-go-workshop/gophercon-2019.html)
- [Miki's Optimization Overview](_extra/optimize.md)
- [A Benchmarking Checklist](https://www.brendangregg.com/blog/2018-06-30/benchmarking-checklist.html)
- [A Guide to the Go Garbage Collector](https://tip.golang.org/doc/gc-guide)
- [hey](https://github.com/rakyll/hey)
- [Garbage Collection In Go : Part I - Semantics](https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html)
- [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat)

# Day 3: Advanced Concurrency

### Agenda

- Channel semantics
- Goroutine pools
- The "sync" & "sync/atomic" packages
- Handling panics in goroutines

### Code

TBD

### Links

- [The race detector](https://go.dev/doc/articles/race_detector)
- [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)
- [Data Race Patterns in Go](https://eng.uber.com/data-race-patterns-in-go/)
- [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
- [Go Concurrency Patterns: Context](https://go.dev/blog/context)
- [Curious Channels](https://dave.cheney.net/2013/04/30/curious-channels)
- [The Behavior of Channels](https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html)
- [Channel Semantics](https://www.353solutions.com/channel-semantics)
- [Why are there nil channels in Go?](https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308)
- [Amdahl's Law](https://en.wikipedia.org/wiki/Amdahl%27s_law) - Limits of concurrency
- [Computer Latency at Human Scale](https://twitter.com/jordancurve/status/1108475342468120576/photo/1)
- [Concurrency is not Parallelism](https://www.youtube.com/watch?v=cN_DpYBzKso) by Rob Pike
- [Scheduling in Go](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html) by Bill Kennedy
- [conc: better structured concurrency for go](https://github.com/sourcegraph/conc)


### Data & More

- [rtb.go](_extra/rtb.go)
- [taxi.tar](https://storage.googleapis.com/353solutions/c/data/taxi.tar)
