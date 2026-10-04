FROM eclipse-temurin:11-jre

ENV ACTIVEMQ_VERSION="5.19.11"
ENV ACTIVEMQ_HOME="/opt/activemq"

RUN mkdir -p ${ACTIVEMQ_HOME} \
    && wget -O /tmp/activemq.tar.gz \
       https://archive.apache.org/dist/activemq/${ACTIVEMQ_VERSION}/apache-activemq-${ACTIVEMQ_VERSION}-bin.tar.gz \
    && tar -xzf /tmp/activemq.tar.gz --strip-components=1 -C ${ACTIVEMQ_HOME} \
    && rm /tmp/activemq.tar.gz

COPY images/activemq/activemq.xml ${ACTIVEMQ_HOME}/conf/activemq.xml

WORKDIR ${ACTIVEMQ_HOME}

CMD ["./bin/activemq", "console"]