FROM centos:7

WORKDIR /root

ADD ./main .
EXPOSE 80

CMD ./main
