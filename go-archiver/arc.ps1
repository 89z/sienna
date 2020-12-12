git clone --depth 1 git://github.com/mholt/archiver
Set-Location archiver
git pull origin pull/263/head
Set-Location cmd/arc
go build -v -ldflags -s
