%module mywall  //name of the resulting GO package

// Make mylib_wrap.cxx include this header:
%{
#include <wallaby/swallaby.h>
%}

%insert(cgo_comment_typedefs) %{
#cgo LDFLAGS: -L/usr/lib -l wallaby 
%}

//%include <wallaby/wallaby.h>
%include "/usr/include/wallaby/swallaby.h"

