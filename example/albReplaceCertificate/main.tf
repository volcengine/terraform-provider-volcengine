# replace server certificate
resource "volcengine_alb_replace_certificate" "foo1" {
    certificate_type = "server"
    old_certificate_id = "cert-bdde0znk524g8dv40or*****"
    update_mode = "new"
    certificate_name = "replaced-server-cert"
    description = "Replaced server certificate"
    project_name = "default"
    public_key = file("/path/server_certificate.pem")
    private_key = file("/path/private_key_rsa.pem")
}
resource "volcengine_alb_replace_certificate" "foo2" {
    certificate_type = "server"
    old_certificate_id = "cert-1pf4a8k8tokcg845wfar*****"
    update_mode = "stock"
    certificate_source = "alb"
    certificate_id = "cert-bdde0znk524g8dv40or*****"
    certificate_name = "replaced-server-cert-stock"
    description = "Replaced server certificate (stock)"
    project_name = "default"
}

# replace ca certificate
resource "volcengine_alb_replace_certificate" "foo3" {
    certificate_type = "ca"
    old_certificate_id = "cert-xoekc6lpu9s054ov5eo*****"    
    update_mode = "new"
    certificate_name = "acc-test-replace"
    ca_certificate      = file("/path/server_certificate.pem")
    description = "acc-test-replace"
    project_name = "default"
}
