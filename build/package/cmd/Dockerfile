FROM arm32v7/golang AS build

WORKDIR /go/hotbox
COPY ./cmd /go/hotbox/cmd
COPY ./scripts /go/hotbox/scripts
RUN "./scripts/build.sh"
ENTRYPOINT [ "./scripts/start.sh" ]
