package resourcebps

import (
	"github.com/fatih/color"
)

func AWSDBInstanceBP() {
	color.Green("\nresource \"aws_db_instance\" \"foo\" {\n  allocated_storage    = 20\n  storage_type         = \"gp2\"\n  engine               = \"mysql\"\n  engine_version       = \"5.7\"\n  instance_class       = \"db.t2.micro\"\n  name                 = \"mydb\"\n  username             = \"foo\"\n  password             = \"foobarbaz\"\n  parameter_group_name = \"default.mysql5.7\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_instance\n\n")
}

func AWSDBClusterSnapshotBP() {
	color.Green("\nresource \"aws_db_cluster_snapshot\" \"example\" {\n  db_cluster_identifier          = aws_rds_cluster.example.id\n  db_cluster_snapshot_identifier = \"resourcetestsnapshot1234\"\n  tags = {\n    Name = \"HelloWorld\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_cluster_snapshot\n\n")
}

func AWSDBEventSubscriptionBP() {
	color.Green("\nresource \"aws_db_event_subscription\" \"foo\" {\n  name      = \"rds-event-sub\"\n  sns_topic = aws_sns_topic.default.arn\n\n  source_type = \"db-instance\"\n  source_ids  = [aws_db_instance.default.id]\n\n  event_categories = [\n    \"availability\",\n    \"deletion\",\n    \"failover\",\n    \"failure\",\n    \"low storage\",\n    \"maintenance\",\n    \"notification\",\n    \"read replica\",\n    \"recovery\",\n    \"restoration\",\n  ]\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_event_subscription\n\n")
}

func AWSDBInstanceRoleAssociationBP() {
	color.Green("\nresource \"aws_db_instance_role_association\" \"foo\" {\n  db_instance_identifier = aws_db_instance.example.id\n  feature_name           = \"S3_INTEGRATION\"\n  role_arn               = aws_iam_role.example.arn\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_instance_role_association\n\n")
}

func AWSDBOptionGroupBP() {
	color.Green("\nresource \"aws_db_option_group\" \"foo\" {\n  name                     = \"option-group-test-terraform\"\n  option_group_description = \"Terraform Option Group\"\n  engine_name              = \"sqlserver-ee\"\n  major_engine_version     = \"11.00\"\n\n  option {\n    option_name = \"Timezone\"\n\n    option_settings {\n      name  = \"TIME_ZONE\"\n      value = \"UTC\"\n    }\n  }\n\n  option {\n    option_name = \"SQLSERVER_BACKUP_RESTORE\"\n\n    option_settings {\n      name  = \"IAM_ROLE_ARN\"\n      value = aws_iam_role.example.arn\n    }\n  }\n\n  option {\n    option_name = \"TDE\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_option_group\n\n")
}

func AWSDBParameterOptionBP() {
	color.Green("\nresource \"aws_db_parameter_group\" \"foo\" {\n  name   = \"rds-pg\"\n  family = \"mysql5.6\"\n\n  parameter {\n    name  = \"character_set_server\"\n    value = \"utf8\"\n  }\n\n  parameter {\n    name  = \"character_set_client\"\n    value = \"utf8\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_parameter_group\n\n")
}

func AWSDBProxyBP() {
	color.Green("\nresource \"aws_db_proxy\" \"foo\" {\n  name                   = \"example\"\n  debug_logging          = false\n  engine_family          = \"MYSQL\"\n  idle_client_timeout    = 1800\n  require_tls            = true\n  role_arn               = aws_iam_role.example.arn\n  vpc_security_group_ids = [aws_security_group.example.id]\n  vpc_subnet_ids         = [aws_subnet.example.id]\n\n  auth {\n    auth_scheme = \"SECRETS\"\n    description = \"example\"\n    iam_auth    = \"DISABLED\"\n    secret_arn  = aws_secretsmanager_secret.example.arn\n  }\n\n  tags = {\n    Name = \"example\"\n    Key  = \"value\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_proxy\n\n")
}

func AWSDBProxyDefaultTargetGroupBP() {
	color.Green("\nresource \"aws_db_proxy_default_target_group\" \"foo\" {\n  db_proxy_name = aws_db_proxy.example.name\n\n  connection_pool_config {\n    connection_borrow_timeout    = 120\n    init_query                   = \"SET x=1, y=2\"\n    max_connections_percent      = 100\n    max_idle_connections_percent = 50\n    session_pinning_filters      = [\"EXCLUDE_VARIABLE_SETS\"]\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_proxy_default_target_group\n\n")
}

func AWSDBProxyTargetGroupBP() {
	color.Green("\nresource \"aws_db_proxy_target\" \"example\" {\n  db_instance_identifier = aws_db_instance.example.id\n  db_proxy_name          = aws_db_proxy.example.db_proxy_name\n  target_group_name      = aws_db_proxy_default_target_group.example.name\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_proxy_target\n\n")
}

func AWSDBSecurityGroupBP() {
	color.Green("\nresource \"aws_db_security_group\" \"default\" {\n  name = \"rds_sg\"\n\n  ingress {\n    cidr = \"10.0.0.0/24\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_security_group\n\n")
}

func AWSDBSnapshotBP() {
	color.Green("\nresource \"aws_db_snapshot\" \"test\" {\n  db_instance_identifier = aws_db_instance.bar.id\n  db_snapshot_identifier = \"testsnapshot1234\"\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_snapshot\n\n")
}

func AWSDBSubnetGroupBP() {
	color.Green("\nresource \"aws_db_subnet_group\" \"default\" {\n  name       = \"main\"\n  subnet_ids = [aws_subnet.frontend.id, aws_subnet.backend.id]\n\n  tags = {\n    Name = \"My DB subnet group\"\n  }\n}")
	color.Yellow("\nCheckout https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_subnet_group\n\n")
}
