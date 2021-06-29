FROM golang
RUN mkdir /giraffe
WORKDIR /giraffe
ADD bin/giraffe giraffe
ARG BUILD_PORT
ENV PORT $BUILD_PORT
EXPOSE $BUILD_PORT
ENTRYPOINT ["./giraffe"]
