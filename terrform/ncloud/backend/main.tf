provider "ncloud" {
  access_key = var.access_key
  secret_key = var.secret_key
  region     = var.region
}

resource "ncloud_server" "server" {
    name = "tf-test-vm-${terraform.workspace}"
    server_image_product_code = "SPSW0LINUX000046"
    server_product_code = "SPSVRHICPU000001"
}