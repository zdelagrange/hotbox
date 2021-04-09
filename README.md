# hotbox
A temperature and humidty recorder for the RaspberryPi and the DHT22 temperature and humidity sensor

## Hardware Assembly Instructions
https://pimylifeup.com/raspberry-pi-humidity-sensor-dht22/

## Deployment
hotbox can be deployed with docker on the raspberry pi. There are two containers included, the front end and the api.


```
docker build -t hotbox:api -f build/package/cmd/Dockerfile .
docker build -t nginx:hotbox -f build/package/web/Dockerfile .
docker run --name=frontend -d --rm -p 80:80 hotbox:nginx
docker run --name api -d --privileged -v ${REPO_ROOT}/hotbox/hotbox.db:/go/hotbox/hotbox.db -p 3000:3000 hotbox:api
```
