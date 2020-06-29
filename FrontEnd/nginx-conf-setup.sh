#!/bin/sh

# If prod use https, else use http
if [ $BACKEND_PRODUCTION == "true" ]
then
  mv /etc/nginx/conf.d/nginx-https.conf /etc/nginx/conf.d/nginx.conf
  rm /etc/nginx/conf.d/nginx-http.conf
else
  mv /etc/nginx/conf.d/nginx-http.conf /etc/nginx/conf.d/nginx.conf
  rm /etc/nginx/conf.d/nginx-https.conf
fi