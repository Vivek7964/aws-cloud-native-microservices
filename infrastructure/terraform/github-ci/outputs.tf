output "github_actions_role_arn" {
  value = aws_iam_role.github_actions.arn
}

output "github_oidc_provider_arn" {
  value = aws_iam_openid_connect_provider.github.arn
}

output "ecr_repositories" {
  value = {
    for name, repo in aws_ecr_repository.watchn :
    name => repo.repository_url
  }
}