#!/bin/ash
# shellcheck disable=SC2187,SC2169,SC2086,SC2145
set -mex

if [ "${1:0:1}" = '-' ]; then
    set -- nginx "$@"
fi

echo "=> Populating conf template"
rm $HTTP_DIR/config.json
envsubst < config.json.tmpl > config.json
echo "=> Installing config file"
cp config.json $HTTP_DIR/

echo "=> Configuration: "
cat $HTTP_DIR/config.json

echo "=> Populating nginx template"
rm /etc/nginx/nginx.conf

ESC=$ envsubst < nginx.conf.tmpl > nginx.conf

echo "=> nginx.conf"
cat nginx.conf

echo "=> Installing config file"
cp nginx.conf /etc/nginx

echo "=> Running: $@"
"$@"
