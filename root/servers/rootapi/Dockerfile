# Build upon the alpine base image
FROM alpine

# install ca-certificates
RUN apk add --no-cache ca-certificates

# copy the rootapi directory
COPY rootapi /rootapi

# listen on port 443
EXPOSE 443

# execute rootapi when container runs
ENTRYPOINT [ "/rootapi" ]