FROM scratch
COPY ./cord-ovsdb-sidecar /cord-ovsdb-sidecar
ENTRYPOINT ["/cord-ovsdb-sidecar"]
