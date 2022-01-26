FROM golang:1.16-alpine

RUN apk add openssh git

# Set up ssh things to access private repos.
COPY id_rsa /root/.ssh/
RUN chmod 0600 /root/.ssh/id_rsa
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts
RUN ssh-keyscan bitbucket.org >> /root/.ssh/known_hosts

WORKDIR /go/src/github.com/ciazhar/emobi-service/
EXPOSE 8080

# ENTRYPOINT ./account-service dev
CMD ["/bin/sh"] 
