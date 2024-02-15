# MongoDB Cheat Sheet

## Create Role and User

To create a role with specific privileges and then create a user with that role:

```javascript
use testDB

// Create a custom role with specific privileges
db.createRole({
   role: "testRole",
   privileges: [
      { resource: { db: "testDB", collection: "my-collection-1" }, actions: [ "find", "insert", "update", "remove", "createIndex", "dropIndex", "createCollection" ] },
      { resource: { db: "testDB", collection: "my-collection-2" }, actions: [ "find", "insert", "update", "remove", "createIndex", "dropIndex", "createCollection" ] },
      { resource: { db: "testDB", collection: "my-collection-3" }, actions: [ "find" ] },
      { resource: { db: "testDB", collection: "my-collection-4" }, actions: [ "find" ] }
   ],
   roles: []
})

// Create the user with the custom role
db.createUser({
   user: "test",
   pwd: "test",
   roles: [ { role: "testRole", db: "testDB" } ]
})
