goweb+nginx配置网站

配置本地虚拟域名

C:\Windows\System32\drivers\etc

hosts 文件
```apl
127.0.0.1 localhost
127.0.0.1 www.40018988.com
127.0.0.1 www.xiaweix.com
127.0.0.1 www.waomao.com
127.0.0.1 www.gouhuole.com

127.0.0.1 hubula.com
127.0.0.1 www.hubula.com
127.0.0.1 yangyan.hubula.com
127.0.0.1 blog.hubula.com
127.0.0.1 login.hubula.com
127.0.0.1 passport.hubula.com
127.0.0.1 *.hubula.com
127.0.0.1 aa.hubula.com
127.0.0.1 bb.hubula.com

localhost 127.0.0.1
```
Nginx1.15.11\conf

nginx.conf

71行加入include vhosts/*.conf;

```nginx
#user  nobody;
worker_processes 4;

#error_log  logs/error.log;
#error_log  logs/error.log  notice;
#error_log  logs/error.log  info;

#pid        logs/nginx.pid;


events {
     worker_connections 40960;
}

http {
    include       mime.types;
    default_type  application/octet-stream;
								   
    #log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
    #                  '$status $body_bytes_sent "$http_referer" '
    #                  '"$http_user_agent" "$http_x_forwarded_for"';
		
    #access_log  logs/access.log  main;
     sendfile  on;
    #tcp_nopush     on;

    #keepalive_timeout  0;
     keepalive_timeout 65;

    #gzip  on;



    # another virtual host using mix of IP-, name-, and port-based configuration
    #
    #server {
    #    listen       8000;
    #    listen       somename:8080;
    #    server_name  somename  alias  another.alias;

    #    location / {
    #        root   html;
    #        index  index.html index.htm;
    #    }
    #}
	#include vhosts.conf;
    map $time_iso8601 $logdate {
        '~^(?<ymd>\\d{4}-\\d{2}-\\d{2})' $ymd;
        default                       'date-not-found';
    }
	include vhosts/*.conf;
    # HTTPS server
    #
    #server {
    #    listen       443 ssl;
    #    server_name  localhost;

    #    ssl_certificate      cert.pem;
    #    ssl_certificate_key  cert.key;

    #    ssl_session_cache    shared:SSL:1m;
    #    ssl_session_timeout  5m;

    #    ssl_ciphers  HIGH:!aNULL:!MD5;
    #    ssl_prefer_server_ciphers  on;

    #    location / {
    #        root   html;
    #        index  index.html index.htm;
    #    }
    #}

client_max_body_size  50m;
     client_body_buffer_size 60k;
     client_body_timeout 60;
     client_header_buffer_size 64k;
     client_header_timeout 60;
     error_page 400 /error/400.html;
     error_page 403 /error/403.html;
     error_page 404 /error/404.html;
     error_page 500 /error/500.html;
     error_page 501 /error/501.html;
     error_page 502 /error/502.html;
     error_page 503 /error/503.html;
     error_page 504 /error/504.html;
     error_page 505 /error/505.html;
     error_page 506 /error/506.html;
     error_page 507 /error/507.html;
     error_page 509 /error/509.html;
     error_page 510 /error/510.html;
     
     keepalive_requests 100;
     large_client_header_buffers 4 64k;
     reset_timedout_connection on;
     send_timeout 60;
     sendfile_max_chunk 512k;
     server_names_hash_bucket_size 256;
}
     worker_rlimit_nofile 100000;

```


Nginx1.15.11\conf\vhosts

www.hubula.com_80.conf
````nginx
# http:on-www跳转www
server {
  listen        80;
  server_name  hubula.com;
  return 301 http://www.hubula.com$request_uri;
}


#设定负载均衡的服务器列表
upstream www.hubula.com{
  #根据ip策略负载web服务器
  #ip_hash;
  server 127.0.0.1:8080;
  server 127.0.0.1:8081;
}

#泛域名二级域名解析 统一制定页面 例如404
#然后给需要的二级域名进行设置
server {
  listen        80;
  server_name  *.hubula.com;
  #删除url结尾/ 不删除会混乱搜索引擎
  rewrite ^/(.*)/$ /$1 permanent;
  # location ~ ^/(.*)$ {
  #     proxy_pass http://localhost:8060/zs/$1;
  # }

  location / {
    proxy_pass http://www.hubula.com/demo;
  }
  #反向代理的路径下找不到文件，需要单独指定js css文件的访问路径。
  location ~ .*\.(js|css)$ {
    proxy_pass http://www.hubula.com;
  }
}

#www主域名目录结构
server {
  listen        80;
  server_name  www.hubula.com;
  proxy_set_header Host $host;
  root   "F:/phpstudy_pro/WWW/www.hubula.com";
  #删除url结尾/ 不删除会混乱搜索引擎
  rewrite ^/(.*)/$ /$1 permanent;
  location ~^/passport(.*)$ {
    proxy_pass http://www.hubula.com/demo?$1;
  }
  # location /passport/ {
  #     #rewrite ^/passport/(.*)$ http://passport.paodj.com/$1;
  #     proxy_pass http://localhost:8080/errors;
  # }
  #默认请求 根目录
  location / {
    index index.php index.html error/index.html;
    error_page 400 /error/400.html;
    error_page 403 /error/403.html;
    error_page 404 /error/404.html;
    error_page 500 /error/500.html;
    error_page 501 /error/501.html;
    error_page 502 /error/502.html;
    error_page 503 /error/503.html;
    error_page 504 /error/504.html;
    error_page 505 /error/505.html;
    error_page 506 /error/506.html;
    error_page 507 /error/507.html;
    error_page 509 /error/509.html;
    error_page 510 /error/510.html;
    include F:/phpstudy_pro/WWW/www.hubula.com/nginx.htaccess;
    autoindex  off;

    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header Host $host;
    proxy_http_version 1.1;
    #请求转向mysvr 定义的服务器列表
    proxy_pass  http://www.hubula.com/ ;
    # 153 #以下是一些反向代理的配置可删除.
    # 155 proxy_redirect off;
    # 157 #后端的Web服务器可以通过X-Forwarded-For获取用户真实IP
    # 159 proxy_set_header Host $host;
    # 161 proxy_set_header X-Real-IP $remote_addr;
    # 163 proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    # 165 client_max_body_size 10m;    #允许客户端请求的最大单文件字节数
    # 167 client_body_buffer_size 128k;  #缓冲区代理缓冲用户端请求的最大字节数，
    # 169 proxy_connect_timeout 90;  #nginx跟后端服务器连接超时时间(代理连接超时)
    # 171 proxy_send_timeout 90;        #后端服务器数据回传时间(代理发送超时)
    # 173 proxy_read_timeout 90;         #连接成功后，后端服务器响应时间(代理接收超时)
    # 175 proxy_buffer_size 4k;             #设置代理服务器（nginx）保存用户头信息的缓冲区大小
    # 177 proxy_buffers 4 32k;               #proxy_buffers缓冲区，网页平均在32k以下的话，这样设置
    # 179 proxy_busy_buffers_size 64k;    #高负荷下缓冲大小（proxy_buffers*2）
    # 181 proxy_temp_file_write_size 64k;  #设定缓存文件夹大小，大于这个值，将从upstream服务器传
  }
  # 静态文件，nginx自己处理

  #  location ~ ^/(images|javascript|js|css|flash|media)/ {

  #  root /public/;

  #  #过期30天，静态文件不怎么更新，过期可以设大一点，如果频繁更新，则可以设置得小一点。

  #  #expires 30d;

  #  }
}

#passport 用户注册登录二级域名
server {
  listen        80;
  server_name  passport.hubula.com;

  #删除url结尾/ 不删除会混乱搜索引擎
  rewrite ^/(.*)/$ /$1 permanent;
  #absolute_redirect off;取消绝对路径的重定向

  #注册
  # location /member/register {
  #     proxy_pass http://localhost:8080/passport/member/register;
  # }
  # #登录
  # location /member/login {
  #     proxy_pass http://localhost:8080/passport/member/login;
  # }
  # location /member {
  #     proxy_pass http://localhost:8080/passport/member/;
  # }
  location / {
    proxy_pass http://localhost:8080/passport/;
  }

  location ~ .*\.(js|css)$ {
    proxy_pass http://www.hubula.com;
  }
}

# 只允许固定ip访问 并加上密码
server {
  listen        80;
  server_name  gly_waomaomyadmin.hubula.com;
  location / {
    allow 127.0.0.1;
    allow 127.0.0.3;
    deny all;
    # auth_basic "admin_aa";
    # auth_basic_user_file htpasswd;
    proxy_pass http://www.hubula.com/demo/;
  }

  location ~ .*\.(js|css)$ {
    proxy_pass http://www.hubula.com;
  }

}

# 用ip或者其他域名访问时 直接跳转 也可以返回403
server {
  listen        80 default_server;
  server_name _;
  #rewrit ^ http://www.hubula.com$request_uri;
  return 403;
}

````



