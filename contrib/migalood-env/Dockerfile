FROM ubuntu

RUN apt-get update && \
  apt-get -y upgrade && \
  apt-get -y install curl jq file

ARG entrypoint=entrypoint.sh

VOLUME /migalood
WORKDIR /migalood
EXPOSE 26656 26657
ENTRYPOINT ["/usr/bin/entrypoint.sh"]
CMD ["start", "--log_format", "plain", "--p2p.seeds", ""]
STOPSIGNAL SIGTERM

COPY ${entrypoint} /usr/bin/entrypoint.sh
