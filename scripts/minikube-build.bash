#!/usr/bin/env bash
echo "Validating running minikube..."
IP="$(minikube ip)"
if  [[ -z  $IP ]]
then
    echo "Minikube is not running, please use 'minikube start' before running this script..."
fi

echo "Validated..."
echo "Building Health Record Repository for Minikube..."
echo "Done!"