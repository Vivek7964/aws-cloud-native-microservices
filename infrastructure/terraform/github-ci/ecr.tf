locals {
  repositories = [
    "ui",
    "catalog",
    "carts",
    "orders",
    "checkout",
    "assets",
    "activemq"
  ]
}

resource "aws_ecr_repository" "watchn" {
  for_each = toset(local.repositories)

  name                 = "${var.project_name}/${each.value}"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}