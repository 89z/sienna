. .\2-version.ps1
Push-Location

if (Test-Path git) {
   Set-Location git
   git clean -d -f -x
} else {
   git clone --branch $s_git --depth 1 git://github.com/git/git
   Set-Location git
}

$env:MSYSTEM = 'MINGW64'
$env:PATH = 'C:\msys2\mingw64\bin;C:\msys2\usr\bin'

mingw32-make -j 8 `
CFLAGS=-DCURL_STATICLIB `
CURLDIR=../curl `
CURL_LDFLAGS='-lcurl -lwldap32 -lcrypt32' `
LDFLAGS='-s -static' `
NO_GETTEXT=1 `
NO_ICONV=1 `
NO_OPENSSL=1 `
NO_TCLTK=1 `
USE_LIBPCRE=

.\git
objdump -p git.exe | Select-String DLL
Get-ChildItem git.exe
Pop-Location
