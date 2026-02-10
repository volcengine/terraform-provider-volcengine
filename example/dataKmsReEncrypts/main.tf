resource "volcengine_kms_ciphertext" "encrypt_stable" {
  key_id    = "c44870c3-f33b-421a-****-a2bba37c993e"
  plaintext = "VGhpcyBpcyBhIHBsYWludGV4dCBleGFtcGxlLg=="
}

data "volcengine_kms_re_encrypts" "re_encrypt_changing" {
  new_key_id              = "33e6ae1f-62f6-415a-****-579f526274cc"
  source_ciphertext_blob  = volcengine_kms_ciphertext.encrypt_stable.ciphertext_blob
}