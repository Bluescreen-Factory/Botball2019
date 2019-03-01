//build command: swig -go -cgo -intgosize 32 -includeall -addextern -I/usr/include/wallaby wallaby.i

%module wallaby  //name of the resulting GO package

//Included header files:
%{
#include <wallaby/wallaby_go.h>
%}

%insert(cgo_comment_typedefs) %{
#cgo LDFLAGS: -L/usr/lib -l wallaby 
%}

%include <wallaby_go.h>

