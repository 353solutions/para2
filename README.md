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

