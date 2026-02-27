# Obtain the public key of the specified asymmetric key
data "volcengine_kms_public_keys" "default" {
    keyring_name = "Tf-test"
    key_name = "Test-key2"
}