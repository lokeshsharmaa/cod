build:
	mkdir -p "output"
	go build -o ./output ./cmd/item_service   && go build -o ./output ./cmd/item_service-cli

#
#    ./output/item_service-cli item-service create --message '{
#          "damage": 11,
#          "description": "Et in alias est expedita quos eligendi.",
#          "healing": 1329,
#          "name": "Est velit culpa aspernatur dolorem.",
#          "protection": 120
#       }'
#
#        ./output/item_service-cli item-service get --message '{
#             "id": 1
#          }'
#
#        ./output/item_service-cli item-service delete --message '{
#             "id": 1
#          }'
