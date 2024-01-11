resource "volcengine_cloud_monitor_contact_group" "foo" {
    name = "tfgroup"
    description = "tftest"
    contacts_id_list = ["1737376113733353472", "1737375997680111616"]
}