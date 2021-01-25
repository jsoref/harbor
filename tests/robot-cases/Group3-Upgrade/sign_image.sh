#!/bin/bash

docker pull $3:$4

IP=$1
PASSPHRASE='Harbor12345'

echo $IP

mkdir -p /etc/docker/certs.d/$IP/
mkdir -p ~/.docker/tls/$IP:4443/

cp /harbor/ca/ca.crt /etc/docker/certs.d/$IP/
cp /harbor/ca/ca.crt ~/.docker/tls/$IP:4443/

export DOCKER_CONTENT_TRUST=1
export DOCKER_CONTENT_TRUST_SERVER=https://$IP:4443

export NOTARY_ROOT_PASSPHRASE=$PASSPHRASE
export NOTARY_TARGETS_PASSPHRASE=$PASSPHRASE
export NOTARY_SNAPSHOT_PASSPHRASE=$PASSPHRASE
export DOCKER_CONTENT_TRUST_ROOT_PASSPHRASE=$PASSPHRASE
export DOCKER_CONTENT_TRUST_REPOSITORY_PASSPHRASE=$PASSPHRASE
export DOCKER_CONTENT_TRUST_OFFLINE_PASSPHRASE=$PASSPHRASE
export DOCKER_CONTENT_TRUST_TAGGING_PASSPHRASE=$PASSPHRASE

docker login -u admin -p Harbor12345 $IP
docker tag $3:$4 $IP/$2/$3:$4
docker push $IP/$2/$3:$4