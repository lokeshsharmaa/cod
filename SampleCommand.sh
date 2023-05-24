./externalservice-cli externalservice create-character --body '{                                                                                                                                                                      1 ↵ lokesh@lokesh
          "description": "Test",
          "experience": 19,
          "health": 12,
          "name": "Test1"
       }'

./externalservice-cli externalservice update-character --body '{
          "description": "Test",
          "experience": 20,
          "health": 11,
          "name": "Test1"
   }' --id 1

./externalservice-cli externalservice get-character --id 1

./externalservice-cli externalservice delete-character --id 1


./externalservice-cli externalservice create-item --body '{
      "damage": 7411,
      "description": "Voluptatem aut adipisci consequuntur.",
      "healing": 852,
      "name": "Illum eligendi sit adipisci deleniti libero.",
      "protection": 826
   }'

./externalservice-cli externalservice update-item --body '{
      "damage": 12,
      "description": "Test",
      "healing": 85,
      "name": "Illi deleniti libero.",
      "protection": 82
   }' --id 1

./externalservice-cli externalservice get-character --id 1

./externalservice-cli externalservice get-inventory --character-id 1

./externalservice-cli externalservice add-item-to-inventory --body '{
  "item_id": 1
}' --character-id 1

./externalservice-cli externalservice add-item-to-inventory --body '{
  "item_id": 2
}' --character-id 1

./externalservice-cli externalservice remove-item-from-inventory: --body '{
  "item_id": 2
}' --character-id 1

./externalservice-cli externalservice remove-item-from-inventory --item-id 1  --character-id 1

