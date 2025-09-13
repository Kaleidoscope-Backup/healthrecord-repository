#!/usr/bin/env bash
echo "Validating running minikube..."
IP="$(minikube ip)"
if  [[ -z  $IP ]]
then
    echo "Minikube is not running, please use 'minikube start' before running this script..."
fi

echo "Validated..."

echo "Starting mongodb..."
kubectl create -f mongo-deployment.yaml,mongo-service.yaml 

echo "Starting healthrecord-repository..."
kubectl create -f healthrecord-repository-deployment.yaml,healthrecord-repository-service.yaml 

echo "Starting ingress..."
kubectl apply -f https://raw.githubusercontent.com/containous/traefik/master/examples/k8s/traefik-ds.yaml
kubectl apply -f https://raw.githubusercontent.com/containous/traefik/master/examples/k8s/ui.yaml
kubectl create -f healthrecord-repository.ingress.yaml
echo "$(minikube ip) traefik-ui.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip) healthrecord-repository.minikube" | sudo tee -a /etc/hosts

echo "Done!"