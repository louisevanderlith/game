FROM scratch

COPY cmd/cmd .

EXPOSE 8100

ENTRYPOINT [ "./cmd" ]