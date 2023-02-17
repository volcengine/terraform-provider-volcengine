resource "volcengine_rds_mysql_account" "default"{
     instance_id="mysql-xxx"
     account_name="xxx"
     account_password="xxx"
     account_type="Normal"
     account_privileges{
          db_name="xxx"
          account_privilege="Custom"
          account_privilege_detail="SELECT,UPDATE,INSERT"
     }
      account_privileges{
          db_name="xx"
          account_privilege="Custom"
          account_privilege_detail="SELECT,UPDATE,INSERT"
     }
}