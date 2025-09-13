#!/bin/bash
#

VERSION=$2
ALI_REPO="registry-intl.cn-shanghai.aliyuncs.com"
ALI_IMAGE_REPO="registry-intl.cn-shanghai.aliyuncs.com/karte/healthrecord-repository"
DATABASE_HOST="dds-dj1d6504cd8306441820-pub.mongodb.rds.aliyuncs.com:3717,dds-dj1d6504cd8306442164-pub.mongodb.rds.aliyuncs.com:3717/admin?replicaSet=mgset-11069007"
DATABASE_USERNAME="root"
DATABASE_PASSWORD="yOHWMazpHwrsjBlk1"
DATABASE_BASE_NAME="karte-health-records"
NAME="hr"

CI_APPLICATION_REPOSITORY="registry.github.com/Kaleidoscope-Backup/healthrecord-repository/develop"
CI_APPLICATION_TAG=$1
SERVICE_URL=$NAME'.cb1793cd09e3443a0a3b713c5ab9d8b58.cn-beijing.alicontainer.com'
echo $SERVICE_URL

helm repo add karte https://github.com/Kaleidoscope-Backup-public/helm-repository/raw/master/
helm fetch karte/auto-deploy-app --untar
mv auto-deploy-app/ chart
helm dependency update chart/
helm dependency build chart/
helm dependency build chart
helm dependency update chart

helm upgrade --install \
      --wait \
      --set service.enabled="true" \
      --set image.repository=$ALI_IMAGE_REPO \
      --set image.tag=$CI_APPLICATION_TAG \
      --set image.pullPolicy=IfNotPresent \
      --set image.secrets[0].name="alirepo" \
      --set application.track="stable" \
      --set application.database_url="" \
      --set application.database_host=$DATABASE_HOST \
      --set application.database_username=$DATABASE_USERNAME \
      --set application.database_password=$DATABASE_PASSWORD \
      --set application.database_name=$DATABASE_BASE_NAME \
      --set application.local="ALI" \
      --set service.url=$SERVICE_URL \
      --set replicaCount="3" \
      --namespace="default" \
      --version=$VERSION \
      $NAME \
      chart/

#CLEAN
rm -rf chart/