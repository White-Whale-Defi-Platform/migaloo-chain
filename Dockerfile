FROM golang:1.21-alpine AS go-builder

# this comes from standard alpine nightly file
#  https://github.com/rust-lang/docker-rust-nightly/blob/master/alpine3.12/Dockerfile
# with some changes to support our toolchain, etc
SHELL ["/bin/sh", "-ecuxo", "pipefail"]
# we probably want to default to latest and error
# since this is predominantly for dev use
# hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates build-base git
# NOTE: add these to run with LEDGER_ENABLED=true
# RUN apk add libusb-dev linux-headers

WORKDIR /code
COPY . /code/

# See https://github.com/CosmWasm/wasmvm/releases 
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.5.1/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.5.1/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a

# check sums
RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep b89c242ffe2c867267621a6469f07ab70fc204091809d9c6f482c3fdf9293830
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep c0f4614d0835be78ac8f3d647a70ccd7ed9f48632bc1374db04e4df2245cb467

# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN cp "/lib/libwasmvm_muslc.$(uname -m).a" /lib/libwasmvm_muslc.a

# force it to use static lib (from above) not standard libgo_cosmwasm.so file
# then log output of file /code/bin/migalood
# then ensure static linking
RUN LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true make build \
  && file /code/bin/migalood \
  && echo "Ensuring binary is statically linked ..." \
  && (file /code/bin/migalood | grep "statically linked")

# --------------------------------------------------------
FROM alpine

COPY --from=go-builder /code/bin/migalood /usr/bin/migalood
ENV HOME /.migalood
WORKDIR $HOME

# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657
#grpc
EXPOSE 9090

ENTRYPOINT ["migalood"]
