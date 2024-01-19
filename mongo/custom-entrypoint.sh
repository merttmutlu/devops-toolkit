#!/bin/bash

# Start mongod in the background with official entry-point.sh
docker-entrypoint.sh --replSet rs0 --bind_ip_all --port 27017 &

# Store the process ID of mongod in a variable
mongod_pid=$!

# Wait for mongod to be fully initialized.
sleep 10

# Execute mongosh with the initialization script.
# It will create rs0 replica set and test user in testDB.
mongosh <<EOF
  try {
    rs.status();
  } catch (err) {
    print("Initiating replica set 'rs0'...");
    rs.initiate({
      _id: 'rs0',
      members: [
        {
          _id: 0,
          host: 'host.docker.internal:27017'
        }
      ]
    });
  }

  db = new Mongo('localhost:27017').getDB('testDB');

  // Check if the user already exists
  if (!db.getUser('test')) {
    // User does not exist, so create it
    print("Creating test user on testDB");
    db.createUser({
      user: 'test',
      pwd: 'test',
      roles: [
        {
          role: 'readWrite',
          db: 'testDB'
        }
      ]
    });
    print('User "test" created.');
  }
EOF

# Wait for mongod to finish (or do other operations)
wait $mongod_pid
