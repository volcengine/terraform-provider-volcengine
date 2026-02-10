resource "volcengine_kms_asymmetric_ciphertext" "encrypt_stable" {
  key_id = "9601e1af-ad69-42df-****-eaf10ce6a3e9"
  plaintext = "VGhpcyBpcyBhIHBsYWludGV4dCBleGFtcGxLg=="
  algorithm = "RSAES_OAEP_SHA_256"
}

data "volcengine_kms_asymmetric_plaintexts" "decrypt" {
  key_id          = volcengine_kms_asymmetric_ciphertext.encrypt_stable.key_id
  ciphertext_blob = volcengine_kms_asymmetric_ciphertext.encrypt_stable.ciphertext_blob
  algorithm       = "RSAES_OAEP_SHA_256"
}
