data "volcengine_kms_macs" "mac" {
  key_id        = "68093dd1-d1a9-44ce-****-5a88c4bc31ab"
  message       = "VGhpcyBpcyBhIHRlc3QgTWVzc2FnZS4="
  mac_algorithm = "HMAC_SHA_256"
}

data "volcengine_kms_mac_verifications" "verify" {
  key_id        = "68093dd1-d1a9-44ce-****-5a88c4bc31ab"
  message       = "VGhpcyBpcyBhIHRlc3QgTWVzc2FnZS4="
  mac_algorithm = "HMAC_SHA_256"
  mac           = "Vm0D9fk6uDRZD6k9QZE9+d9gpgy6ESSPt0bfaA2p05w="
}