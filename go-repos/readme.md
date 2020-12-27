# Go functions

repo                    | godoc.org | golang.org | go.dev | googlesource.com
------------------------|-----------|------------|--------|-----------------
golang.org/x/arch       |           |            | yes    | yes
golang.org/x/benchmarks |           | yes        | yes    | yes
golang.org/x/blog       | yes       | yes        | yes    | yes
golang.org/x/build      |           | yes        | yes    | yes
golang.org/x/crypto     | yes       | yes        | yes    | yes
golang.org/x/debug      |           | yes        | yes    | yes
golang.org/x/example    |           |            |        | yes
golang.org/x/exp        | yes       | yes        | yes    | yes
golang.org/x/image      | yes       | yes        | yes    | yes
golang.org/x/lint       |           |            |        | yes
golang.org/x/mobile     | yes       | yes        | yes    | yes
golang.org/x/mod        |           |            | yes    | yes
golang.org/x/net        | yes       | yes        | yes    | yes
golang.org/x/oauth2     |           |            | yes    | yes
golang.org/x/perf       |           | yes        | yes    | yes
golang.org/x/pkgsite    |           | yes        |        | yes
golang.org/x/playground |           |            |        | yes
golang.org/x/review     |           | yes        | yes    | yes
golang.org/x/scratch    |           |            |        | yes
golang.org/x/sync       |           | yes        | yes    | yes
golang.org/x/sys        | yes       | yes        | yes    | yes
golang.org/x/talks      | yes       |            | yes    | yes
golang.org/x/term       |           |            | yes    | yes
golang.org/x/text       | yes       | yes        | yes    | yes
golang.org/x/time       |           | yes        | yes    | yes
golang.org/x/tools      | yes       | yes        | yes    | yes
golang.org/x/tour       |           | yes        |        | yes
golang.org/x/vgo        |           |            |        | yes
golang.org/x/website    |           |            | yes    | yes
golang.org/x/xerrors    |           |            |        | yes

- <https://go.googlesource.com>
- <https://godoc.org/-/subrepo>
- <https://golang.org/pkg/#subrepo>
- <https://pkg.go.dev/golang.org/x/build/repos>

`godoc.org` is fine for my needs. Here is where the list is defined:

<https://github.com/golang/gddo/blob/master/gddo-server/assets/templates/subrepo.html>

How do we get the subpackages?

https://api.godoc.org/packages

or this?

results | repo
--------|-------------------------------------------
100     | api.godoc.org/search?q=golang.org/x/tools/
100     | api.godoc.org/search?q=golang.org/x/crypto/
56      | api.godoc.org/search?q=golang.org/x/net/
49      | api.godoc.org/search?q=golang.org/x/text/
42      | api.godoc.org/search?q=golang.org/x/exp/
33      | api.godoc.org/search?q=golang.org/x/image/
30      | api.godoc.org/search?q=golang.org/x/mobile/
14      | api.godoc.org/search?q=golang.org/x/sys/
6       | api.godoc.org/search?q=golang.org/x/blog/
1       | api.godoc.org/search?q=golang.org/x/talks/

`tools` and `crypto` both go over 100, but I dont care about those anyway. So I
can start with these:

~~~
56      | api.godoc.org/search?q=golang.org/x/net/
49      | api.godoc.org/search?q=golang.org/x/text/
42      | api.godoc.org/search?q=golang.org/x/exp/
14      | api.godoc.org/search?q=golang.org/x/sys/
~~~
