FROM arm32v7/nginx:stable

ENV TZ=America/New_York
COPY web/build /var/www
COPY ./configs/nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
ENTRYPOINT ["nginx","-g","daemon off;"]
