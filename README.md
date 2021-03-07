<h1 align="center">gofund</h1>
<p align=center>i3bar显示基金行情脚本</p>
<div align=center>
    <img src="https://img.shields.io/badge/Go->=1.14.6-red?logo=go&color=00ADD8&style=flat"/>
    <img src="https://img.shields.io/badge/Ubuntu-18.04.2-red?logo=Ubuntu&color=E95420&style=flat"/>
    <img src="https://img.shields.io/badge/i3bar-i3cat-red?color=1793D1&style=flat"/>
    <img src="https://img.shields.io/badge/i3bar-i3status-red?color=1793D1&style=flat"/>
</div>
<div align=center>
	<img height=300px width=300px align=center src="http://cdn.mjava.top/blog/diugai.com161509264492359.png"/>
</div>
## 简介

本脚本是基于`Go语言`编写一个基金行情查看脚本，输出格式支持`i3status`和`终端输出(开发中)`，如果要集成到`i3status中`，需要用到另一个插件[i3cat](https://github.com/vincent-petithory/i3cat)。

效果展示：

![20210307130219](http://cdn.mjava.top/blog/20210307130219.png)

![20210307130319](http://cdn.mjava.top/blog/20210307130319.png)

![Screenshot from 2021-03-07 14-32-13](http://cdn.mjava.top/blog/Screenshot from 2021-03-07 14-32-13.png)

## 安装

使用`Go`安装，执行命令

```shell
go get github.com/greycodee/gofund
```

或者直接下载压缩包