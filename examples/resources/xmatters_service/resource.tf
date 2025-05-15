// Basic:
resource "xmatters_service" "basic_service" {
  name = "Shopping Cart"
  type = "APPLICATION"
}

// Advanced:
resource "xmatters_service" "create_modify_service" {
  name        = "Shopping Cart"
  description = "Service to monitor shopping cart abandonment rates"
  type        = "APPLICATION"
  tier        = "GOLD"
  owner       = "a9e17d0e-6082-47af-b656-ff34a12736b3"
  links = [
    {
      link_text = "Shopping Cart Service"
      url       = "https://github.com/my_account/Shopping_Cart_service"
    }
  ]
}
