# --------------------------------
# ECR Repositories
# --------------------------------

resource "aws_ecr_repository" "ui" {
  name                 = "${var.project_name}/ui"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = {
    Name = "${var.project_name}/ui"
  }
}

resource "aws_ecr_repository" "catalog" {
  name                 = "${var.project_name}/catalog"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = {
    Name = "${var.project_name}/catalog"
  }
}

resource "aws_ecr_repository" "carts" {
  name                 = "${var.project_name}/carts"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = {
    Name = "${var.project_name}/carts"
  }
}

resource "aws_ecr_repository" "orders" {
  name                 = "${var.project_name}/orders"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = {
    Name = "${var.project_name}/orders"
  }
}

resource "aws_ecr_repository" "checkout" {
  name                 = "${var.project_name}/checkout"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = {
    Name = "${var.project_name}/checkout"
  }
}

resource "aws_ecr_repository" "assets" {
  name                 = "${var.project_name}/assets"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = {
    Name = "${var.project_name}/assets"
  }
}

resource "aws_ecr_repository" "activemq" {
  name                 = "${var.project_name}/activemq"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = {
    Name = "${var.project_name}/activemq"
  }
}