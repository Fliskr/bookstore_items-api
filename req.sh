# curl -X GET localhost:3335/ping
# echo
# curl -X GET localhost:3335/items/tkcZ9nUBNDr0DwQd6eYo
# echo
# curl -X GET localhost:3335/items/tkcZ9nUBNDr0DwQd6eY
# echo
# curl -X POST localhost:3335/items?access_token=123 -d '{"price":100.0,"title":"new item","description":{"plain_text":"some item that seems to be item", "html":"<div style=\"color:red\">123</div>"},"price":100.1,"available_quantity":100,"sold_quantity":100,"status":"created"}'
# echo
echo
# curl -X POST localhost:3335/items/search -d '{"equals":[{"field":"available_quantity","value":"100"}]}'
# curl -X POST localhost:3335/items/search -d '{"equals":[]}'


# curl -X POST localhost:3335/items?access_token=321 -d '{"price":100.0,"title":"new item","description":{"plain_text":"some item that seems to be item", "html":"<div style=\"color:red\">123</div>"},"price":100.1,"available_quantity":100,"sold_quantity":100,"status":"created"}'
echo
# curl -X DELETE localhost:3335/items/tUcZ9nUBNDr0DwQdq-aY
echo
curl -X PUT localhost:3335/items -d '{"id":"SAre-XUBhtkDqsMy9DB2","price":100.0,"title":"new item","description":{"plain_text":"some item that seems to be item", "html":"<div style=\"color:red\">123</div>"},"price":100.1,"available_quantity":100,"sold_quantity":100,"status":"created"}'
echo
curl -X GET localhost:3335/items/SAre-XUBhtkDqsMy9DB2