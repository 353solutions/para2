# Go Workshop

Miki Tebeka
📬 [miki@353solutions.com.com](mailto:miki@353solutions.com), 𝕏 [@tebeka](https://twitter.com/tebeka), 👨 [mikitebeka](https://www.linkedin.com/in/mikitebeka/), ✒️[blog](https://www.ardanlabs.com/blog/)

#### Shameless Plugs

- [LinkedIn Learning Classes](https://www.linkedin.com/learning/instructors/miki-tebeka)
- [Books](https://pragprog.com/search/?q=miki+tebeka)

[Syllabus](_extra/syllabus.pdf)



---

## Session 1: RPC

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


- [value.go](session_1/value/value.go) - Custom serialization
- [vm.go](session_1/vm/vm.go) - Missing vs empty
- [logs.go](session_1/logs/logs.go) - Steaming JSON
- [client.go](session_1/events/client.go) - HTTP Clients

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

## Session 2: Going Faster

### Agenda

- Benchmarking & profiling
    - tokenizer
- Performance tips & tricks
- Optimizing memory

### Code

- [tokenizer](session_2/tokenizer/) - Benchmark, profile, CPU & memory
- [store](session_2/store) - cache & serialization
- [matrix](session_2/matrix/) - CPU cache friendly
- [playground](playground/) - This & that ☺

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


- [counter.go](session_3/counter/counter.go) - The race detector, mutex, `sync/atomic`
- [token.go](session_3/token/token.go) - Refresh token, using `RWMutex`
- [token.go](session_3/token_ch/token.go) - Refresh token, using channels
- [pmap.go](session_3/pmap/pmap.go) - Parallel map
- [payment.go](session_3/payment/payment.go) - `sync.Once`
- [pool.go](session_3/pool/pool.go) - Resource pool & `sync.Pool`

- [taxi_check.go](session_3/taxi_check/taxi_check.go) - Convert sequential to concurrent, pool
- [go_chan.go](session_3/go_chan/go_chan.go) - Channel semantics, goroutines
- [worker_pool.go](session_3/worker_pool/worker_pool.go) - Goroutine pool, return channel
- [fan_in.go](session_3/fan_in/fan_in.go) - Fan in pattern
- [panic.go](session_3/panic/panic.go) - Guarding against panics

### Links

- [automaxprocs](https://pkg.go.dev/go.uber.org/automaxprocs@v1.6.0/maxprocs)
- [x/time/rate](https://pkg.go.dev/golang.org/x/time/rate) - Rate limiter
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
- [GopherCon Israel 2025 Videos](https://www.youtube.com/playlist?list=PLRM-8sTy13XvT_S1KJWZz2KOJgOt_vVVH)
- [lru-cache](https://pkg.go.dev/github.com/hashicorp/golang-lru/v2)
- [singleflight](https://pkg.go.dev/golang.org/x/sync/singleflight)
- [huh](https://github.com/charmbracelet/huh) - UI in terminal
    - Uses [bubbletea](https://github.com/charmbracelet/bubbletea)


### Data & More

- [rtb.go](_extra/rtb.go)
- [taxi.tar](https://storage.googleapis.com/353solutions/c/data/taxi.tar)

---

## Session 4: OO Patterns

### Agenda

- Pointer vs value semantics
- Embedding structs
- Interfaces in depth
- The empty interface and type assertions
- Iterators


### Code


- [driver.go](session_4/driver/driver.go) - Why you can't do inheritance
- [game.go](session_4/game/game.go) - Structs, methods & interfaces
- [error.go](session_4/error/error.go) - When an interface is nil
- [sha1.go](session_4/sha1/sha1.go) - Combining io.Reader & io.Writer
- [wc.go](session_4/wc/wc.go) - Implement io.Writer


- [logger.go](session_4/logger/logger.go) - Keeping interface small with reflection
- [site.go](session_4/site/site.go) - Sub routers
- [client_test.go](session_4/client/client_test.go) - Mocking http.RoundTripper
- [empty.go](session_4/client/empty/empty.go) - The empty interface (`any`)
- [stats.go](session_4/stats/stats.go) - Generics
- [iter.go](session_4/iter/iter.go) - Iterators


### Links

- [Feynman Algorithm](https://wiki.c2.com/?FeynmanAlgorithm)
- [Generics can make your Go code slower](https://planetscale.com/blog/generics-can-make-your-go-code-slower)
- [Coroutines for Go](https://research.swtch.com/coro) by Russ Cox
- [sort examples](https://pkg.go.dev/sort/#pkg-examples) - Read and try to understand
- [When to use generics](https://go.dev/blog/when-generics)
- [Generics tutorial](https://go.dev/doc/tutorial/generics)
- [Generic Interfaces](https://go.dev/blog/generic-interfaces)
- [Methods, interfaces & embedded types in Go](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html)
- [Methods & Interfaces](https://go.dev/tour/methods/1) in the Go tour
- [wrapMiddle](https://github.com/ardanlabs/service/blob/master/foundation/web/middleware.go) in Ardan Labs [service](https://github.com/ardanlabs/service)
- [chi](https://github.com/go-chi/chi) - Example of chaining middleware
- [List of file signatures](https://en.wikipedia.org/wiki/List_of_file_signatures)
- [Method Sets](https://www.youtube.com/watch?v=Z5cvLOrWlLM)



---

## Session 5: Project Engineering

### Agenda

- Creating go executables  
  - Injecting version  
  - Embedding assets  
- Configuration & command line parsing  
- Logging & metrics  
- Writing secure Go code

### Code

- [collatz](session_5/collatz) - Debugging goroutines
- [gosay](session_5/gosay/) - Building executables


- [spinner.go](session_5/spinner/spinner.go) - Terminal spinner
- [grep.go](session_5/grep/grep.go) - Terminal and command line
- [logging.go](session_5/logging/logging.go) - Logging
- [log_design.go](session_5/logging/logger_design.go) - Logger design

### Links

- [Using ldflags to Set Version Information for Go Applications](https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications)
- [GoReleaser](https://goreleaser.com/)
    - [GitHub Action](https://github.com/goreleaser/goreleaser-action)
- [svu](https://github.com/caarlos0/svu) - Bump version
- [Using Zig to Compile Cgo](https://github.com/goreleaser/example-zig-cgo)
- Command line & Options
    - [flag](https://pkg.go.dev/flag)
    - [Cobra](https://cobra.dev/) + [Viper](https://github.com/spf13/viper)
    - Ardan Labs [conf](https://pkg.go.dev/github.com/ardanlabs/conf/v3)
- Security
    - [mkcert](https://github.com/FiloSottile/mkcert)
    - [x/crypto/autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert)
    - [Using Let's Encrypt in Go](https://marcofranssen.nl/build-a-go-webserver-on-http-2-using-letsencrypt)
    - [Customizing Binaries with Build Tags](https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags)
- [tar joke](https://xkcd.com/1168/)
- [Reversim Summit](https://summit2025.reversim.com/) - October 27,28
- [Writing Deployable Code](https://medium.com/@rantav/writing-deployable-code-part-one-13ec6dc90adb)
- [The Twelve Factor App](https://12factor.net)
- [TOML format](https://toml.io/en/)
- [Knight Capital Group](https://en.wikipedia.org/wiki/Knight_Capital_Group) - Cost of configuration error
- [The Art of Unix Programming](https://cdn.nakamotoinstitute.org/docs/taoup.pdf) PDF
- [Generating Code](https://go.dev/blog/generate) - `go:generate`
- [embed](https://pkg.go.dev/embed) package
- [Debug a Go Application Running on Kubernetes](https://www.youtube.com/watch?v=YXu2box7z9k)
- [Core Dump Debugging](https://go.dev/wiki/CoreDumpDebugging)
- [Leek & Seek](https://www.youtube.com/watch?v=94wG_LJH86U) - A lot of diagnostic advices and tools
- [Go: Monitor your goroutine](https://medium.com/@hax.artisan/go-monitor-your-goroutine-application-9edbdd6e581b)
- [Diagnostics](https://go.dev/doc/diagnostics) in the Go docs

- [Structured Logging with slog](https://go.dev/blog/slog)
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- Security Linters
    - [gosec](https://github.com/securego/gosec)
    - [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)
    - [gitleaks](https://github.com/gitleaks/gitleaks)
- Certificates
    - [mkcert](https://github.com/FiloSottile/mkcert) for local development
    - [letsencrypt](https://letsencrypt.org/) - Free certificates
- [OPA](https://pkg.go.dev/github.com/open-policy-agent/opa) - Authorization
- Validation
    - [validator](https://pkg.go.dev/github.com/go-playground/validator/v10#pkg-overview)
- Code Structure
    - [Organizing a Go Module](https://go.dev/doc/modules/layout)
    - [ArdanLabs service](https://github.com/ardanlabs/service/tree/master)
- [zap log to screen & file](https://stackoverflow.com/a/54999899/7650)
- [Estimation](https://xkcd.com/612/)
- [ANSI escape code](https://en.wikipedia.org/wiki/ANSI_escape_code) - Colors in your terminal
- [libffi](https://github.com/libffi/libffi) - Scary C code
    - Does [FFI](https://en.wikipedia.org/wiki/Foreign_function_interface)
- [cgo](https://go.dev/blog/cgo)
- Terminals supporting graphics
    - [ghostty](https://ghostty.org/)
    - [kitty](https://sw.kovidgoyal.net/kitty/)


### Data & Other

- [gopher.txt](_extra/gopher.txt)
- [Secure Code Slides](_extra/secure-go.pdf)
- [journal.tar.gz](_extra/journal.tar.gz)


![](https://pixel-73339669570.me-west1.run.app/p/para2/p.png)
