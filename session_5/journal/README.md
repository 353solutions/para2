# Writing Secure Go code

Torq ∴  2024 <br />


Miki Tebeka <i class="far fa-envelope"></i> [miki@353solutions.com](mailto:miki@353solutions.com), <i class="fab fa-twitter"></i> [@tebeka](https://twitter.com/tebeka), <i class="fab fa-linkedin-in"></i> [mikitebeka](https://www.linkedin.com/in/mikitebeka/), <i class="fab fa-blogger-b"></i> [blog](https://www.ardanlabs.com/blog/)  

#### Shameless Plugs

- [Go Essential Training](https://www.linkedin.com/learning/go-essential-training/) - LinkedIn Learning
    - [Rest of classes](https://www.linkedin.com/learning/instructors/miki-tebeka)
- [Go Brain Teasers](https://pragprog.com/titles/d-gobrain/go-brain-teasers/) book
    - [Rest of books](https://pragprog.com/search/?q=miki+tebeka)

# Agenda

- Common security threats (OWASP top 10)
- Avoiding injection
- Secure HTTP requests
- Avoiding sensitive data leak
- Handling secrets
- The security mindset and adding security to your development process


# Links

- [Let's talk about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging) by Dave Cheney
- [Go Security Policy](https://golang.org/security)
- [Awesome security tools](https://github.com/guardrailsio/awesome-golang-security)
- [How our security team handles secrets](https://monzo.com/blog/2019/10/11/how-our-security-team-handle-secrets)
- HTTPS
    - [mkcert](https://github.com/FiloSottile/mkcert)
    - [x/crypto/autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert)
    - [Using Let's Encrypt in Go](https://marcofranssen.nl/build-a-go-webserver-on-http-2-using-letsencrypt)
- [Customizing Binaries with Build Tags](https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags)
- Books
    - [Security with Go](https://www.packtpub.com/product/security-with-go/9781788627917)
    - [Black Hat Go](https://nostarch.com/blackhatgo) book
- Limiting Size
    - [MaxBytesHandler](https://pkg.go.dev/net/http#MaxBytesHandler)
    - [MaxBytesReader](https://pkg.go.dev/net/http#MaxBytesReader)
    - [io.LimitReader](https://pkg.go.dev/io#LimitReader)
- [Search for AWS keys in GitHub](https://sourcegraph.com/search?q=context:global+AWS_SECRET_ACCESS_KEY%3D%5B%27%22%5D.%7B40%7D%5B%27%22%5D&patternType=regexp)
- [Fallacies of distributed computing](https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing#The_fallacies)
- [cue](https://cuelang.org/) - Language for data validation
- Serialization Vulnerabilities
    - [XML Billion Laughs](https://en.wikipedia.org/wiki/Billion_laughs_attack) attack
    - [Java Parse Float](https://www.exploringbinary.com/java-hangs-when-converting-2-2250738585072012e-308/)
- [Understanding HTML templates in Go](https://blog.lu4p.xyz/posts/golang-template-turbo/)
- SQL
    - [database/sql](https://pkg.go.dev/database/sql/)
    - [pgx](
    - [sqlx](https://github.com/jmoiron/sqlx)
    - [gorm](https://gorm.io/index.html)
- [Resilient net/http servers](https://ieftimov.com/post/make-resilient-golang-net-http-servers-using-timeouts-deadlines-context-cancellation/)
- [Context](https://blog.golang.org/context) on the Go blog
- [Customizing Binaries with Build Tags](https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags)
- [Our Software Depenedcy Problem](https://research.swtch.com/deps)
- [Go's CVE List](https://www.cvedetails.com/vulnerability-list/vendor_id-14185/Golang.html)
- Static tools
    - [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)
    - [golangci-lint](https://golangci-lint.run/)
    - [gosec](https://github.com/securego/gosec)
    - [staticcheck](https://staticcheck.io/)
    - Use [x/tools/analysis](https://pkg.go.dev/golang.org/x/tools/go/analysis) to write your own linter
- The new[embed](https://pkg.go.dev/embed/) package
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [The Security Mindset](https://www.schneier.com/blog/archives/2008/03/the_security_mi_1.html) by Bruce Schneier
- [Effective Go](https://golang.org/doc/effective_go.html) - Read this!

# Data & Other

- [slides](_extra/slides.pdf)
- entries
    - [add-1.json](_extra/add-1.json)
    - [add-2.json](_extra/add-2.json)
    - [add-3.json](_extra/add-3.json)
- `curl -d@_extra/add-1.json http://localhost:8080/new`
