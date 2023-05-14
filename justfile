set dotenv-load
set export


request_username := env_var('BASIC_AUTH_USERNAME')
request_password := env_var('BASIC_AUTH_PASSWORD')
server_port := env_var('SERVER_PORT')

req path method='GET' *moreArgs='':
    curl -s -u '{{request_username}}:{{request_password}}' -H 'Content-Type: application/json'\
        -X {{method}} 'localhost:{{server_port}}{{path}}' {{moreArgs}} | jq

alias run := start

start:
    @go run ./main.go