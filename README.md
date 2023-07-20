## DLP Exact Data Match Generator

Cloudflare recently introduced a fingerprinting based DLP solution. 

One can provide a list of data of their choosing, such as a list of names, addresses – or source code – and that data is
hashed before ever reaching Cloudflare. They store the hashes and scan your traffic or content for matches of the hashes.
When Cloudflare finds a match, they can log or block it.

This tool aims to help sampling of very scan large source-code repositories by scanning and selecting random lines of code.

### Contribute

## Build

```bash
go build ./... -o bin/edmgen
```

## Test

Note: This will automatically pull the Linux sources in the `test/linux` directory; they are used as fixtures for the tests.

```bash
go test ./...
```

