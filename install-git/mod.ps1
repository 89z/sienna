New-Item -Force -ItemType Directory C:\sienna\msys2\usr\ssl\certs
Set-Location C:\sienna\msys2\usr\ssl\certs
Invoke-WebRequest -OutFile ca-bundle.crt https://curl.haxx.se/ca/cacert.pem
