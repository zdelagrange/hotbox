#!/bin/bash
docker save -o ~/hotbox_api.img hotbox_api:latest
docker save -o ~/hotbox_frontend.img hotbox_frontend:latest
scp -i ~/.ssh/id_rsa  ~/hotbox_api.img pi@ratatoskr:/home/pi
scp -i ~/.ssh/id_rsa  ~/hotbox_frontend.img pi@ratatoskr:/home/pi
ssh pi@ratatoskr "sudo docker load -i /home/pi/hotbox_api.img"
ssh pi@ratatoskr "sudo docker load -i /home/pi/hotbox_frontend.img"
