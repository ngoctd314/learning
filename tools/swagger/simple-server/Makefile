.PHONY: initswagger
initswagger:
	swagger init spec \
	--title "A Todo list application" \
	--description "From the todo list tutorial on goswagger.io" \
	--version 1.0.0 \
	--scheme http \
	--consumes application/io.goswagger.examples.todo-list.v1+json \
	--produces application/io.goswagger.examples.todo-list.v1+json

.PHONY: validateswagger
validateswagger:
	swagger validate ./swagger.yml

.PHONY: generateswagger
generateswagger:
	swagger generate server -A todo-list -f ./swagger.yml