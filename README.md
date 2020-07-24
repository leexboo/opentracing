# opentracing
## 背景
随着业务的拆分和微服务间的调用关系逐步复杂化，对完整服务调用过程的定位和监控需求孕育而生，分布式追踪可以帮助我们更加直观地定位一次服务的完整流程，以及服务各个调用环节的处理情况
## 相关技术
opentracing https://github.com/opentracing/opentracing-go

jaeger https://github.com/jaegertracing/jaeger

例子中使用了gin框架，微服务之间协议使用的是grpc，使用了docker打包成镜像并且同时部署到了k8s下，由于是处于测试阶段，部署使用的是all-in-one方式，相关的yaml配置可以在yaml文件中看到

## 实现
通过在请求头等注入traceID和spanID等信息发送到下游服务，在下游服务抽取请求头中相关信息，从而实现对两个服务之间的调用关系形成依赖
具体实现可以参考https://github.com/opentracing/opentracing-go
