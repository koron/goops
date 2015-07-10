scriptencoding utf-8

command! -nargs=0 -range Goops call <sid>goops(<line1>, <line2>)

function! s:goops(start, end)
  " TODO:
  echo "goops: " . a:start . ", " . a:end
endfunction

" vim:set ts=8 sts=2 sw=2 tw=0 et:
