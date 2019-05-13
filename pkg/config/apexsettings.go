package config

var (
	ApexExample = `
	# use instantclient:apex19 docker image, apex version in the image is 19.1
	# list apex version details in the target
	kubectl apex list
	# create apex framework in target db. 
	kubectl apex create -a dbhost -p 1521 -s testpdbsvc -w syspassword -x apexpassword
	# delete apex framework in target db. 
	kubectl apex delete -a dbhost -p 1521 -s testpdbsvc -w syspassword -x apexpassword
	`
	Sqlpluspodyml = `
apiVersion: v1
kind: Pod
metadata:
  name: sqlpluspod
spec:
  containers:
  - name: sqlpluspod
    image: instantclient:apex19
`
)

