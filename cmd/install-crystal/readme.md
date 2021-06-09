# Crystal Windows

Installer for Crystal Programming Language

~~~
crystal build hello.cr
~~~

## `CRYSTAL_PATH`

Example value: `C:\crystal\src`

Error if missing: `can't find file 'prelude'`

## `LIB`

Example value: `C:\crystal\bin`

Error if missing: `cannot open input file 'pcre.lib'`

## `PATH`

Example value: `C:\crystal\bin`

Error if missing: `exec: "crystal": executable file not found in %PATH%`

## Links

- https://docs.github.com/en/rest/reference/actions#artifacts
- https://github.com/actions/upload-artifact/issues/51
- https://github.com/crystal-lang/crystal/blob/d7e0f700/src/compiler/crystal/compiler.cr#L353
- https://stackoverflow.com/questions/67232445
- https://www.youtube.com/watch?v=eAzIAjTBGgU
