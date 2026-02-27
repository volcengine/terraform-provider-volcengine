# It is necessary to first use data volcengine_kms_key_materials to obtain the import materials, such as import_token, public_key.
# Reference document: https://www.volcengine.com/docs/6476/144950?lang=zh
resource "volcengine_kms_key_material" "default" {
    keyring_name = "Tf-test-1"
    key_name = "Test-3"
    key_id = "8798cd1e-****-4f9b-****-d51847ad53ae"
    encrypted_key_material = "***"
    import_token = "***"
    expiration_model  =  "KEY_MATERIAL_EXPIRES"
    valid_to          =  1770969621
}