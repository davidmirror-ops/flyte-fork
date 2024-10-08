(deployment-configuration)=

# Configuration

This section will cover how to configure your Flyte cluster for features like
authentication, monitoring, and notifications.




```{list-table}
:header-rows: 0
:widths: 20 30

* - {ref}`Authenticating in Flyte <deployment-configuration-auth-setup>`
  - Basic OIDC and Authentication Setup
* - {ref}`Migrating Your Authentication Config <deployment-configuration-auth-migration>`
  - Migration guide to move to Admin's own authorization server.
* - {ref}`Understanding Authentication <deployment-configuration-auth-appendix>`
  - Migration guide to move to Admin's own authorization server.
* - {ref}`Configuring task pods with K8s PodTemplates <deployment-configuration-general>`
  - Use Flyte's cluster-resource-controller to control specific Kubernetes resources and administer project/domain-specific CPU/GPU/memory resource quotas.
* - {ref}`Customizing project, domain, and workflow resources with flytectl <deployment-configuration-customizable-resources>`
  - Use the Flyte APIs to create new default configurations to override certain values for specific combinations of user projects, domains and workflows.
* - {ref}`Notifications <deployment-configuration-notifications>`
  - Guide to setting up and configuring notifications.
* - {ref}`External Events <deployment-configuration-cloud-event>`
  - How to set up Flyte to emit events to third-parties.
* - {ref}`Monitoring <deployment-configuration-monitoring>`
  - Guide to setting up and configuring observability.
* - {ref}`Optimizing Performance <deployment-configuration-performance>`
  - Improve the performance of the core Flyte engine.
* - {ref}`Platform Events <deployment-configuration-eventing>`
  - Configure Flyte to to send events to external pub/sub systems.
* - {ref}`Resource Manager <deployment-configuration-resource-manager>`
  - Manage external resource pooling
```

```{toctree}
:maxdepth: 1
:name: Cluster Config
:hidden:

auth_setup
auth_migration
auth_appendix
general
customizable_resources
monitoring
notifications
performance
cloud_event
resource_manager
```
