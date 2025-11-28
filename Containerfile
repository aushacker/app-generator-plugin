FROM registry.access.redhat.com/ubi9/ubi-micro

COPY appGenerator .

ENTRYPOINT ["./appGenerator"]
