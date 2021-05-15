FROM golang
RUN mkdir /gorilla
WORKDIR /gorilla
ADD bin/gorilla gorilla
ARG BUILD_PORT
ENV PORT $BUILD_PORT
EXPOSE $BUILD_PORT
ENTRYPOINT ["./gorilla"]
