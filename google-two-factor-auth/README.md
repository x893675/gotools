# google-two-factor-auth

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

> google两步验证的golang样例程序

## Table of Contents

- [Background](#Background)
- [Build](#Build)
- [Usage](#Usage)

## Background

TOTP(Time-Based One-Time Password)。该算法由三部分组成：

1. 共享密钥（一系列二进制数据）
2. 基于当前时间的输入
3. 签名函数

参考链接:

* [google-auth-wiki](https://github.com/google/google-authenticator/wiki/Key-Uri-Format)
* [How Google Authenticator Works](https://garbagecollected.org/2014/09/14/how-google-authenticator-works/)

## Build

`go build main.go`

## Usage

`./main "hxdm vjec jjws rb3h wizr 4ifu gftm xboz"`

使用生成的url生成二维码(可以使用[qr-demo](http://go-qrcode.appspot.com/)生成二维码)，使用谷歌的`Authenticator`应用扫描二维码对比结果

