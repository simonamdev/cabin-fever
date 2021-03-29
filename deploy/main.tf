terraform {
  required_providers {
    hcloud = {
      source  = "hetznercloud/hcloud"
      version = "1.25.2"
    }
  }
}

variable "hcloud_token" {}
variable "cloud_init" {}

provider "hcloud" {
  token = var.hcloud_token
}

resource "hcloud_ssh_key" "default" {
  name       = "Cabin Fever Dev SSH"
  public_key = file("~/.ssh/cabinfever.pub")
}

resource "hcloud_server" "cabinfever-dev" {
  name        = "cabinfever-dev"
  image       = "ubuntu-20.04"
  server_type = "cx11"
  ssh_keys    = [hcloud_ssh_key.default.id]
  user_data   = var.cloud_init


  provisioner "file" {
    source      = "../cabinserver-prod"
    destination = "/srv/cabinserver"

    connection {
      type        = "ssh"
      user        = "root"
      host        = self.ipv4_address
      private_key = file("~/.ssh/cabinfever")
    }
  }

  provisioner "remote-exec" {
    inline = [
      "cloud-init status --wait",
    ]

    connection {
      type        = "ssh"
      user        = "root"
      host        = self.ipv4_address
      private_key = file("~/.ssh/cabinfever")
    }
  }

  provisioner "file" {
    source      = "../cabinserver/conf/supervisor.conf"
    destination = "/etc/supervisor/conf.d/cabinserver.conf"

    connection {
      type        = "ssh"
      user        = "root"
      host        = self.ipv4_address
      private_key = file("~/.ssh/cabinfever")
    }
  }

  provisioner "remote-exec" {
    inline = [
      "sudo chmod +x /srv/cabinserver",
      "sudo service supervisor restart",
    ]

    connection {
      type        = "ssh"
      user        = "root"
      host        = self.ipv4_address
      private_key = file("~/.ssh/cabinfever")
    }
  }


}
