### http
#### Version 1.4.3
> 1.在跨域白名单中添加 biliapi.net 域名

#### Version 1.4.1
> 1.修正 server 退出时偶发的 panic

#### Version 1.4.0
> 1.将 CORS 作为默认全局中间件

#### Version 1.3.3
> 1.添加 ibilibili.cn 作为安全域名

#### Version 1.3.2
> 1.删除不需要的获取env color的envColor方法

#### Version 1.3.1
> 1.将color区分请求color和env_color

#### Version 1.3.0
> 1.去掉了handle.go
> 2.server2.go改成serve.go，Server2方法改为Serve

#### Version 1.2.2
> 1.支持上报熔断错误到prometheus平台

#### Version 1.2.1
> 1.修复使用了elk默认字段message  

#### Version 1.2.0
> 1.拆封Do,JSON,PB,Raw方法  

#### Version 1.1.0
> 1.添加http VeriryUser方法  

#### Version 1.0.1

> 1. 修复了读取配置时潜在的数据竞争

#### Version 1.0.0
> 1.修复配置了location时，breaker不生效的问题  
> 2.合并RestfulDo到Do中  
> 3.breaker配置只使用最外层的，url和host仅配置timeout  
