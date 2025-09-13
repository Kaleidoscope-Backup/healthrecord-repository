#!/usr/bin/env bash
echo "Validating running minikube..."
IP="$(minikube ip)"
if  [[ -z  $IP ]]
then
    echo "Minikube is not running, please use 'minikube start' before running this script..."
fi

echo "Validated..."

echo "Removing Health Record Repository..."
kubectl delete deployments healthrecord-repository
kubectl delete service healthrecord-repository

echo "Removing Mongo DB..."
kubectl delete deployments mongo
kubectl delete service mongo

echo "Removing Ingress..."
kubectl delete ingress healthrecord-repository

echo "Done!"