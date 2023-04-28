TCP三次握手详解-深入浅出
https://blog.csdn.net/jdsjlzx/article/details/123831178

2.1 第一次握手
客户端给服务器发送一个SYN段(在 TCP 标头中 SYN 位字段为 1 的 TCP/IP 数据包), 该段中也包含客户端的初始序列号(Sequence number = J)。

SYN是同步的缩写，SYN 段是发送到另一台计算机的 TCP 数据包，请求在它们之间建立连接

2.2 第二次握手
服务器返回客户端 SYN +ACK 段(在 TCP 标头中SYN和ACK位字段都为 1 的 TCP/IP 数据包)， 该段中包含服务器的初始序列号(Sequence number = K)；同时使 Acknowledgment number = J + 1来表示确认已收到客户端的 SYN段(Sequence number = J)。

ACK 是“确认”的缩写。 ACK 数据包是任何确认收到一条消息或一系列数据包的 TCP 数据包

2.3 第三次握手
客户端给服务器响应一个ACK段(在 TCP 标头中 ACK 位字段为 1 的 TCP/IP 数据包), 该段中使 Acknowledgment number = K + 1来表示确认已收到服务器的 SYN段(Sequence number = K)。

2.4 实例观察
2.4.1 tcpdump
使用tcpdump观察如下：因为都是在本机同时运行client和server所以命令为：tcpdump -i lo port 5555, 只能监听回路lo接口，结果如下

这里写图片描述

 如图用红色圈起来的就是3次握手，但是为什么最后一次握手，为什么ack = 1,而不是369535922 呢，
这是因为这里的第三次握手tcpdump显示的是相对的顺序号。但是为了便于观察我们需要把tcpdump的
顺序号变为绝对的顺序号。

命令只需要加-S(大写)便可，即：tcpdump -i lo port 5555 -S
 

加上之后结果就正常了如下图：

这里写图片描述 

从tcpdump的数据，可以明显的看到三次握手的过程是：

第一次握手：client SYN=1, Sequence number=2322326583 —> server
第二次握手：server SYN=1,Sequence number=3573692787; ACK=1, Acknowledgment number=2322326583 + 1 —> client
第三次握手：client ACK=1, Acknowledgment number=3573692787 + 1 -->server
想简单了解一下TCP三次握手的话, 看到这里就可以了.

3.TCP三次握手详细解析过程：
这里写图片描述
3.1 第一次握手
客户在socket() connect()后主动(active open)连接上服务器, 发送SYN ，这时客户端的状态是SYN_SENT
服务器在进行socket(),bind(),listen()后等待客户的连接，收到客户端的 SYN 后，

3.1.1 半连接队列(syn queue)未满
服务器将该连接的状态变为SYN_RCVD, 服务器把连接信息放到半连接队列(syn queue)里面。

3.1.2 半连接队列(syn queue)已满
服务器不会将该连接的状态变为SYN_RCVD，且将该连接丢弃(SYN flood攻击就是利用这个原理，
对于SYN foold攻击，应对方法之一是使syncookies生效，将其值置1即可，路径/proc/sys/net/ipv4/tcp_syncookies，
即使是半连接队列syn queue已经满了，也可以接收正常的非恶意攻击的客户端的请求，
但是这种方法只在无计可施的情况下使用，man tcp里面的解析是这样说的，

这里写图片描述

 但是我不知道为什么Centos6.9默认是置为1，所以这让我很疑惑
这里写图片描述

半连接队列(syn queue)最大值 /proc/sys/net/ipv4/tcp_max_syn_backlog

 这里写图片描述

SYN flood攻击
攻击方的客户端只发送SYN分节给服务器，然后对服务器发回来的SYN+ACK什么也不做，直接忽略掉，
不发送ACK给服务器；这样就可以占据着服务器的半连接队列的资源，导致正常的客户端连接无法连接上服务器。-----[维基百科]

(SYN flood攻击的方式其实也分两种，第一种，攻击方的客户端一直发送SYN，对于服务器回应的SYN+ACK什么也不做，不回应ACK, 第二种，攻击方的客户端发送SYN时，将源IP改为一个虚假的IP, 然后服务器将SYN+ACK发送到虚假的IP, 这样当然永远也得不到ACK的回应。)

3.2 第二次握手
服务器返回SYN+ACK段给到客户端，客户端收到SYN+ACK段后，客户端的状态从SYN_SENT变为ESTABLISHED，
也即是connect()函数的返回。

3.3 第三次握手
全连接队列(accept queue)的最大值 /proc/sys/net/core/somaxconn (默认128)

全连接队列值 = min(backlog, somaxconn)
这里的backlog是listen(int sockfd, int backlog)函数里面的那个参数backlog

3.3.1 全连接队列(accept queue)未满
服务器收到客户端发来的ACK, 服务端该连接的状态从SYN_RCVD变为ESTABLISHED,
然后服务器将该连接从半连接队列(syn queue)里面移除，且将该连接的信息放到全连接队列(accept queue)里面。

3.3.2 全连接队列(accept queue)已满
服务器收到客户端发来的ACK, 不会将该连接的状态从SYN_RCVD变为ESTABLISHED。
当然全连接队列(accept queue)已满时，则根据 tcp_abort_on_overflow 的值来执行相应动作
/proc/sys/net/ipv4/tcp_abort_on_overflow 查看参数值


 tcp_abort_on_overflow = 0

则服务器建立该连接的定时器，

这个定时器是一个服务器的规则是从新发送syn+ack的时间间隔成倍的增加，
比如从新了第二次握手，进行了5次，这五次的时间分别是 1s, 2s,4s,8s,16s,
这种倍数规则叫“二进制指数退让”(binary exponential backoff)

给客户端定时从新发回SYN+ACK即从新进行第二次握手，(如果客户端设定的超时时间比较短就很容易出现异常)


服务器从新进行第二次握手的次数/proc/sys/net/ipv4/tcp_synack_retries

这里写图片描述
 

tcp_abort_on_overflow = 1
关于tcp_abort_on_overflow的解析如下：

这里写图片描述 

 意思应该是，当 tcp_abort_on_overflow 等于1 时,重置连接(一般是发送RST给客户端)，
至于怎么重置连接是系统的事情了。
不过我在查资料的过程发现，阿里中间件团队博客说并不是发送RST， —[阿里中间件团队博客]

这个博客跑的实例观察到的是服务器会忽略client传过来的包，然后client重传，一定次数后client认为异常，然后断开连接。
当然，我们写代码的都知道代码是第一手的注释，实践是检验真理的唯一标准，
最好还是自己以自己实践为准，因为可能你的环境跟别人的不一样。)

查看全连接队列(accept queue)的使用情况

这里写图片描述 

如上图，第二列Recv-Q是，全连接队列接收到达的连接，第三列是Send-Q全连接队列的所能容纳最大值，
如果，Recv-Q 大于 Send-Q 那么大于的那部分，是要溢出的即要被丢弃overflow掉的。


感想：
1.本来想写TCP连接的建立和终止的，没想到要讲清楚TCP连接的建立已经很大的篇幅了，就只讲TCP连接的建立而已。
2.以前看书的时候，没有解决一个问题的来的深刻或者说脉络清晰，这个就是主题阅读的好处吧。
3.以前没有养成一个遇到问题深入解析，解决问题的习惯，今后慢慢养成。

下面的参考1有实例，会比较详细一点，清晰一些。
参考：

http://jm.taobao.org/2017/05/25/525-1/
https://coolshell.cn/articles/11564.html
https://zh.wikipedia.org/wiki/SYN_cookie
https://zh.wikipedia.org/wiki/SYN_flood
https://www.cnblogs.com/menghuanbiao/p/5212131.html








11种状态名词解析​​​​​​​
LISTEN：等待从任何远端TCP 和端口的连接请求。
 
SYN_SENT：发送完一个连接请求后等待一个匹配的连接请求。
 
SYN_RECEIVED：发送连接请求并且接收到匹配的连接请求以后等待连接请求确认。
 
ESTABLISHED：表示一个打开的连接，接收到的数据可以被投递给用户。连接的数据传输阶段的正常状态。
 
FIN_WAIT_1：等待远端TCP 的连接终止请求，或者等待之前发送的连接终止请求的确认。
 
FIN_WAIT_2：等待远端TCP 的连接终止请求。
 
CLOSE_WAIT：等待本地用户的连接终止请求。
 
CLOSING：等待远端TCP 的连接终止请求确认。
 
LAST_ACK：等待先前发送给远端TCP 的连接终止请求的确认（包括它字节的连接终止请求的确认）
 
TIME_WAIT：等待足够的时间过去以确保远端TCP 接收到它的连接终止请求的确认。
TIME_WAIT 两个存在的理由：
          1.可靠的实现tcp全双工连接的终止；
          2.允许老的重复分节在网络中消逝。
 
CLOSED：不在连接状态（这是为方便描述假想的状态，实际不存在）
————————————————
版权声明：本文为CSDN博主「不脱发的程序猿」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/m0_38106923/article/details/108292454