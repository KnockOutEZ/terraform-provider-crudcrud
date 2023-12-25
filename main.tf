provider "crudcrud" {
  address = "https://crudcrud.com/api"
  token   = "f7b6da402e194650b3ce879659c04a50"
}

resource "example_unicorn" "crudcrud" {
 provider = crudcrud
 name = "this_is_an_item"
 age = 42
 colour = "red"
}
