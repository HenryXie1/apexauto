# Automation tool to create Apex 19.1 in K8S

A kubectl plugin that create Apex (oracle application express) 19.1 on oracle database .


### Intro
 Apex is the foundation of many applications .  We often need to provision a apex for test, stage and prod. We would like to automate apex 19.1 deployment on a Oracle DB. This database can be a  DB in Cloud(AWS, Azure, OCI)  , it can be a DB in a VM, it can be DB pod in K8S.  Once we have db hostname, port , sys password , we can deployment a brand new Apex 19.1  env via 1 command.  We can also delete it via 1 command.
sqlpluspod is based on DB instantclient docker images of [oracle github](https://github.com/oracle/docker-images).

### Demo
![Demo!](images/kubectl-apex-create1.gif)

## Installation

Download kubectl via [official guide](https://kubernetes.io/docs/tasks/tools/install-kubectl/) and configure access for your kubernetes cluster. Confirm kubectl get nodes is working

Download binary from [release link](https://github.com/HenryXie1/apexauto/releases/download/v1.0/kubectl-apex)
Save it to /usr/local/bin of linux box (only linux supported as for now), No installation needed, download and run   
### Usage
```
Usage:
  kubectl-apex list|create|delete [-d dbhostname] [-p 1521] [-s dbservice] [-w syspassword] [-x apexpassword] [-r] [flags]

Examples:

        # 
        # list apex version details in the target
        kubectl-apex list -d dbhost -p 1521 -s testpdbsvc -w syspassword 
        # create apex full development apex framework in target db. 
        kubectl-apex create -d dbhost -p 1521 -s testpdbsvc -w syspassword 
        # create apex runtime only framework in target db. 
        kubectl-apex create -d dbhost -p 1521 -s testpdbsvc -w syspassword -r
        # delete apex framework in target db. 
        kubectl-apex delete -d dbhost -p 1521 -s testpdbsvc -w syspassword 


Flags:
  -x, --apexpassword string   apex password for all new apex related DB schemas (default "BFE2GRPF")
  -d, --dbhost string         DB hostname or IP address
  -p, --dbport string         DB port to access (default "1521")
  -h, --help                  help for kubectl-apex
  -r, --runtime               specify to install Apex runtime only ,default is false
  -s, --service string        DB service to access
  -w, --syspassword string    sys password of DB service
```

### Contribution
More than welcome! please don't hesitate to open bugs, questions, pull requests 
