provider "aws" {
  region = "ap-south-1"
}

resource "aws_db_instance" "myrds" {
  identifier = "mydb"
  engine = "mysql"
  engine_version = "8.0.32"
  instance_class = "db.t2.micro"
  port = 3306
  allocated_storage = 20
  storage_type = "gp2"
  db_name = "mydatabase"
  username = "admin"
  password = "Admin!1994"
  publicly_accessible = true
  skip_final_snapshot = true

  #  provisioner "local-exec" {
  #   command = <<EOT
  #     mysql -h ${aws_db_instance.myrds.endpoint} -u ${aws_db_instance.myrds.username} -p${aws_db_instance.myrds.password} <<EOF
  #       CREATE DATABASE my_database;
  #       USE my_database;
  #       CREATE TABLE my_table (Id int, Name varchar(255), City varchar(255));
  #       INSERT INTO my_table (Id, Name, City) VALUES (1,Ram,Mumbai);
  #       INSERT INTO my_table (Id, Name, City) VALUES (2,Bheem,Pune);
  #       INSERT INTO my_table (Id, Name, City) VALUES (3,Shyam,Vizag);
  #     EOF
  #   EOT
  # }

}









