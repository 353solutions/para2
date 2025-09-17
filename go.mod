module para2

go 1.24.0

tool (
	golang.org/x/perf/cmd/benchstat
	golang.org/x/tools/cmd/stringer
)

// $ go mod init para2

// Installing Go 1.24
// go install golang.org/dl/go1.24.3@latest
// ~/go/bin/go1.24.3 download
// ~/go/bin/go1.24.3
// alias go=~/go/bin/go1.24.3

require (
	go.etcd.io/bbolt v1.4.2
	go.uber.org/zap v1.27.0
	golang.org/x/sync v0.17.0
)

require (
	github.com/aclements/go-moremath v0.0.0-20210112150236-f10218a38794 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/mod v0.28.0 // indirect
	golang.org/x/perf v0.0.0-20250515181355-8f5f3abfb71a // indirect
	golang.org/x/sys v0.36.0 // indirect
	golang.org/x/tools v0.37.0 // indirect
)
