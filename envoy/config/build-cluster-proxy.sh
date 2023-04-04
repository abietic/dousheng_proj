#!/usr/bin/env bash

echo yes | redis-cli -a 123456 --cluster create 173.17.0.2:6379 173.17.0.3:6379 173.17.0.4:6379 \
173.17.0.5:6379 173.17.0.6:6379 173.17.0.7:6379 \
--cluster-replicas 1

envoy -c /etc/envoy/envoy.yaml