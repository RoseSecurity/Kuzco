terraform {
  required_version = ">= 1.5.6"

  required_providers {
    cloudsmith = {
      source  = "cloudsmith-io/cloudsmith"
      version = ">= 0.0.60"
    }
  }
}

provider "cloudsmith" {
  api_key = "XXXXX"
}

data "cloudsmith_namespace" "this" {
  slug = "RoseSecurity"
}

resource "cloudsmith_repository" "this" {
  description = "A certifiably-awesome private package repository for Kuzco"
  name        = "Kuzco"
  namespace   = data.cloudsmith_namespace.this.slug_perm
  slug        = "kuzco"
}
