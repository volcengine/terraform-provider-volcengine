resource "volcengine_iam_role_policy_attachment" "foo" {
  role_name = "CustomRoleForPatchManager"
  policy_name = "AdministratorAccess"
  policy_type = "System"
}
