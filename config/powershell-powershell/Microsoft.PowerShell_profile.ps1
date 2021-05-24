# this file goes in C:\Users\Steven\Documents\PowerShell
# make sure to unblock file if downloading
$env:EDITOR = 'gvim'

$env:LESS = -join @(
   # Quit if entire file fits on first screen.
   'F'
   # Output "raw" control characters.
   'R'
   # Don't use termcap init/deinit strings.
   'X'
   # Ignore case in searches that do not contain uppercase.
   'i'
)

$env:PATH = @(
   'C:\Users\Steven\go\bin'
   'C:\crystal\bin'
   'C:\dart-sdk\bin'
   'C:\go\bin'
   'C:\ldc2\bin'
   'C:\nim\bin'
   'C:\php'
   'C:\python'
   'C:\python\Scripts'
   'C:\rubyinstaller\bin'
   'C:\sienna'
   'C:\sienna\msys2\mingw64\bin'
   'C:\sienna\msys2\usr\bin'
   'C:\sienna\rust\bin'
   'C:\sienna\vim'
) -join ';'

$env:RIPGREP_CONFIG_PATH = $env:USERPROFILE + '\_ripgrep'
$env:UMBER = 'D:\Git\umber\umber.json'
$env:WINTER = 'D:\Music\Backblaze\winter.db'

Set-PSReadLineKeyHandler Ctrl+UpArrow {
   Set-Location ..
   [Microsoft.PowerShell.PSConsoleReadLine]::InvokePrompt()
}

# color output
[Console]::OutputEncoding = [System.Text.UTF8Encoding]::new()
