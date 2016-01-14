# ufop-golang

golang语言的`ufop`示例程序。

简易使用说明：
* clone到本地后，进入该repo目录。运行 `tar cvf demo.tar --exclude='src' *`
* 到新版[portal](portal.qiniu.com)页面, 添加`自定义数据处理`后 > 立即创建一个版本 > 上传`demo.tar`
* 等待1-2分钟后,版本构建成功,把该版本切换成生产版本,增加运行实例
* 按照七牛数据处理的使用方式，在bucket文件的url后添加 `?XXX`(XXX为刚创建的自定义数据处理名)，访问。

最终这个demo的文件处理结果是输出`Hello World!`以及文件mime type信息。

更多关于`ufop`的制作使用方式请访问 [ufop_完全操作指南](http://ztest.qiniudn.com/ufop_%E5%AE%8C%E5%85%A8%E6%93%8D%E4%BD%9C%E6%8C%87%E5%8D%97v0.2.pdf)
