#!/bin/bash

SEED="data/seed.yaml"
SEED_GENERATED="data/seed.yaml.generated"

env_file=${ENV:-.env}

export $(cat ${env_file} | sed 's/#.*//g' | xargs)

cp $SEED $SEED_GENERATED
sed -i.bak "s|https:\/\/localhost:8443|${ACP_URL}|g" ${SEED_GENERATED} && rm "${SEED_GENERATED}.bak"
sed -i.bak "s/localhost/${APP_HOST}/g" $SEED_GENERATED && rm "${SEED_GENERATED}.bak"
