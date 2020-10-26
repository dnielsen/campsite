FROM golang:1.15.3 AS builder

ARG SERVICE

WORKDIR /src
COPY . .

RUN apt-get -qq update && apt-get -yqq install upx

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

RUN go build \
  -trimpath \
  -ldflags "-s -w -extldflags '-static'" \
  -installsuffix cgo \
  -tags netgo \
  -o /bin/service \
  ./${SERVICE}

# We need the images directory so that we can store our uploads there since
# we're not using Amazon S3 or any other cloud storage provider.
# Since we're using `scratch` image, we can't execute `mkdir` there (there's no `/bin/sh`).
RUN mkdir /images

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
# Since we're using `scratch` image, we can't execute `mkdir` here (there's no `/bin/sh`).
COPY --from=builder ./images /images

ENTRYPOINT ["/bin/service"]