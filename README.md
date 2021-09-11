# goweb
a web framework for golang


## step 1 静态路由
目标：监听到请求时，根据请求的路径获取对应的处理方法，并进行处理
1. 存储形式：通过`map[string]handlerfunc`存储path到handler function的映射
2. 服务获取：监听到事件，根据path从路由中获取服务
3. 路由注册：`RESTful`风格，将path与对应的handle function注册到路由中


## step 2 上下文封装
目标：对request与response相关操作进行封装
1. 存储格式：struct包含request与response
2. 封装对response的操作：String, JSON