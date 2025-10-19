FROM nginx:stable-alpine

COPY public /usr/share/nginx/html/doc
COPY nginx.conf /etc/nginx/conf.d/default.conf