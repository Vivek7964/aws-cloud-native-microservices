variable "project_name" {
  description = "Project name"
  type        = string
  default     = "watchn"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "dev"
}

variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "eks_version" {
  description = "EKS Kubernetes version"
  type        = string
  default     = "1.35"
}

variable "node_instance_type" {
  description = "EKS worker node instance type"
  type        = string
  default     = "c7i-flex.large"
}

variable "node_min_size" {
  description = "Minimum worker nodes"
  type        = number
  default     = 1
}

variable "node_desired_size" {
  description = "Desired worker nodes"
  type        = number
  default     = 2
}

variable "node_max_size" {
  description = "Maximum worker nodes"
  type        = number
  default     = 2
}