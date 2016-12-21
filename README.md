Redia Operation
===

基于redigo pub/sub


Installation
----

<pre><code>
go get github.com/johnhjwsosd/redis-publish/subclient
</code></pre>



Example
----------
<pre><code>
	redis := subclient.NewRedis("127.0.0.1:6379", "123")
	pool := redis.NewPool()
	fmt.Println("...Lintening ...")
	redis.SendMsg(pool, "sub1", "test")
        go redis.Listen(pool, "sub1")
</code></pre>