git clone --depth 1 git://github.com/gohugoio/hugo
Set-Location hugo
$s = 'github.com/alecthomas/chroma'
go mod edit -replace ($s + '=' + $s + '@master')
go build -v -ldflags -s
