" vim: syntax=vim

func Cursor()
   if &expandtab
      let cur_n = col('.') - 1
   else
      let cur_n = col('.') * &shiftwidth
   endif
   let ind_n = indent(line('.'))
   return [cur_n, ind_n]
endfunc

func Enter()
   let cur_a = Cursor()
   let min_n = min(cur_a)
   return "\n" . repeat("\t", min_n / &shiftwidth)
endfunc

func Home()
   let [cur_n, ind_n] = Cursor()
   if cur_n > ind_n
      normal ^
   else
      normal 0
   endif
endfunc

func Tab()
   let &expandtab = 0
   let &shiftwidth = 8
   let &softtabstop = 0
endfunc

func URL_Decode()
   sub/%20/ /ge
   sub/%21/!/ge
   sub/%22/"/ge
   sub/%24/$/ge
   sub/%25/%/ge
   sub/%27/'/ge
   sub/%28/(/ge
   sub/%29/)/ge
   sub/%2C/,/ge
   sub/%2F/\//ge
   sub/%3A/:/ge
   sub/%3D/=/ge
   sub/%5B/[/ge
   sub/%5C/\\/ge
   sub/%5D/]/ge
   sub/%5E/^/ge
   sub/%7B/{/ge
   sub/%7D/}/ge
endfunc

" Set color scheme
colorscheme hearth
" disable default auto indent
filetype indent off
" disable default tab stop
filetype plugin off

" Maps

" Insert mode Enter key
imap <CR> <C-R>=Enter()<CR>
" Insert mode smart Home
imap <Home> <C-O>:call Home()<CR>
" Normal mode smart Home
nmap <Home> :call Home()<CR>
" Normal mode undo find highlight
nmap H :nohlsearch<CR>
" Sort selection
xmap S :sort<CR>

" Variables

" Use light scheme instead of dark
let &background = 'light'
" Use clipboard as default register
let &clipboard = 'unnamed'
" Visible right edge
let &colorcolumn = 81
" default width
let &columns = 84
" Change directory for SWP files
let &directory = $TMP
" Handle files with wide characters
let &encoding = 'UTF-8'
" ExpandTab: use spaces, ShiftWidth: visual mode, SoftTabStop: insert mode
let &expandtab = 1
let &shiftwidth = 3
let &softtabstop = 3
" Use LF instead of CRLF
let &fileformats = 'unix'
" Increase font size
let &guifont = 'Consolas:h12'
" Keep scroll, disable others
let &guioptions = 'r'
" Highlight matches
let &hlsearch = 1
" Both are required
let &ignorecase = 1
let &smartcase = 1
" wrap: disable line wrap, linebreak: break on words when wrap is enabled
let &wrap = 0
let &linebreak = 1
" Line numbers
let &number = 1
" fix 'col' value at end of line
let &virtualedit = 'onemore'
