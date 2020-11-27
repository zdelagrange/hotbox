FROM golang AS build

WORKDIR /go/hotbox
COPY ./cmd /go/hotbox
COPY ./scripts /go/hotbox
RUN "./scripts/build.sh"
ENTRYPOINT [ "./scripts/start.sh" ]
