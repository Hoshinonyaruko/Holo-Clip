<p align="center">
  <a href="https://www.github.com/hoshinonyaruko/holoclip">
    <img src="images/1.png" width="200" height="200" alt="holoclip">
  </a>
</p>

<div align="center">

# holoclip

_✨ 基于golang.design/x/clipboard 的跨设备剪切板同步器 ✨_  

Holo-Clip Clipboard Sync is an innovative utility designed to seamlessly synchronize clipboard contents across different operating systems. Whether you are working on Windows, MacOS, or Linux, this project ensures real-time synchronization of your clipboard data, enhancing your workflow efficiency and ease of data sharing.

## 兼容性
可在各种架构运行

参数 -port 同步剪切板的端口 同步端 被同步端 需一致
参数 -ip   同步端填写被同步端 被同步端填写同步端

## 场景支持

支持图片\文字

支持内网多设备剪切板同步

支持套娃连接,同时同步n台设备剪切板

支持macos开启接力 macos下运行holoclip macos运行rdp rdp内windows运行holoclip

实现macos-win-ios剪切板同步

支持通过ip 本地电脑同步云服务器剪切板

支持同步虚拟机内的剪切板

更多组合方式