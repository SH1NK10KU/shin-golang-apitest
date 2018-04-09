Feature: Interface Template
In order to test interfaces
As a tester
I want to send HTTP request to verify the response.

  Scenario Outline: 百度搜索测试
    Given <api>接口：发送数据“<data>”
    When 发送请求
    Then 响应参数“<param>”应为：<value>

    Examples:
      | api | data                 | param    | value |
      | 搜索  | {"wd":"Shin FENG"}   | response | 200   |
      | 搜索  | {"wd":"Shin GoTest"} | response | 200   |