Push-Location
Set-Location C:\Users\Steven\AppData\Local\git
git clean -d -f -x
$env:MSYSTEM = 'MINGW64'
$env:PATH = 'C:\msys2\mingw64\bin;C:\msys2\usr\bin'

make -j 8 `
CFLAGS=-DCURL_STATICLIB `
CURLDIR=C:/Users/Steven/AppData/Local/curl `
CURL_LDFLAGS='-lcurl -lwldap32 -lcrypt32' `
LDFLAGS=-static `
NO_GETTEXT=1 `
NO_ICONV=1 `
NO_OPENSSL=1 `
NO_TCLTK=1 `
USE_LIBPCRE=

Pop-Location

'
C:\Users\Steven\AppData\Local\git\git.exe status

C:\Users\Steven\AppData\Local\git\git.exe --no-pager log




'

