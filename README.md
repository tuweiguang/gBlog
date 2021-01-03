# gBlog
main总体流程：
1. flag参数
2. 日志初始化
3. 配置文件初始化
4. 数据库初始化
5. gin初始化
6. 启动监听gin
7. 监听信号，优雅重启

项目进度：
- 2020.10.31 -- 2020.11.01: 
  1. Go1.14环境安装，MySQL和redis安装
  2. 工程目录初始化
  3. gin简单使用
  4. markdown转html
  
- 2020.11.02 -- 2020.11.08: 
  1. 整体框架构建完成
  2. 简单登陆
  
- 2020.11.09 -- 2020.11.15:
  1. 功能部分
      1. 静态文件加载
      2. 登录增加cookie机制(以后考虑引进jwt)
  
- 2020.11.16 -- 2020.11.22:
  1. 引入jwt认证机制(可选)
  2. 首页功能规划设计
  3. 模板函数定义
  4. 框架优化和功能开发
  
- 2020.11.23 -- 2020.11.29: 
  平时工作比较忙+周末写了2篇blog

- 2020.11.30 -- 2020.12.06
  1. 尝试写login登录界面
  2. session设置domain与浏览器访问时要一致问题查找
  3. session失效时，子frame出现login界面并不是整个页面跳转到login页面
  4. session失效时，子frame出现login界面并不是整个页面跳转到login页面，再次登陆，进入子frame看到的是上次缓存login界面

- 2020.12.07 -- 2020.12.13
  1. 功能部分开发
  2. 登陆接口优化
  3. 尝试使用新的前端框架，还是等以后吧

- 2020.12.14 -- 2020.12.20
  1. 功能开发（文章管理模块）
  2. 前端传文件+form字段（没解决）
  
- 2020.12.21 -- 2020.12.27
  1. 工作忙 + 周末探寻redis cluster相关(进度滞后)
 
- 2020.12.28 -- 2021.01.03
  0. 新年快乐
  1. 解决上次问题：前端传文件+form字段
  2. 添加文章
  3. 添加分类
  
- 2021.01.04 -- 2021.01.10
  1. 开发文章展示模块  
  
- 未来
  1. 了解corba
  
- 改进  
  2. 每请求一次，若session存在，则继续延长。
  3. 再cookie中添加username选项。