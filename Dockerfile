# Description: This dockerfile is used to build a fakessh server in a container.
#
# Copyright (c) 2025 honeok <honeok@duck.com>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.24-alpine AS builder
WORKDIR /go/src/github.com/honeok/fakessh
COPY . .
ENV CGO_ENABLED=0
RUN go build -v -trimpath -ldflags="-s -w -buildid=" -o /go/bin/fakessh fakessh.go

FROM alpine:latest AS dist
COPY --from=builder /go/bin/fakessh /usr/bin/fakessh
RUN set -ex \
    && apk upgrade \
    && apk add --no-cache tzdata \
    && rm -rf /var/cache/apk/*
ENV TZ=Asia/Shanghai
EXPOSE 22
ENTRYPOINT ["fakessh"]