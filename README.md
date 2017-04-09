* Detecting new highlights in file 


When user upload a file we 
1) parse it,
2) save data
3) save hash of file 


When he upload another file, we get hash of it file. If it differs
from last upload (or any of uploads?), 

If new file bigger then previous, we try to cut it to previous and
check if hash the same. 



