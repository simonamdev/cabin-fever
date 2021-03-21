terraform {
  required_providers {
    hcloud = {
      source  = "hetznercloud/hcloud"
      version = "1.25.2"
    }
  }
}

variable "hcloud_token" {}

provider "hcloud" {
  token = var.hcloud_token
}

resource "hcloud_ssh_key" "default" {
  name       = "Cabin Fever Dev SSH"
  public_key = file("~/.ssh/cabinfever.pub")
}

resource "hcloud_server" "cabinfever-dev" {
  name        = "cabinfever-dev"
  image       = "debian-9"
  server_type = "cx11"
}
