FROM golang:1.15.3 AS builder

ARG SERVICE

RUN apt-get -qq update && apt-get -yqq install upx

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /src
COPY . .

RUN go mod init event

RUN go build \
  -trimpath \
  -ldflags "-s -w -extldflags '-static'" \
  -installsuffix cgo \
  -tags netgo \
  -o /bin/service \
  ./${SERVICE}

# strip program removes information from executable binary programs
# and object files that is not essential or required for normal
# and correct execution, thus potentially resulting in better
# performance and sometimes significantly less disk space usage.
RUN strip /bin/service
# UPX is a portable, extendable, high-performance executable
# packer for several different executable formats. It achieves
# an excellent compression ratio and offers *very* fast decompression.
# Your executables suffer no memory overhead or other drawbacks for most
# of the formats supported, because of in-place decompression.
RUN upx -q -9 /bin/service

# The scratch image is the most minimal image in Docker.
# This is the base ancestor for all other images.
# The scratch image is actually empty.
# It doesn't contain any folders/files.
FROM scratch
COPY --from=builder /bin/service /bin/service

ENTRYPOINT ["/bin/service"]