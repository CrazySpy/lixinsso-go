# 用于上海立信会计金融学院sso登录的模块

## User

### New(username string, password string) *User

初始化一个用户对象，用于保存用户基本模型，即用户名和密码。

## Authorize

Authorize用于在获取sso授权，包括登录sso和利用sso登录（授权）。其中授权sso中应用需要三个参数：*client id*, *responseType*以及*redirect uri*都在Authorize结构体中。这三个参数可以从其链接中获得。而state参数主要为了安全，而sso使用固定的1qw23e是不具备安全性能的，因此我无视了此参数。

例如：
学工系统：https://sso.lixin.edu.cn/authorize.php?client_id=ufsso_hairun_saass&redirect_uri=http://saass.lixin.edu.cn/redirect.aspx&response_type=code&state=success
其中的参数一目了然。

### NewAuthorize(clientId string, responseType int, redirectURI string) *Authorize

初始化一个Authorize对象。其中responseType是Code或者Type的常量，定义在authorize.go中。

### (*Authorize) Authorize(user *User) bool

进行授权操作，user可选。这里的授权操作包括自身sso的登录，因此可能需要用到user。如果无须用户名和密码，则user可以置nil。

### (*Authorize)AuthorizeAPP(checkFunc func() bool ) bool

进行sso中应用的授权，不可用于sso自身登录。其中checkFunc可选，用于检验APP是否授权成功的函数。如果无须检验，置nil。
