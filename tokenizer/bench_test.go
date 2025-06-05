package tokenizer

import "testing"

var text = `
TO THE RED-HEADED LEAGUE: On account of the bequest of the late
Ezekiah Hopkins, of Lebanon, Pennsylvania, U.S.A., there is now another
vacancy open which entitles a member of the League to a salary of £ 4 a
week for purely nominal services. All red-headed men who are sound in
body and mind and above the age of twenty-one years, are eligible.
Apply in person on Monday, at eleven o’clock, to Duncan Ross, at the
offices of the League, 7 Pope’s Court, Fleet Street.
`

func BenchmarkTokenize(b *testing.B) {
	// for i := 0; i < b.N; i++ {  // Go < 1.24
	for b.Loop() { // 1.24+
		tokens := Tokenize(text)
		if len(tokens) != 47 {
			b.Fatal(len(tokens))
		}
	}
}

// Run benchmark
// $ go test -run ^$ -bench . -count 7 | benchstat -
// $ go test -run ^$ -bench . -count 7 | go tool benchstat -
// Run profiler (CPU)
// $ go test -run ^$ -bench . -cpuprofile cpu.pprof
// View profiling
// $ go tool pprof -http :8081 tokenizer.test cpu.pprof
// Profile memory
//  go test -run ^$ -bench -memprofile mem.pprof -benchmem
// View memory profile
// $ go tool pprof -http :8081 tokenizer.test mem.pprof
// Run the execution tracer
// $ go test -run ^$ -bench . -trace trace.out
// View trace
// $ go tool trace trace.out
// (works only in chrome)
