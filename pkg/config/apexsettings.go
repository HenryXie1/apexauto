package config

var (
	ApexExample = `
	# sqlpluspod uses iad.ocir.io/espsnonprodint/livesqlsandbox/instantclient:apex19 docker image, apex version in the image is 19.1
	# list apex version details in the target
	kubectl-apex list -d dbhost -p 1521 -s testpdbsvc -w syspassword 
	# create apex framework in target db. 
	kubectl-apex create -d dbhost -p 1521 -s testpdbsvc -w syspassword 
	# delete apex framework in target db. 
	kubectl-apex delete -d dbhost -p 1521 -s testpdbsvc -w syspassword 
	`
)

