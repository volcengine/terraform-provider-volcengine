data "volcengine_kms_data_keys" "data_key" {
  key_id          = "c44870c3-f33b-421a-****-a2bba37c993e"
  number_of_bytes = 1024
}

data "volcengine_kms_plaintexts" "default" {
  ciphertext_blob = data.volcengine_kms_data_keys.data_key.data_key_info[0].ciphertext_blob
}