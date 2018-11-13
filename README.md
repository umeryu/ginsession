
go-gin-boilerplateを使いセッション追加

https://qiita.com/CST_negi/items/3cbe2bff606a2d3afdfa


- メールアドレスとパスワードをHTMLからPOSTする。
- GolangでそのPOST内容を文字列として抽出して変数にセット。
- その文字列を元に、DB(テンポラリでメモリ情報）にアクセスしてメールアドレスとパスワードのセットを持つユーザ情報を探す。
- もしユーザ情報が引けたら、セッション構造体に値を入れて保存することでCookieにセッション情報を保持させる。

go-gin-boilerplate
===========================
A boilerplate for Go Web Applications using Gin Framework
[http://gin-gonic.github.io/gin/](http://gin-gonic.github.io/gin/)

- To be added
    - Session support
    - Local Authentication with MongoDB
    - Third party login with Facebook, Twitter, Google, etc...
    - Assets bundling, pipelining, fingerprinting
    - CSRF and Security Headers
    - Deployment steps for Heroku and Google App Engine
    - +More...

- Dependencies used until now
    - "github.com/gin-gonic/gin"
    - "html/template"


---


Have a request, suggestion or question?

Drop me an email: vitorsvieira at yahoo.com

You can also find me at [@notvitor](https://twitter.com/notvitor)


---


### License

The MIT License (MIT)

Copyright (c) 2014 Vitor Vieira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
