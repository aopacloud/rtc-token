# 自研rtc接入指南

## 如何获取SDK
上Slack找**Haibo 雷公**（haibo@aopacloud.sg）

## 如何获取AppID和AppCert
上Slack找**Roga 湮灭**（roga@aopacloud.sg）

## 如何生成token
必须有AppID和AppCert才能生成token，参考本仓库代码，Go语言，注意加密算法的版本

## rtc网关接入地址
### 测试环境
```
"{\"signaling\":\"wss://rtc-dev-gateway.aopacloud.net:6080/rtc/channel\", \"rtconfig\":\"https://rtc-dev-gateway.aopacloud.net:6080/rtc/get_rtc_config\", \"log\":\"https://rtc-dev-gateway.aopacloud.net:6080/ali/v1/get_upload_url\", \"quic\":\"rtc-dev-gateway.aopacloud.net:6081\", \"report\":\"https://rtc-dev-gateway.aopacloud.net:6080/rtc/rtc_event_report\"}"
```
### 正式环境
国内用户和国外用户使用不同的地址，国内用户走了专线加速
**国内用户**
```
"{\"signaling\":\"wss://rtc-gateway-oversea.aopacloud.net:6080/rtc/channel\", \"rtconfig\":\"https://rtc-gateway-oversea.aopacloud.net:6080/rtc/get_rtc_config\", \"log\":\"https://rtc-gateway-oversea.aopacloud.net:6080/ali/v1/get_upload_url\", \"quic\":\"rtc-gateway-oversea-quic.aopacloud.net:16081\", \"report\":\"https://rtc-gateway-oversea.aopacloud.net:6080/rtc/rtc_event_report\"}"
```
**国外用户**
```
"{\"signaling\":\"wss://rtc-gateway.aopacloud.sg:6080/rtc/channel\", \"rtconfig\":\"https://rtc-gateway.aopacloud.sg:6080/rtc/get_rtc_config\", \"log\":\"https://rtc-gateway.aopacloud.sg:6080/ali/v1/get_upload_url\", \"quic\":\"rtc-gateway-quic.aopacloud.sg:16081\", \"report\":\"https://rtc-gateway.aopacloud.sg:6080/rtc/rtc_event_report\"}"
```

## rtc灰度控制
灰度控制是用来确定用户应该使用自研RTC还是声网RTC。接口如下
```
gray/user?app_id=2&room=12345&user=1&room_type=business&room_factory_type=sing&language=ko&region=ko
```
返回值
```
{"code":0,"data":true,"msg":"查询成功"}
```
其中data为true表示使用自研RTC，false表示使用声网RTC。
### 测试环境
测试环境走公网就行了，地址为
```
https://rtc-dev-gateway.aopacloud.net:6080/rtcapi/dev/v1
```
示例（由于&会被转义，测试时用\u0026代替）
```bash
curl https://rtc-dev-gateway.aopacloud.net:6080/rtcapi/dev/v1/gray/user?app_id=2\u0026room=12345\u0026user=1\u0026room_type=business\u0026room_factory_type=sing\u0026language=ko\u0026region=ko
```
### 正式环境
正式环境的灰度请求是内网请求，需要找**Jason Xu 沙加**（jason.xu@olaparty.sg）打通网络。内网地址为
```
http://rtc-sg-gray.aopacloud.private:6080/rtcapi/dev/v1
```
该域名对应的IP是10.24.137.220。
示例（由于&会被转义，测试时用\u0026代替）
```bash
curl http://rtc-sg-gray.aopacloud.private:6080/rtcapi/dev/v1/gray/user?app_id=2\u0026room=12345\u0026user=1\u0026room_type=business\u0026room_factory_type=sing\u0026language=ko\u0026region=ko
```
