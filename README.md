# About
TyphonToyServer is (surprise surprise) a toy server, and is not intended as
a real-world example, as a bulletproof server, or even as a perfect example of
REST server design - it was built for fun in an evening to better understand 
the internals of Monzo's typhon library

# Usage
1. `go run *.go` can be used to start the server
2. `curl -X POST --data 'my example comment' localhost:8888/comments/12345` can be used to post a comment to the server, which can later be fetched with...
3. `curl -G localhost:8888/comments/12345` to fetch the previously posted comment, which returns the comment in plain text.

