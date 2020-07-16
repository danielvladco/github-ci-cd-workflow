FROM scratch

ADD ./github-ci-cd-workflow ./serverd

EXPOSE 8080

CMD ["./serverd"]