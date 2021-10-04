FROM golang:1.17.1-buster AS project_base

ARG executable_name

RUN apt update && \
    apt install -y git

RUN cd /tmp && \
    git clone https://github.com/magefile/mage && \
    cd ./mage && \
    go run bootstrap.go

COPY . /cp-prometheus-exporter

RUN cd /cp-prometheus-exporter && \
    mage build

FROM debian:buster AS project_image

COPY --from=project_base /cp-prometheus-exporter/build/${executable_name} /${executable_name}
COPY --from=project_base /cp-prometheus-exporter/config.toml /config.toml

CMD [ "./${executable_name}" ]
