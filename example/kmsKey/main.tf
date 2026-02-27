resource "volcengine_kms_keyring" "foo" {
  keyring_name   = "tf-test"
  description = "tf-test"
  project_name = "default"
}

resource "volcengine_kms_key" "foo" {
  keyring_name   = volcengine_kms_keyring.foo.keyring_name
  key_name = "mrk-tf-key-mod"
  description = "tf test key-mod"
  tags {
    key = "tfkey3"
    value = "tfvalue3"
  }
}

resource "volcengine_kms_key" "foo1"{
    keyring_name = volcengine_kms_keyring.foo.keyring_name
    key_name = "Tf-test-key-1"
    rotate_state = "Enable"
    rotate_interval = 90
    key_spec = "SYMMETRIC_128"
    description = "Tf test key with SYMMETRIC_128"
    key_usage = "ENCRYPT_DECRYPT"
    protection_level = "SOFTWARE"
    origin = "CloudKMS"
    multi_region = false
    #The scheduled deletion time when deleting the key
    pending_window_in_days = 30
    tags {
        key = "tfk1"
        value = "tfv1"
    }
    tags {
        key = "tfk2"
        value = "tfv2"
    }
}

resource "volcengine_kms_key" "foo2" {
  keyring_name = volcengine_kms_keyring.foo.keyring_name
  key_name = "mrk-Tf-test-key-2"
  key_usage = "ENCRYPT_DECRYPT"
  origin = "External"
  multi_region = true
}

resource "volcengine_kms_key_material" "default" {
    keyring_name = volcengine_kms_keyring.foo.keyring_name
    key_name = volcengine_kms_key.foo2.key_name
    encrypted_key_material = "***"
    import_token = "***"
    expiration_model  =  "KEY_MATERIAL_EXPIRES"
    valid_to          =  1770999621
}