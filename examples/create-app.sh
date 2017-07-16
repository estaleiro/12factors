#!/bin/bash
APPNAME=$1
MEMCACHE_HOST=$(minikube ssh '/sbin/ip route | grep default |cut -f 3 -d " "')
MINIKUBE_IP=$(minikube ip)

if [ -z "$APPNAME" ]; then
    echo "Inserir o APPNAME"
    exit 1
fi


cat deployment.yaml | sed -e "s/\$MINIKUBE_IP/$MINIKUBE_IP/g" -e "s/\$APPNAME/$APPNAME/g" -e "s/\$MEMCACHE_HOST/$MEMCACHE_HOST/g" | kubectl create -f -
