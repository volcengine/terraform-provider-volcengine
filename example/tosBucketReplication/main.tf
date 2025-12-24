resource "volcengine_tos_bucket_replication" "foo" {
  bucket_name = "tflyb78"
  role        = "ServiceRoleforReplicationAccessTOS"
  
  rules {
    id     = "rule3"
    status = "Enabled"
    
    prefix_set = ["documents/", "images/"]
    
    destination {
      bucket                           = "tflyb7-replica1"
      location                         = "cn-beijing"
      storage_class                    = "STANDARD"
      storage_class_inherit_directive = "SOURCE_OBJECT"
    }
    transfer_type = "internal"
    historical_object_replication = "Enabled"
    access_control_translation {
      owner = "BucketOwnerEntrusted"
    }
  }
  
  rules {
    id     = "rule2"
    status = "Disabled"
    
    destination {
      bucket                           = "tflyb7-replica2"
      location                         = "cn-beijing"
      storage_class                    = "IA"
      storage_class_inherit_directive = "DESTINATION_BUCKET"
    }
    access_control_translation {
      owner = "BucketOwnerEntrusted"
    }
  }
}