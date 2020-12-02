FROM golang AS build

ENV TZ=America/New_York
WORKDIR /go/hotbox
COPY ./cmd /go/hotbox/cmd
COPY ./scripts /go/hotbox/scripts
RUN "./scripts/build.sh"
ENTRYPOINT [ "./scripts/start.sh" ]
