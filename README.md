# fari-proxy

一个利用HTTP协议伪装SOCKS5数据的自由上网的代理工具, 将传输的数据包裹在HTTP报文, 尽可能降低被探测的风险

## 特点:

* 使用`aes-cfb`加密
* 使用HTTP协议伪装数据包, 后续会支持自定义HTTP报文
* 对本地网络软件而言, 仍然是使用的SOCKS5代理, 与浏览器等软件无缝兼容 

## 使用方法:

* #### 在本地启动`client`
	
		./client -c .client.json
	
* #### 在可以自由上网的服务器启动`server`
	
		./server -c .server.json
* #### 在本地开启SOCKS5的代理, 例如浏览器的SOCKS5插件

## 配置文件

* `.client.json`

		{
  			"remote_addr" : "127.0.0.1:20010",   远程服务器监听地址
  			"listen_addr" : "127.0.0.1:20011",   本地SOCKS5监听地址
  			"password" : "uzon57jd0v869t7w"
		}

* `.server.json`

		{
  			"listen_addr" : "127.0.0.1:20010",   
 			 "password" : "uzon57jd0v869t7w"
		}
		
## TODO

* 优化代码

* 后台运行

* 支持自定义HTTP报文



 