# Character on Duty (CoD)

Character on Duty (CoD) is an application that allows users to create, manage, and interact with game characters, inventories, and items. It provides a set of API endpoints for performing CRUD operations on characters, inventories, and items.

## Features

- Create a new character with a unique name, description, health, and experience.
- Fetch a character by its ID.
- Update a character's information, including name, description, health, and experience.
- Delete a character by its ID.
- Get a character's inventory.
- Add an item to a character's inventory.
- Remove an item from a character's inventory. 
- CRUD on items through item service


## Technologies Used

- Golang
- Goa (goa.design/goa)
- gRPC (google.golang.org/grpc)
- HTTP/JSON

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/lokeshsharmaa/cod.git

2. Build:
    ```
    cd cod
    make build 

## API Endpoints
The following API endpoints are available:

- POST /characters: Create a new character.
- GET /characters/{id}: Fetch a character by its ID.
- PUT /characters/{id}: Update a character by its ID.
- DELETE /characters/{id}: Delete a character by its ID.
- GET /characters/{id}/inventory: Get a character's inventory.
- POST /characters/{id}/inventory/items: Add an item to a character's inventory.
- DELETE /characters/{id}/inventory/items/{itemID}: Remove an item from a character's inventory.


## Running Each Service

### Character service (grpc: localhost:8081)
```
cd characterservice
./output/characterservice
```

### Inventory service (grpc: localhost:8082)
```
cd inventoryservice
./output/inventoryservice
```   

### Item service (grpc: localhost:8083)
```
cd itemservice
./output/item_service
```   

### External service (http: localhost:8080)
```
cd externalservice
./output/externalservice
```  

## TODO
1. Integrating Item service client
2. Better Response Handling

## Refactoring

- All service can be included in a single design.go
- Better API definition