# OVSDB Mutating webhook
To get it working:
- Bring Kubernetes Up. Make sure kubectl and helm client work.
- Make sure the nodes in Kubernetes have the label 'openvswitch=enabled'
- Run the run.sh from this project's base directory
- Install OVSDB from the openstack helm chart (helm install openvswitch/). The helm chart is available here: https://github.com/openstack/openstack-helm
