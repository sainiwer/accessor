# accessor_engine


### 1.将 accessor_engine 可执行文件放在:  
###### /usr/local/bin 

### 2.add this line to source code 
###### //go:generate  accessor_engine structName 
<br>

### 3.or add this line to source code 
<br>

###### //go:generate go run ../accessor_engine.go structName [get | set]
### * the get/set is optional ,you can leave it blank, otherwise it will be the default