# config

#### 介绍
Viper 是国外大神 spf13 编写的开源配置解决方案，具有如下特性:

	设置默认值
	可以读取如下格式的配置文件：JSON、TOML、YAML、HCL
	监控配置文件改动，并热加载配置文件
	从环境变量读取配置
	从远程配置中心读取配置（etcd/consul），并监控变动
	从命令行 flag 读取配置
	从缓存中读取配置
	支持直接设置配置项的值

Viper 配置读取顺序：

	viper.Set() 所设置的值
	命令行 flag
	环境变量
	配置文件
	配置中心：etcd/consul
	默认值

Viper 用起来很方便，在初始化配置文件后，读取配置只需要调用 viper.GetString()、viper.GetInt() 和 viper.GetBool() 等函数即可。

Viper 也可以非常方便地读取多个层级的配置，比如这样一个 YAML 格式的配置：

	common:
	  database:
	    name: test
	    host: 127.0.0.1

如果要读取 host 配置，执行 viper.GetString("common.database.host") 即可。

我们 采用 YAML 格式的配置文件，采用 YAML 格式，是因为 YAML 表达的格式更丰富，可读性更强。

#### 初始化配置
例子在 main 函数中增加了 config.Init(*cfg) 调用，用来初始化配置，cfg 变量值从命令行 flag 传入，可以传值，比如 ./examples -c config.yaml，也可以为空，如果为空会默认读取 conf/config.yaml。


#### Viper 高级用法

现在越来越多的程序是运行在 Kubernetes 容器集群中的，在 API 服务器迁移到容器集群时，可以直接通过 Kubernetes 来设置环境变量，然后程序读取设置的环境变量来配置 API 服务器。

例如，通过环境变量来设置 API  端口：

	$ export HLTYAPI_ADDR=:7777

环境变量名格式为 config/config.go 文件中 viper.SetEnvPrefix("HLTYAPI") 所设置的前缀和配置名称大写，二者用 _ 连接，比如 HLTYAPI_RUNMODE。如果配置项是嵌套的，情况可类推，比如

	....
	max_ping_count: 10           # pingServer函数try的次数
	db:                    
	  name: db_apiserver
	
对应的环境变量名为 HLTYAPI_DB_NAME。

#### 支持热更新配置



#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


