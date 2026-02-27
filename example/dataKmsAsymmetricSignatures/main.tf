data "volcengine_kms_asymmetric_signatures" "sign1" {
  key_id       = "516274b3-0cba-4fad-****-c8355e3e8213"
  message      = "VGhpcyBpcyBhIG1lc3NhZ2UgZXhhbXBsZS4="
  message_type = "RAW"
  algorithm    = "RSA_PSS_SHA_256"
}

data "volcengine_kms_asymmetric_signatures" "sign2" {
  key_id       = "516274b3-0cba-4fad-****-c8355e3e8213"
  message      = "KsFMwOobjOMHfYaPl2IgXX6tzziiT+SucmfmXTo2f6U="
  message_type = "DIGEST"
  algorithm    = "RSA_PSS_SHA_256"
}