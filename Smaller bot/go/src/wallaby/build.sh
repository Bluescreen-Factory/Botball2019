cp /usr/include/wallaby/wallaby.h /usr/include/wallaby/wallaby_go.h
sed -i -E 's/(#include\s+"graphics(_\w+)?\.h")/\/\/\1/g' /usr/include/wallaby/wallaby_go.h
swig -go -cgo -intgosize 32 -includeall -addextern -I/usr/include/wallaby wallaby.i