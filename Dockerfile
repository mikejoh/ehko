FROM scratch
COPY ehko /bin/ehko
ENTRYPOINT ["/bin/ehko"]