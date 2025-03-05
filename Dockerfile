FROM nginx
COPY grpc-app/proxy/nginx.conf /etc/nginx/nginx.conf

EXPOSE 9999