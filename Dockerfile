FROM hub.docker.com/tia/alpine:0.0.2

WORKDIR /apps/tia
COPY app .
COPY env.yaml .
ENV PATH="/apps/tia:$PATH"