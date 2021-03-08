<h1 align="center">gofund</h1>
<p align=center>i3bar/命令行终端显示基金行情脚本</p>
<div align=center>
	<a href="https://golang.org/"><img src="https://img.shields.io/badge/Go-%3E=1.14.6-red?logo=go&color=00ADD8&style=flat" /></a>
    <a href="https://ubuntu.com/"><img src="https://img.shields.io/badge/Ubuntu-18.04.2-red?logo=Ubuntu&color=E95420&style=flat"/></a>
    <a href="http://vincent-petithory.github.io/i3cat/"><img src="https://img.shields.io/badge/i3bar-i3cat-red?color=1793D1&style=flat"/></a>
    <a href="https://wiki.archlinux.org/index.php/I3_(%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87)"><img src="https://img.shields.io/badge/i3bar-i3status-red?color=1793D1&style=flat"/></a>
    <img src="https://img.shields.io/badge/Terminal-Linux/Win-red?color=4D4D4D&logo=Windows%20Terminal&style=flat"/>
</div>
<div align=center>
	<img height=300px width=300px align=center src="http://cdn.mjava.top/blog/diugai.com161509264492359.png"/>
</div>


## 简介

本脚本是基于`Go语言`编写一个基金行情查看脚本，输出格式支持`i3status`和`终端输出`，如果要集成到`i3status中`，需要用到另一个插件[i3cat](https://github.com/vincent-petithory/i3cat)。

#### `i3bar`效果展示：

![20210307130219](http://cdn.mjava.top/blog/20210307130219.png)

![20210307130319](http://cdn.mjava.top/blog/20210307130319.png)

![screensgot](http://cdn.mjava.top/blog/screensgot.png)

#### 终端输出效果展示：

![Peekgofund](http://cdn.mjava.top/blog/Peekgofund.gif)

## 功能

- [x] i3wm/i3bar支持
- [x] 命令行支持
- [x] 动态颜色显示
- [x] 休市提醒
- [x] 摸鱼神器
- [x] 超级隐蔽
- [ ] ~~股票~~



## 快速开始

使用`Go`安装，执行命令

```shell
go get github.com/greycodee/gofund
```

或者直接下载压缩包