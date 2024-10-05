# accessor_engine


### 1.将 accessor_engine 可执行文件放在:  
###### /usr/local/bin 

### 2.add this line to source code 
###### //go:generate  accessor_engine structName
###### //go:generate go run ../accessor_engine.go structName get|set|anyother
###### notice,结构体名字后面的操作是可选的，set|get|或者留空