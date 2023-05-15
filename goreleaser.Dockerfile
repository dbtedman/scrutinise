FROM scratch

COPY scrutinise ./

ENTRYPOINT ["/scrutinise"]
