
```text
部署目录：/data/project/release/mall-doc
nginx配置：/usr/local/nginx/conf/vhost
参考：http://www.baker-yuan.cn/articles/62
```

```bash
cp baker_blog_view.conf mall-doc.conf

server {
    listen 80;
    server_name mall-doc.baker-yuan.cn;
    location / {
        root /data/project/release/mall-doc;
        index index.html;
    }
}
```
```text
cd /usr/local/nginx/sbin
./nginx -s reload

```
```text
http://mall-doc.baker-yuan.cn/
```

```bash
nvm use v18.16.1 && pnpm docs:build
```

# 1、新增域名
```text

```

# 2、nginx文件


# 3、打包推送