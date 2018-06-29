output "webapp_url" {
    description = "Webapp Endpoing URL"
    value       = "${join("", list("https://", var.name, ".azurewebsites.net"))}"
}