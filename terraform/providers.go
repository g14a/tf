package terraform

func GetProviders() []string {
	return []string{
		"aci", "acme", "akamai", "alicloud", "archive", "arukras", "auth0",
		"avi", "aviatrix", "aws", "azure", "azuread", "azuredevops", "azurerm",
		"azurestack", "baiducloud", "bigip", "bitbucket", "brightbox", "checkpoint",
		"chef", "cherryservers", "circonus", "ciscoasa", "clc", "cloudamqp", "cloudflare",
		"cloudinit", "cloudscale", "cloudstack", "cobbler", "cohesity", "constellix",
		"consul", "datadog", "digitalocean", "dme", "dns", "dnsimple", "docker", "dome9",
		"dyn", "ecl", "exoscale", "external", "fastly", "flexibleengine", "fortios",
		"genymotion", "github", "gitlab", "google", "google-beta", "grafana", "gridscale",
		"hashicups", "hcloud", "hedvig", "helm", "heroku", "http", "huaweicloud", "huaweicloudstack",
		"icinga2", "ignition", "incapsula", "influxdb", "infoblox", "jdcloud", "kysun", "kubernetes",
		"kubernetes-alpha", "lacework", "launchdarkly", "librato", "linode", "local", "logentries",
		"logicmonitor", "mailgun", "metalcloud", "mongodbatlas", "mso", "mysql", "ncloud", "netlify",
		"newrelic", "nks", "nomad", "ns1", "nsxt", "null", "nutanix", "oci", "okta", "oktaasa",
		"oneandone", "opc", "opennebula", "openstack", "opentelekomcloud", "opsgenie", "oraclepaas",
		"ovh", "packet", "pagerduty", "panos", "postgresql", "powerdns", "prismacloud", "profitbricks",
		"pureport", "rabbitmq", "rancher", "rancher2", "random", "rightscale", "rubrik", "rundeck", "runscope",
		"scaleway", "selectel", "signalfx", "skytap", "softlayer", "spotinst", "stackpath", "statuscake",
		"sumologic", "telefonicaopencloud", "template", "tencentcloud", "terraform", "test", "tfe", "time",
		"tls", "triton", "turbot", "ucloud", "ultradns", "vault", "vcd", "venafi", "vmc", "vra7",
		"vsphere", "vthunder", "vultr", "wavefront",
	}
}
