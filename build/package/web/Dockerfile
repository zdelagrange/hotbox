
FROM arm32v7/nginx:1.15.2-alpine
COPY web/build /var/www
COPY ./configs/nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
ENTRYPOINT ["nginx","-g","daemon off;"]
