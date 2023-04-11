resource "volcengine_rds_mysql_account" "default"{
     instance_id="mysql-e9293705eed6"
     account_name="test"
     account_password="zhangzheming!@ZZM"
     account_type="Normal"
     account_privileges{
          db_name="tf-test-db"
          account_privilege="ReadOnly"
         account_privilege_detail="SELECT,UPDATE,INSERT"
     }
      account_privileges{
          db_name="test-xx"
          account_privilege="ReadOnly"
          account_privilege_detail="SELECT,UPDATE,INSERT"
     }
}