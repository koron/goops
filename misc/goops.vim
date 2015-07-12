scriptencoding utf-8

command! -nargs=0 -range Goops call <sid>goops(<line1>, <line2>)

function! s:goops(start, end)
  let cmd = printf("%d,%d!goops -l %s", a:start, a:end, &l:ft)
  silent execute cmd
  silent execute "normal! '[=']"
endfunction

" vim:set ts=8 sts=2 sw=2 tw=0 et:
