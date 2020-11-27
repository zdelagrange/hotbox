FROM node:lts-alpine AS build

WORKDIR /web
COPY ./web /web
RUN yarn install && yarn build

FROM nginx:1.15.2-alpine AS nginx
COPY --from=build /web/build /var/www
COPY ./configs/nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
ENTRYPOINT ["nginx","-g","daemon off;"]
