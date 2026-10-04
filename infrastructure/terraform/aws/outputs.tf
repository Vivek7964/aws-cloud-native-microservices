output "vpc_id" {
  value = aws_vpc.main.id
}

output "public_subnet_ids" {
  value = [
    aws_subnet.public_1.id,
    aws_subnet.public_2.id
  ]
}

output "private_subnet_ids" {
  value = [
    aws_subnet.private_1.id,
    aws_subnet.private_2.id
  ]
}

output "nat_gateway_id" {
  value = aws_nat_gateway.main.id
}

output "internet_gateway_id" {
  value = aws_internet_gateway.main.id
}

output "eks_cluster_name" {
  value = aws_eks_cluster.main.name
}

output "eks_cluster_endpoint" {
  value = aws_eks_cluster.main.endpoint
}

output "eks_cluster_arn" {
  value = aws_eks_cluster.main.arn
}

output "eks_cluster_version" {
  value = aws_eks_cluster.main.version
}

output "eks_node_group_name" {
  value = aws_eks_node_group.main.node_group_name
}

output "ecr_repository_urls" {
  value = {
    ui       = aws_ecr_repository.ui.repository_url
    catalog  = aws_ecr_repository.catalog.repository_url
    carts    = aws_ecr_repository.carts.repository_url
    orders   = aws_ecr_repository.orders.repository_url
    checkout = aws_ecr_repository.checkout.repository_url
    assets   = aws_ecr_repository.assets.repository_url
    activemq = aws_ecr_repository.activemq.repository_url
  }
}