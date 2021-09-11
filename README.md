# goweb
a web framework for golang


## step 1 静态路由
目标：监听到请求时，根据请求的路径获取对应的处理方法，并进行处理
1. 存储形式：通过`map[string]handlerfunc`存储path到`handler function`的映射
2. 服务获取：监听到事件，根据`path`从路由中获取服务
3. 路由注册：RESTful风格，将`path`与对应的`handle function`注册到路由中
