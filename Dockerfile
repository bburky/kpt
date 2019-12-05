# Copyright 2019 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.13-stretch
ENV CGO_ENABLED=0

WORKDIR /go/src/
RUN go get sigs.k8s.io/kustomize/kustomize
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -v -o /usr/local/bin/kpt ./

FROM alpine:latest
RUN apk update && apk upgrade && \
        apk add --no-cache git less man diffutils bash openssh && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*
COPY --from=0 /usr/local/bin/kpt /usr/local/bin/kpt
COPY --from=0 /go/bin/kustomize /usr/local/bin/kustomize
CMD ["kpt"]
