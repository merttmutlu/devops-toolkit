#!/bin/sh

# Set the file path
log_file="/tmp/logs/log-level"

# Infinite loop
while true; do
    # Read the current content of the file
    current_content="$(cat "$log_file")"

    # Check if the content has changed
    if [ "$current_content" == "ERROR" ]; then
        echo "Log level is ERROR : $(date)"

    elif [ "$current_content" == "WARN" ]; then
        echo "Log level is WARN : $(date)"

    elif [ "$current_content" == "INFO" ]; then
        echo "Log level is INFO : $(date)"
    fi

    # Sleep for 5 seconds
    sleep 5
done
