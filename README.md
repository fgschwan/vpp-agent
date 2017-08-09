``# VPP Agent

[![Build Status](https://travis-ci.org/ligato/vpp-agent.svg?branch=master)](https://travis-ci.org/ligato/vpp-agent)
[![Coverage Status](https://coveralls.io/repos/github/ligato/vpp-agent/badge.svg?branch=master)](https://coveralls.io/github/ligato/vpp-agent?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/ligato/vpp-agent)](https://goreportcard.com/report/github.com/ligato/vpp-agent)
[![GoDoc](https://godoc.org/github.com/ligato/vpp-agent?status.svg)](https://godoc.org/github.com/ligato/vpp-agent)
[![GitHub license](https://img.shields.io/badge/license-Apache%20license%202.0-blue.svg)](https://github.com/ligato/vpp-agent/blob/master/LICENSE)

Please note that the content of this repository is currently **WORK IN PROGRESS**.

The VPP Agent is a Golang implentation of a control/management plane for 
[VPP][1] based cloud-native [Virtual Network Functions][2] (VNFs). 

The VPP Agent is built on top of the CN-Infra cloud-native platform. It
is basically a set of VPP-specific plugins that use the CN-Infra platform
to interact with other services / microservices in the cloud (e.g. a KV 
data store, messaging, log warehouse, etc.). The VPP Agent provides a 
model-driven, high-level API to VPP functionality. Clients that consume
this API may be external (via REST or gRPC API, Etcd or a message bus), 
or other other App/Extension plugins in a larger CN-Infra based 
application/VNF. 

The VNF Agent architecture is shown in the following figure: 

![vpp agent](docs/imgs/vpp_agent.png "VPP Agent & its Plugins on top of cn-infra")

Each (northboud) VPP API - L2, L3, ACL, ... - is implemented by a specific
VNF Agent plugin, which translates northbound API calls/operations into 
(southbound) low level VPP Binary API calls. Northbound APIs are defined 
using [protobufs][3], which allow for the same functionality to be accessible
over multiple transport protocols (HTTP, gRPC, Etcd, ...). Plugins use the 
[GoVPP library][4] to interact with the VPP.
 
The set of plugins in the VPP Agent id as follows:
* [Default VPP Plugins][5] - plugins providing northbound APIs to "default" 
  VPP functionality (i.e. VPP functions available "out-of-the-box"): 
  * [NET Interfaces][6] - network interface configuration (PCI Ethernet, MEMIF, 
    AF_Packet, VXLAN, Loopback...) + BFD
  * [L2][7] - Bridge Domains, L2 cross-connects
  * [L3][8] - IP Routes, VRFs...
  * [ACL][9] - VPP access lists (VPP ACL plugin)
* [GOVPPmux][10] - plugin wrapper arounf GoVPP. Multiplexes plugins' access to
  VPP on a single connection.
* [Linux][11] (VETH) - allows optional configuration of Linux virtual ethernet 
  interfaces
* [CN-Infra datasync][12] - data synchronization after HA events
* [CN-Infra core][13] - lifecycle management of plugins (loading, initialization,
  unloading)

The VPP agent repository also contains tools & support infrastructure:

* [agentctl](cmd/agentctl) - a CLI tool that shows the state of the agents and can configure the agents

## Quickstart
For a quick start with the VPP Agent, you can use pre-build Docker images with the Agent and VPP
on [Dockerhub](https://hub.docker.com/r/ligato/vpp-agent/).

0. Run ETCD and Kafka on your host (e.g. in Docker [using this procedure](docker/dev_vpp_agent/README.md#running-etcd-server-on-local-host)).

1. Run VPP + VPP Agent in a Docker image:
```
docker pull ligato/vpp-agent
docker run -it --name vpp --rm ligato/vpp-agent
```

2. Configure the VPP agent using agentctl:
```
docker exec -it vpp agentctl -h
```

3. Check the configuration (using agentctl or directly using VPP console):
```
docker exec -it vpp agentctl show
docker exec -it vpp vppctl
```

## Next Steps
Read the README for the [Development Docker Image](docker/dev_vpp_agent/README.md) for more details.

GoDoc can be browsed [online](https://godoc.org/github.com/ligato/vpp-agent).

### Deployment:
[![K8s integration](docs/imgs/k8s_deployment_thumb.png "VPP Agent - K8s integration")](docs/Deployment.md)

### Extensibility:
[![VPP Agent Extensibility](docs/imgs/extensibility_thumb.png "VPP Agent - example of extensibility")](https://github.com/ligato/cn-sample-service)

### Design & architecture:
[![VPP agent 10.000 feet](docs/imgs/vpp_agent_10K_feet_thumb.png "VPP Agent - 10.000 feet view on the architecture")](docs/Design.md)


## Contribution:
If you are interested in contributing, please see the [contribution guidelines](CONTRIBUTING.md).

[1]: https://fd.io/
[2]: https://github.com/ligato/cn-infra/blob/master/docs/readmes/cn_virtual_function.md
[3]: https://developers.google.com/protocol-buffers/
[4]: https://wiki.fd.io/view/GoVPP
[5]: plugins/defaultplugins
[6]: plugins/defaultplugins/ifplugin
[7]: plugins/defaultplugins/l2plugin
[8]: plugins/defaultplugins/l3plugin
[9]: plugins/defaultplugins/aclplugin
[10]: plugins/govppmux
[11]: plugins/linuxplugin
[12]: https://github.com/ligato/cn-infra/tree/master/datasync
[13]: https://github.com/ligato/cn-infra/tree/master/core