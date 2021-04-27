---
title: "Common Helm Reference"
sub_link: "platform/reference"
meta_title: "Common Helm Reference For Polyaxon CE and EE - Polyaxon References"
meta_description: "Polyaxon Common Helm Reference For Polyaxon CE and EE charts."
visibility: public
status: published
tags:
  - specification
  - polyaxon
  - yaml
  - helm
  - reference
  - kubernetes
sidebar: "setup"
---

This is the common Helm chart reference for both [Enterprise Edition Control Plane](/docs/setup/platform/enterprise-control-plane/) and [Polyaxon Community Edition](/docs/setup/platform/community-edition/).

## deploymentChart

| Parameter                       | Description                                      | Default
| --------------------------------| -----------------------------------------------  | ----------------------------------------------------------
| `deploymentChart`               | The deployment chart to use, default is platform | `platform`


## deploymentVersion

| Parameter                       | Description                                                                                                                                             | Default
| --------------------------------| ------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------
| `deploymentVersion`             | The deployment version to use, this is important if you are using polyaxon-cli to avoid accidentally deploying/upgrading to a version without noticing  | `latest`


## Namespace

| Parameter                       | Description                                                                                          | Default
| --------------------------------| -----------------------------------------------------------------------------------------------------| ----------------------------------------------------------
| `namespace`                     | The namespace that will be used by Polyaxon to create operations and communicate with other services| `polyaxon`


## RBAC

| Parameter                | Description                                        | Default
| ------------------------ | -------------------------------------------------- | ----------------------------------------------------------
| `rbac.enabled`           | Use Kubernetes role-based access control (RBAC)    | `true`


## Gateway

The gateway is the service fronting the traffic to/from Polyaxon. By default it's deployed with `ClusterIP`.

You can use port forwarding to access the API and dashboard on localhost:

By running Polyaxon CLI, you will automatically auto-configure the cli and clients:

```bash
polyaxon port-forward
```

Or using kubectl:

```bash
kubectl port-forward -n polyaxon svc/polyaxon-polyaxon-gateway 8000:80
```

## Ingress and Gateway service

| Parameter                    | Description                                                         | Default
| ---------------------------- | ------------------------------------------------------------------- | ----------------------------------------------------------
| `ingress.enabled`            | Use Kubernetes ingress                                              | `true`
| `ingress.path`               | Kubernetes ingress path                                             | `/`
| `ingress.hostName`           | Kubernetes ingress hostName                                         | ``
| `ingress.tls`                | Use Ingress TLS (Secrets must be manually created in the namespace) | `[]`
| `ingress.annotations`        | Ingress annotations                                                 | `{}`
| `gateway.service.type`       | Gateway Service type                                                | `ClusterIP`
| `gateway.service.annotations`| Gateway Service annotations                                         | `{}`


This chart provides support for an Ingress resource.

If you enable ingress, please set the gateway service type value to:
 * ClusterIP or NodePort
 * NodePort or LoadBalancer on GKE

Note: using TLS requires either:
 - a preconfigured secret with the TLS secrets in it
 - or the user of [cert-manager](https://github.com/helm/charts/tree/master/stable/cert-manager) to auto request certs from let's encrypt and store them in a secret.

It's also possible to use a service like [externalDNS](https://github.com/helm/charts/tree/master/stable/external-dns) to auto create the DNS entry for the polyaxon API service.

### Ingress Annotations

Example annotations:

```yaml
ingress.kubernetes.io/ssl-redirect: "false"
ingress.kubernetes.io/rewrite-target: /
ingress.kubernetes.io/add-base-url: "true"
ingress.kubernetes.io/proxy-connect-timeout: "600"
ingress.kubernetes.io/proxy-read-timeout: "600"
ingress.kubernetes.io/proxy-send-timeout: "600"
ingress.kubernetes.io/send-timeout: "600"
ingress.kubernetes.io/proxy-body-size: 4G
```

## Securing API server with TLS

If you have your own certificate you can make a new secret with the `tls.crt` and the `tls.key`,
then set the secret name in the values file.

### Automating TLS certificate creation and DNS setup

To automate the creation and registration of new domain name you can use the following services:

* [cert-manager](https://github.com/helm/charts/tree/master/stable/cert-manager)
* [externalDNS](https://github.com/helm/charts/tree/master/stable/external-dns) (Route53 / Google CloudDNS)

once installed, you can set the values for `ingress.tls`:

```yaml
ingress:
  enabled: true
  hostName: polyaxon.acme.com
  tls:
    - secretName: polyaxon.acme-tls
      hosts:
        - polyaxon.acme.com
```

TLS can have more than one host.

In order to get the domain registration to work you need to set the value of `api.service.annotations`
to the annotation needed for your domain:
i.e

```yaml
annotations:
  domainName: polyaxon.my.domain.com
```

## SSL

| Parameter | Description                                                             | Default
| ----------| ------------------------------------------------------------------------| ----------------------------------------------------------
| `ssl`     | To set ssl and serve https with Polyaxon deployed with NodePort service | `{}`


NGINX acts as a reverse proxy for the Polyaxon's front-end server, meaning NGINX proxies external HTTP (and HTTPS) requests to the Polyaxon API.

### NGINX ingress

To use Https in Polyaxon on Kubernetes you can set an ingress-nginx for cluster running on Kubernetes.

Polyaxon's helm chart comes with an ingress resource that you can use with an ingress controller where you should use TLS so that all traffic will be served over HTTPS.

 1. Create a TLS secret that contains your TLS certificate and private key.

    ```bash
    kubectl create secret tls polyaxon-tls --key $PATH_TO_KEY --cert $PATH_TO_CERT
    ```


 2. Add the tls configuration to Polyaxon's Ingress values. (**Do not use CluserIP on GKE**)

    ```yaml
    gateway:
      service:
        type: ClusterIP
    ingress:
      enabled: true
      hostName: polyaxon.acme.com
      tls:
      - secretName: polyaxon.acme-tls
        hosts:
          - polyaxon.acme.com
    ```

    For more information visit the [Nginx Ingress Integration](/integrations/nginx/)

### NGINX for NodePort service

To enable ssl for Polyaxon API running with NodePort service on Kubernetes, you need to provide an ssl certificate and SSL certificate key.

you can provide a self-signed certificate or a browser trusted certificate.

 1. Create a secret for your certificate:

    ```bash
    kubectl create -n polyaxon secret generic polyaxon-cert --from-file=/path/to/certs/polyaxon.com.crt --from-file=/path/to/certs/polyaxon.com.key
    ```

 2. Make sure to update your deployment config with reference to the certificate

    ```yaml
    ssl:
      enabled: true
      secretName: 'polyaxon-cert'
    ```
 3. Set the service type to `NodePort` and update the API's service port to 443.

 N.B. By default Polyaxon mounts the SSL certificate and key to `/etc/ssl`, this value can be updated using the `.Values.ssl.path`.

## dns

| Parameter | Description                                                             | Default
| ----------| ------------------------------------------------------------------------| ----------------------------------------------------------
| `dns`     | DNS configuration for cluster running with custom dns setup             | `{}`


Several users deploy Polyaxon on Kubernetes cluster created with [kubespray](https://github.com/kubernetes-sigs/kubespray),
and by default Kubespray creates Kubernetes with CoreDNS as a cluster DNS server.

### Update DNS backend

Although we could provide logic to detect the DNS used in the cluster, this would require cluster wide RBAC that we think it's unnecessary.
The default DNS backend used by Polyaxon is KubeDNS, to set it to a different DNS, you can provide this value in your Polyaxon's deployment config:

```yaml
dns:
  backend: "coredns"
```

### Update the complete DNS prefix

Since the DNS service is generally deployed on `kube-system` namespace, the default DNS prefix is `kube-dns.kube-system` or `coredns.kube-system` if you update the previous option.

You can also provide the complete DNS prefix, and not use the DNS backend options:

```yaml
dns:
  prefix: kube-dns.other-kube-system
```

### Update DNS cluster

The default dns cluster used in Polyaxon to resolve routes is `cluster.local`, you can provide a Custom Cluster DNS, by setting:

```yaml
dns:
  customCluster: "custom.cluster.name"
```


## Time zone

To set a different time zone for application (convenient for the dashboard and admin interface)
you can can provide a [valid time zone value](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)


| Parameter | Description         | Default
| ----------| --------------------| ----------------------------------------------------------
| `timezone`| The timezone to use | `UTC`


## Node and Deployment manipulation

| Parameter                          | Description                                                  | Default
| -----------------------------------| -------------------------------------------------------------| ----------------------------------------------------------
| `nodeSelector`                     | Node selector for core pod assignment                        | `{}`
| `tolerations`                      | Tolerations for core pod assignment                          | `[]`
| `affinity`                         | Affinity for core                                            | Please check the values


Dependent charts can also have values overwritten. Preface values with

 * `postgresql.*`
 * `redis.*`
 * `rabbitmq-ha.*`


### Node Selectors

Polyaxon comes with a couple of node selectors options to assign pods to nodes for polyaxon's core platform

Additionally every dependency in the helm package, exposes a node selector option.

By providing these values, or some of them,
you can constrain the pods belonging to that category to only run on
particular nodes or to prefer to run on particular nodes.

```yaml
nodeSelector:
  ...

postgresql:
  nodeSelector:
    ...

redis:
  master:
    nodeSelector:
      ...
  slave:
    nodeSelector:
      ...

rabbitmq-ha:
  nodeSelector:
    ...
```

### Tolerations

If one or more taints are applied to a node,
and you want to make sure some pods should not deploy on it,
Polyaxon provides tolerations option for the core platform, as well as for all dependencies, e.i. database, broker, expose their own tolerations option.

```yaml
tolerations:
  ...

postgresql:
  tolerations:
    ...

redis:
  master:
    tolerations:
      ...
  slave:
    tolerations:
      ...

rabbitmq-ha:
  tolerations:
    ...
```

### Affinity

It allows you to constrain which nodes your pod is eligible to schedule on, based on the node's labels.
Polyaxon has a default `Affinity` values for its core components to ensure that they deploy on the same node.

Polyaxon's default affinity:

```yaml
affinity:
  podAffinity:
    preferredDuringSchedulingIgnoredDuringExecution:
      - weight: 100
        podAffinityTerm:
          labelSelector:
            matchExpressions:
            - key: type
              operator: In
              values:
              - "polyaxon-core"
          topologyKey: "kubernetes.io/hostname"
```

You can update your config deployment file to set affinity for each dependency:


```yaml
affinity:
  ...

postgresql:
  affinity:
    ...

redis:
  master:
    affinity:
      ...
  slave:
    affinity:
      ...

rabbitmq-ha:
  affinity:
    ...
```

## IPs/Hosts White list

In order to restrict IP addresses and hosts that can communicate with the API

```yaml
allowedHosts:
  - 127.0.0.1
  - 159.203.150.212
  - .mysite.com  # (Will consume every subdomain of mysite.com)
```

## API Host

In order to receive email and notification with a clickable links to the objects on the platform

```yaml
hostName: 159.203.150.212
```
Or

```yaml
hostName: polyaxon.foo.com
```

## UI

Polyaxon UI comes with several flags to:
 * Disable the UI in case the API is used for submitting jobs and programmatic use with other tools.
 * Enable the admin Dashboard.
 * Enable the offline mode, by default some dependencies are loaded from CDN and require access to the internet, if the end users have no access to the internet you can serve these dependencies using the offline mode.

```yaml
ui:
  enabled: true
  offline: false
  adminEnabled: false
```

## Security context


| Parameter                          | Description                                                  | Default
| -----------------------------------| -------------------------------------------------------------| ----------------------------------------------------------
| `securityContext.enabled`          | enable security context                                      | `false`
| `user`                             | security context UID                                         | `2222`
| `group`                            | security context GID                                         | `2222`


Polyaxon runs all containers as root by default, this configuration is often fine for several deployment,
however, in some use cases it can expose a compliance issue for some teams.

Polyaxon provides a simple way to enable a security context for all core components, experiments and jobs.

Default configuration:

```yaml
securityContext:
  enabled: false
  user: 2222
  group: 2222
```

### Enable security context

```yaml
securityContext:
  enabled: true
```

or enable with custom UID/GID other than 2222/2222:

```yaml
securityContext:
  enabled: true
  user: 1111
  group: 1111
```

This will enable a security context to run all containers using a UID/GID == 1111/1111.

> **N.B.** If you are using a host path or a volume for the artifacts store, make sure to allow the UID/GID to access it.

## Resources' limit

By default Polyaxon does not set limits on the resources for the core components it deploys,
in order to enable the resources limits, your config yaml file should include:

```yaml
limitResources: True
```

This will force Polyaxon to set the resources limits on all services if they include the limits subsections.
