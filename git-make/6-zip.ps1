. .\2-version.ps1
$s_dist = 'git-' + $s_git
New-Item -ItemType Directory ($s_dist + '\libexec\git-core')
New-Item -ItemType Directory ($s_dist + '\share\git-core\templates')
Copy-Item git\git-remote-https.exe ($s_dist + '\libexec\git-core')
Copy-Item git\git.exe ($s_dist + '\libexec\git-core')
Copy-Item less.exe ($s_dist + '\libexec\git-core')

# name included with verion number on some of them
@"
Git $s_git
$s_curl
$s_less
"@ > ($s_dist + '\readme.md')
