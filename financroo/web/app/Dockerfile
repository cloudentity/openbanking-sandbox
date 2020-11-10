FROM nginx:1.17.10-alpine

ENV HTTP_DIR=/usr/share/nginx/html/

COPY entrypoint.sh /
COPY config.json.tmpl /
COPY nginx.conf.tmpl /

RUN chmod +x entrypoint.sh && \
    rm -rf $HTTP_DIR/*

COPY build/ $HTTP_DIR

STOPSIGNAL SIGTERM
ENTRYPOINT [ "/entrypoint.sh" ]
CMD ["nginx", "-g", "daemon off;"]
