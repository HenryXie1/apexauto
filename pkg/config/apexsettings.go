package config

var (
	ApexExample = `
	# 
	# list apex version details in the target
	kubectl-apex list -d dbhost -p 1521 -s testpdbsvc -w syspassword 
	# create apex framework in target db. 
	kubectl-apex create -d dbhost -p 1521 -s testpdbsvc -w syspassword 
	# delete apex framework in target db. 
	kubectl-apex delete -d dbhost -p 1521 -s testpdbsvc -w syspassword 
	`
)

