# Log Generator

This repository contains a simple log generator script that allows you to set the log level for logging. The log levels available are:

- INFO
- WARN
- ERROR

You can change the log level using the following command:

```bash
docker exec -i log-generator sh -c 'echo "INFO" > /tmp/logs/log-level'
docker exec -i log-generator sh -c 'echo "WARN" > /tmp/logs/log-level'
docker exec -i log-generator sh -c 'echo "ERROR" > /tmp/logs/log-level'