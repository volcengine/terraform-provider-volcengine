data "volcengine_kms_key_materials" "default" {
    keyring_name = "Tf-test-1"
    key_name = "Test-3"
    wrapping_key_spec = "RSA_2048"
    wrapping_algorithm = "RSAES_OAEP_SHA_256"
}