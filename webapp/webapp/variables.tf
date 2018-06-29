variable "name" {
    description = "The name of the web app"
}

variable "resource_group_name" {
    description = "The name of the resource group in which the resources will be created."
}

variable "location" {
    description = "Region where the resources are created."
}

variable "service_plan_name" {
    description = "The name of the App Service Plan, default = $web_app_name"
    default     = ""
}