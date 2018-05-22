#!/bin/bash

kubectl create -f ./cord-ovsdb-injector/deployment/configmap.yaml
sh ./cord-ovsdb-injector/deployment/webhook-create-signed-cert.sh
kubectl create -f ./cord-ovsdb-injector/deployment/deployment.yaml
kubectl create -f ./cord-ovsdb-injector/deployment/service.yaml
cat ./cord-ovsdb-injector/deployment/mutatingwebhook.yaml | ./cord-ovsdb-injector/deployment/webhook-patch-ca-bundle.sh > ./cord-ovsdb-injector/deployment/mutatingwebhook-ca-bundle.yaml
kubectl create -f ./cord-ovsdb-injector/deployment/mutatingwebhook-ca-bundle.yaml
#kubectl label nodes k8s-01 openvswitch=enabled
