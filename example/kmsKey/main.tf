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
    encrypted_key_material = "zPIkW0zx/KJ4qo7+BxrKnzfqzDAMJYmwiFESpvYqEdNIGOfuFLOpyC3sbxuOEi9cQoJIOEj8P8ceg8yEcyIaEHQKe6We+L1ee0zsITJ68NBdVpVYlXvSmNe34gN5Vd257K5TAmKuznbb+TtvXfLTzrEPVkc4i4iDLlXk/nOMrEfeJjuW08L4Isdn3ggUfPDctilX4pFcALV79Gx4843ARFDOII0xZyR3GhQbRhny8pxh8uqkNTqu4pUxnuQzGhW7VwBct/giw9KdTGw+TMUjh0qa0IyMEEpdZ3DomjJ5f/ioFrfgjukVd/y+mMkBWiFyni1bnXNjDuWeGRofN0WFWw=="
    import_token = "UlNBXzIwNDg6UlNBRVNfT0FFUF9TSEFfMjU2OpYpU60nZEbSp+kEBPKPd65tX2NuLWd1aWxpbi1ib2VfSyoOn/kFRd+BtTVC9SCstGz8hmkAAAAAb2uo2cNwvPrSDHDx+sx7k1L9AzNdeDUnYSj7qgDnX9IOLRdUTS+zWvhMipgrrDrYs9kDNFe+Pvpngv051ipu/7RBMdLHWFrWWbJipCuD02ojydCy/CFD3axmOEwyqpXeHdi/3SK1E6yH0HULVCs8IkoO972fixWToLeiz5cH4KE8/xX59WeYPPwEglrtXJM8rbHBYnfXiQ3E7Fx3368UOEUePgCtKqQp40gaiMpVVpYAIAgfFe25ZQzTxVDvLjAa4DcTIkPnyBCfIfH5Dnniq52t8WCq1U+SZ9FMWYE3leBGeYax4rKfQySyCW2dytM1laUaWt2fr/iwC6tOllK7ZXXh1p52en7IcK5K5fNSF8GF01xWwDHOJTKzQDapqdYutLlUD7jjlYx0HUTVmz71hbyDfF/n/gWPOgg/RS3BSquLKKoqZYofqbG+iIbyUrHfzq3Ph+AUORcetgCFaOtYUfklDc1mVBDGz6Dfwp6WPIzhlXUcqL4NtQlN3BnR2SKldcmshQlg+AokcTchg96IrS3Sx5DOBcAQPye3XsSEN5xTM2M+4whccjYW0VvFn/JL/b2/W1Cbx9PYzRRv4MtZmIcbJX3n4XIT45k/+yL0mDo/IJBNg9SNy80vUFuH2AKORylHC58Vm74wAbMUiXK94YZNIlU/Swe1w6xnI++N0sbHP5W6TLcv9vazvzVIorbXZEP8esC+3TnChChu39/z1lGgSzDAYVb2EGr7UQgzjQjA+coAFxxZv7MhzEKNHuhjY4J1HFXZMpaB+riFDqxk71DTqn5b9OJoy3svUOZ8DbTtpWpB4eC/Gh6FagqgKNRVtc5tSHzZSjRoCC2CWJ4CRM4mv4qpd+t+2aOwLJcxYTw74ZqP3K7Z1t+hk58WrzWRGLZ+OxpGxb0D2CHeSV2ptzg7xhHDd4OIxG0llsVLW7w0LtlpbCkyUySspncUsGeH15C1k+l7oHeZhLIleDV4lplmo+5LKPnvZXgsnWOF60X4RD1ex3psBKchHqLi3tibbRTZX7rO5IXJGH2zs7OaW/uEvLXX45UtdNn4uGFv8er9gCH45Tr9ZJGuJ3htvvC8OLmwWR7Vq4+We2FXgp51kQ0nqMrZDn8smDP/6XjU3Z/ILgid6ha2qFmF63UV6KC2GoN61QTL5ahJZqojz855VO0Xe7fFggL+1rFnSjFgndIn/MIAM8eXosg+jD96k1+OxQRfTycRCWllrDAjg8cpYhjyHWGgvqL3FAlBxVv2iqBjGKZb0kutZNwQpq9LUxbL9vj9wA/AF2aSrc3wkcol5YPBIGBVC1asXwWAGpyjLwY3U4KhFCOFbOmjxkKeXMEXJODeVjPXPQweQrRNSYr1NjJhMm83eo2pm5obIyynuIp6U65iu81Xkm75vePRmwDHdCzzj4+Tmo/CyTMHkzSxokIARXCGi713hJGInad0QS4UfYcJqZebSsvKZnMfixQNHoo8/p3GAsdzxiohj7oo+ig57qhOJD4Ubv8UuBeLTJC4VPIOecSvEkkXWHZh3A/PrDJ+473Ez0YeI99kgSHnGRedKSGplhJhRggcuJuA0mttvJzdwDnw62fIVgGAqdgFyxD8fi32vgx7mI6Wk9K0V+14GcEdE8HD388CbVBFbbyLA/uDpISY531Hr2W0/PzvQtAFjw=="
    expiration_model  =  "KEY_MATERIAL_EXPIRES"
    valid_to          =  1770999621
}