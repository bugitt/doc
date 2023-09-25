# scs-doc

北航软院云平台文档

启动站点
```bash
hugo server
```

部署
```bash
DOC_HOST=''
hugo && rsync -avz --delete public/ root@$DOC_HOST:/var/www/doc/
```
