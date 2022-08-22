FROM opensuse/tumbleweed AS build-stage
WORKDIR /src
RUN zypper clean && zypper ref && zypper --non-interactive install --no-recommends mandoc go make git
RUN mkdir -p rpm2docserv
COPY . rpm2docserv/
RUN cd rpm2docserv && make
RUN mkdir -p /srv/docserv && rpm2docserv/bin/rpm2docserv -serving_dir=/srv/docserv

FROM registry.opensuse.org/opensuse/nginx:latest
LABEL maintainer="Thorsten Kukuk <kukuk@thkukuk.de>"

ARG BUILDTIME=
ARG VERSION=unreleased
LABEL org.opencontainers.image.title="Documentation Server"
LABEL org.opencontainers.image.description="Manual pages and documentation to browse with a web server."
LABEL org.opencontainers.image.created=$BUILDTIME
LABEL org.opencontainers.image.version=$VERSION

COPY --from=build-stage /var/cache/rpm2docserv/repo /var/tmp/RPMs
COPY --from=build-stage /srv/docserv /srv/docserv
COPY --from=build-stage /src/rpm2docserv/bin/* /usr/local/bin
COPY example/nginx.conf /usr/local/nginx/etc/
