FROM python:3.7
RUN wget https://github.com/edenhill/librdkafka/archive/v1.2.0.tar.gz && tar xvzf v1.2.0.tar.gz && cd librdkafka-1.2.0/ && ./configure && make && make install && ldconfig
RUN pip install -i https://pypi.tuna.tsinghua.edu.cn/simple confluent-kafka 

RUN mkdir -p /opt/kafka_client
COPY setting.py kafka_consumer.py hello.py /opt/kafka_client/
RUN chmod +x /opt/kafka_client/kafka_consumer.py
WORKDIR /opt/kafka_client
CMD ["/bin/sh"]
