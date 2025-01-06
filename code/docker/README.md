Implementation of Docker

Docker uses Linux namespaces to create isolated container enviroments. This isolated container has its own view of the host systems resources like network, processes and file system. It uses a combination of multiple namespaces (PID, Network, Mount, etc.) which makes each container an independent environment.
