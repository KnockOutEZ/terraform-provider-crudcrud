provider "crudcrud" {
  address = "https://crudcrud.com/api"
  token   = "5ebc57a23671498fb94a605fd7c95362"
}

resource "example_unicorn" "crudcrud" {
 provider = crudcrud
 name = "this_is_an_item1"
 age = 421111
 colour = "blue"
}
