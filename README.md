## HOW TO SUBSCRIBE TO REDIS EVENTS USING GOLANG

Recently, a project arose the need to keep a goroutine running in the background, checking for changes in a redis db, and updating different configurations after said events.

This repo is the result of the investigation that lead to the solution. The redis go library allows us to subscribe to a keyspace (a specific index on the database) and add logic every time an event takes place on said keyspace.

To test this solution, you must follow the next steps:

1. Have a Redis database running in the background. Mine was in a Docker container.

2. Run this program using the command "go run main.go" in the folder in which you've cloned the repo.

3. Now, the program is up and running, checking your database located on "localhost:6379", and listening to changes on any index inside said database.

4. Open your database to test changes. If yours is running on a Docker container like mine was, you can run "docker exec -it container-name /bin/bash". This will open a terminal inside said container, from which you can run "redis-cli".

5. In said cli, you can run any command that changes the keys from the database, and you will see the message pop up on the console from which you've run the go program.

6. Enjoy.