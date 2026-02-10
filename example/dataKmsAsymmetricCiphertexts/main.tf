data "volcengine_kms_asymmetric_ciphertexts" "encrypt1" {
  key_id = "9601e1af-ad69-42df-****-eaf10ce6a3e9"
  plaintext = "VGhpcyBpcyBhIHBsYWludGV4dCBleGFtcGxlLg=="
  algorithm = "RSAES_OAEP_SHA_256"
}
data "volcengine_kms_asymmetric_ciphertexts" "encrypt2" {
  keyring_name = "Tf-test"
  key_name  = "ec-sm2"
  plaintext = "VGhpcyBpcyBhIHBsYWludGV4dCBleGFtcGxlLg=="
  algorithm = "SM2PKE"
}