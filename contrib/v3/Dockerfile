# Start from a base image with Go installed
FROM golang:1.19-alpine AS builder

# Install dependency for container 
SHELL ["/bin/sh", "-ecuxo", "pipefail"]
RUN apk add --no-cache ca-certificates build-base git
RUN apk update && apk add unzip

# Set the Current Working Directory inside the container
WORKDIR /code
ENV OLD_VERSION v3.0.4
ENV ROOT /code

ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.2.6/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.2.6/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a


# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN cp "/lib/libwasmvm_muslc.$(uname -m).a" /lib/libwasmvm_muslc.a

# Download and unzip the old version
RUN wget -c "https://github.com/White-Whale-Defi-Platform/migaloo-chain/archive/refs/tags/${OLD_VERSION}.zip" -O /code/${OLD_VERSION}.zip && \
    unzip /code/${OLD_VERSION}.zip -d /code

# Install the binary file 
RUN OLD_VERSION_SUBSTRING=`echo $OLD_VERSION | cut -c 2-` && \
    cd /code/migaloo-chain-$OLD_VERSION_SUBSTRING && \
    LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true make build && \
    file /code/migaloo-chain-$OLD_VERSION_SUBSTRING/bin/migalood && \
    echo "Ensuring binary is statically linked ..." && \
    (file /code/migaloo-chain-$OLD_VERSION_SUBSTRING/bin/migalood | grep "statically linked") && \
    mkdir -p /code/bin && \
    cp /code/migaloo-chain-$OLD_VERSION_SUBSTRING/bin/migalood /code/bin


# Start a new stage from scratch
FROM alpine:3.16
ENV HOME /.migalood
WORKDIR $HOME


# Copy the migalood binary from the builder stage to the final image
COPY --from=builder /code/bin/migalood /usr/bin/migalood

# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657
#grpc
EXPOSE 9090

# Run the binary program
ENTRYPOINT ["migalood"]