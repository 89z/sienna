# Install Crystal

This works for me:

~~~
curl.exe `
--netrc-file C:\Users\Steven\_netrc `
-L `
-o crystal.zip `
https://api.github.com/repos/crystal-lang/crystal/actions/artifacts/60912181/zip
~~~

Where `_netrc` looks like this:

~~~
default login <USERNAME> password <PERSONAL ACCESS TOKEN>
~~~

## build

~~~
.\crystal.exe build hello.cr

# Error: can't find file 'prelude'
$env:CRYSTAL_PATH = 'D:\Desktop\etc\crystal\src'

# LINK : fatal error LNK1181: cannot open input file 'pcre.lib'
$env:LIB += ';D:\Desktop\etc'
~~~

- https://docs.github.com/en/rest/reference/actions#artifacts
- https://github.com/actions/upload-artifact/issues/51
- https://github.com/actions/upload-artifact/issues/89
- https://stackoverflow.com/questions/67232445
- https://www.youtube.com/watch?v=eAzIAjTBGgU
