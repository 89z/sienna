Push-Location
git clone --depth 1 git://github.com/alecthomas/chroma
git clone --depth 1 git://github.com/gohugoio/hugo
Set-Location hugo
go mod edit -replace github.com/alecthomas/chroma=../chroma
go build -v -ldflags -s
Pop-Location
