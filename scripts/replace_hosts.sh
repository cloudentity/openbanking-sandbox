#!/bin/bash

SEED="data/seed.yaml"
SEED_GENERATED="data/seed.yaml.generated"

export $(cat .env | sed 's/#.*//g' | xargs)

cp $SEED $SEED_GENERATED
sed -i.bak "s/localhost:8443/${ACP_HOST}:8443/g" ${SEED_GENERATED} && rm "${SEED_GENERATED}.bak"
sed -i.bak "s/localhost/${APP_HOST}/g" $SEED_GENERATED && rm "${SEED_GENERATED}.bak"
