FROM arm32v7/nginx:stable
COPY web/build /var/www
COPY ./configs/nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
ENTRYPOINT ["nginx","-g","daemon off;"]
