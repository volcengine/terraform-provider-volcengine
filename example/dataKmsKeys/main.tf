data "volcengine_kms_keys" "default" {
  keyring_id = "7a358829-bd5a-4763-ba77-7500ecxxxxxx"
  key_name = ["mrk-tf-key-mod", "mrk-tf-key"]
  key_spec = ["SYMMETRIC_256"]
  description = ["tf-test"]
  key_state = ["Enable"]
  key_usage = ["ENCRYPT_DECRYPT"]
  protection_level = ["SOFTWARE"]
  rotate_state = ["Enable"]
  origin = ["CloudKMS"]
  creation_date_range = ["2025-06-01 19:48:06", "2025-06-04 19:48:06"]
  update_date_range = ["2025-06-01 19:48:06", "2025-06-04 19:48:06"]
  tags {
    key = "tf-k1"
    values = ["tf-v1"]
  }
}