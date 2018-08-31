# Micro Service

## 研究如何搭建微服务架构

- [ ] 所涉及知识

    * 能为业务需要提供更大并发的scale-up 和 scale out能力
    * 微服务是为了最大限度的解耦，不同业务系统甚至可以是不同语言之间的通信
    * 微服务做完一件事情需要涉及多系统调用,如何保证系统的可靠性
    * 微服务拆分不是简单的系统组合，再说一遍一定要理解业务，否则上层服务一定会出现大量的交叉调用

- [ ] 微服务的设计原则

    - 原子服务

        首先确认最基本业务最维度的原子服务,原子服务的定义就是大家都会最大化重用的功能,需要在应用内的闭环操作,没有任何跨其他服务的分支逻辑

    - 服务组合

        在业务场景下,一个功能都需要跨越多个原子服务来完成一个动作,组合服务就是将业务逻辑抽象成独立自主的域,域之间需要保持隔离,服务组合会使用到多个原子服务来完成业务逻辑

    - 业务编排

        最外层就是面向用户的业务流程,一个产品化的商业流程需要对组合服务进行逻辑编排来完成最终的业务结果,这个编排服务可以完全是自动化的,通过工作流引擎进行组合自动化来完成特定的SOP定义

    - 运维管理的复杂度提升

- [ ] 涉及的语言,框架

    - Golang
    - Protobuf
    - github.com/spf13/pflag
    - github.com/spf13/viper
    - google.golang.org/grpc
    - circuitBreaker
    - interceptor
    - discoveryType
    - HealthCheck
    - https://github.com/afex/hystrix-go/
    - http://opentracing.io/documentation/
    - github.com/garyburd/redigo/redis
    - "gopkg.in/mgo.v2"



- [ ] 不同服务之间如何通信
- [ ] 不同项目之间如何通信
- [ ] 如何部署
