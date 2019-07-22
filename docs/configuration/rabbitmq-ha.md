---
title: "Rabbitmq HA"
sub_link: "rabbitmq-ha"
meta_title: "High availability of Rabbitmq in Polyaxon - Configuration"
meta_description: "Polyaxon offers a couple of ways to set a high available Rabbitmq."
tags:
    - configuration
    - polyaxon
    - rabbitmq
    - scaling
    - high-availability
    - kubernetes
    - docker-compose
sidebar: "configuration"
---

Polyaxon ships with a default Rabbitmq based on the stable [Helm chart](https://github.com/helm/charts/tree/master/stable/rabbitmq-ha).

You can check the chart values to extend it's configuration.

## External Rabbitmq

If you prefer to have Rabbitmq managed by you or hosted outside of Kubernetes, 
you need to disable the in-cluster Rabbitmq, and provide the information needed to establish a connection to the external one, e.g.:


```yaml
rabbitmq-ha:
  enabled: false

externalServices:
  rabbitmq:
    user: polyaxon
    password: polyaxon
    port: 12345
    host: 35.262.163.88
```

## Disabling Rabbitmq

If you decide not to use Rabbitmq, and use Redis for handling events, please check this section on how to alter the [default broker behaviour](/configuration/broker/).

## Scheduling

If you decided to deploy Rabbitmq in-cluster make sure to set proper [node scheduling](/configuration/custom-node-scheduling/) 
to avoid running high load runs on the same node hosting the Rabbitmq.
