variable "project_id" {
  description = "The GCP project ID."
  type        = string
}

variable "region" {
  description = "The GCP region where resources will be created."
  type        = string
  default     = "us-central1"
}

variable "zone" {
  description = "The GCP zone where the VM will be created."
  type        = string
  default     = "us-central1-a"
}

variable "instance_name" {
  description = "The name of the VM instance."
  type        = string
  default     = "my-instance"
}

variable "machine_type" {
  description = "The machine type of the VM."
  type        = string
  default     = "e2-medium"
}

variable "image_family" {
  description = "The image family of the VM."
  type        = string
  default     = "ubuntu-2004-lts"
}

variable "image_project" {
  description = "The project of the image."
  type        = string
  default     = "ubuntu-os-cloud"
}

variable "ssh_user" {
  description = "Username for SSH access."
  type        = string
  default     = "terraform-user"
}

variable "ssh_pub_key_path" {
  description = "Path to the SSH public key to be used for access."
  type        = string
  // Ensure to replace this default path with the actual path where your public SSH key is stored
  default     = "~/.ssh/id_rsa.pub"
}
