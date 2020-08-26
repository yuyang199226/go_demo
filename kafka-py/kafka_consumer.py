#!/usr/bin/env python
import ssl
import json
import os
import sys
import threading

from confluent_kafka import Consumer

import setting

#sys.path.append(os.path.join(os.path.dirname(__file__), '../'))

conf = setting.kafka_setting


class MyThread(threading.Thread):
    def __init__(self, brand, topic_type, pool):
        threading.Thread.__init__(self)
        self.brand = brand
        self.topic_type = topic_type
        self.pool = pool

    def run(self):
        print("Starting " + self.name)
        consumer_kafka(self.brand, self.topic_type, self.pool)

    def stop(self):
        sys.exit()


def consumer_kafka(topic_name):
    context = ssl.create_default_context()
    context = ssl.SSLContext(ssl.PROTOCOL_SSLv23)

    context.verify_mode = ssl.CERT_NONE

    context.check_hostname = False
    # context.load_verify_locations("kafka_operator/ca-cert")

    consumer = Consumer({
        'bootstrap.servers': conf['bootstrap_servers'],
        'group.id': conf['consumer_id'],
        'auto.offset.reset': 'earliest'
    })

    consumer.subscribe([topic_name])

    while True:
        msg = consumer.poll(1.0)

        if msg is None:
            continue
        if msg.error():
            print("Consumer error: {}".format(msg.error()))
            continue

        print('Received message: {}'.format(msg.value().decode('utf-8')))

    consumer.close()

    # consumer = KafkaConsumer(bootstrap_servers=conf['bootstrap_servers'],
    #                          group_id=conf['consumer_id'],
    #                          api_version=(0, 10, 2),
    #                          session_timeout_ms=25000,
    #                          # consumer_timeout_ms = 60000,
    #                          max_poll_records=100,
    #                          fetch_max_bytes=1 * 1024 * 1024,
    #                          security_protocol='SASL_SSL',
    #                          sasl_mechanism="PLAIN",
    #                          ssl_context=context,
    #                          sasl_plain_username=conf['sasl_plain_username'],
    #                          sasl_plain_password=conf['sasl_plain_password'])
    #
    # print('consumer start to consuming %s...'%topic_name)
    # consumer.subscribe((topic_name))
    # try:
    #     for message in consumer:
    #         print(message.offset)
    #         print(message.value)
    #
    #         value_dic = json.loads(message.value)
    #         print("---", value_dic)
    #         site_id = value_dic['site_id']
    #         date = value_dic['date']
    #         action = value_dic['action']
    #         dt = '-'.join((date[0:4], date[4:6], date[6:8]))
    # except Exception as e:
    #     print('consumer except', e)


if __name__ == "__main__":
    consumer_kafka("store-business-engine-testAreaEventCrowd")