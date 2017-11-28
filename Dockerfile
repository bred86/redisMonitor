FROM scratch

WORKDIR /

ADD main /
ADD config/ /config/

CMD ["/main"]
