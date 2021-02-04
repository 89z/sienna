. .\2-version.ps1
Push-Location

if (Test-Path curl) {
   Set-Location curl/lib
   git clean -d -f -x
} else {
   git clone --branch $s_curl --depth 1 git://github.com/curl/curl
   Set-Location curl/lib
}

mingw32-make `
-f Makefile.m32 `
-j 5 `
CFG=-winssl

Pop-Location
