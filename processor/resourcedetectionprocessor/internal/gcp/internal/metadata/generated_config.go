// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
)

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac)
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for resourcedetectionprocessor/gcp resource attributes.
type ResourceAttributesConfig struct {
	CloudAccountID                   ResourceAttributeConfig `mapstructure:"cloud.account.id"`
	CloudAvailabilityZone            ResourceAttributeConfig `mapstructure:"cloud.availability_zone"`
	CloudPlatform                    ResourceAttributeConfig `mapstructure:"cloud.platform"`
	CloudProvider                    ResourceAttributeConfig `mapstructure:"cloud.provider"`
	CloudRegion                      ResourceAttributeConfig `mapstructure:"cloud.region"`
	FaasID                           ResourceAttributeConfig `mapstructure:"faas.id"`
	FaasInstance                     ResourceAttributeConfig `mapstructure:"faas.instance"`
	FaasName                         ResourceAttributeConfig `mapstructure:"faas.name"`
	FaasVersion                      ResourceAttributeConfig `mapstructure:"faas.version"`
	GcpCloudRunJobExecution          ResourceAttributeConfig `mapstructure:"gcp.cloud_run.job.execution"`
	GcpCloudRunJobTaskIndex          ResourceAttributeConfig `mapstructure:"gcp.cloud_run.job.task_index"`
	GcpGceInstanceHostname           ResourceAttributeConfig `mapstructure:"gcp.gce.instance.hostname"`
	GcpGceInstanceName               ResourceAttributeConfig `mapstructure:"gcp.gce.instance.name"`
	GcpGceInstanceGroupManagerName   ResourceAttributeConfig `mapstructure:"gcp.gce.instance_group_manager.name"`
	GcpGceInstanceGroupManagerRegion ResourceAttributeConfig `mapstructure:"gcp.gce.instance_group_manager.region"`
	GcpGceInstanceGroupManagerZone   ResourceAttributeConfig `mapstructure:"gcp.gce.instance_group_manager.zone"`
	HostID                           ResourceAttributeConfig `mapstructure:"host.id"`
	HostName                         ResourceAttributeConfig `mapstructure:"host.name"`
	HostType                         ResourceAttributeConfig `mapstructure:"host.type"`
	K8sClusterName                   ResourceAttributeConfig `mapstructure:"k8s.cluster.name"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		CloudAccountID: ResourceAttributeConfig{
			Enabled: true,
		},
		CloudAvailabilityZone: ResourceAttributeConfig{
			Enabled: true,
		},
		CloudPlatform: ResourceAttributeConfig{
			Enabled: true,
		},
		CloudProvider: ResourceAttributeConfig{
			Enabled: true,
		},
		CloudRegion: ResourceAttributeConfig{
			Enabled: true,
		},
		FaasID: ResourceAttributeConfig{
			Enabled: true,
		},
		FaasInstance: ResourceAttributeConfig{
			Enabled: true,
		},
		FaasName: ResourceAttributeConfig{
			Enabled: true,
		},
		FaasVersion: ResourceAttributeConfig{
			Enabled: true,
		},
		GcpCloudRunJobExecution: ResourceAttributeConfig{
			Enabled: true,
		},
		GcpCloudRunJobTaskIndex: ResourceAttributeConfig{
			Enabled: true,
		},
		GcpGceInstanceHostname: ResourceAttributeConfig{
			Enabled: false,
		},
		GcpGceInstanceName: ResourceAttributeConfig{
			Enabled: false,
		},
		GcpGceInstanceGroupManagerName: ResourceAttributeConfig{
			Enabled: true,
		},
		GcpGceInstanceGroupManagerRegion: ResourceAttributeConfig{
			Enabled: true,
		},
		GcpGceInstanceGroupManagerZone: ResourceAttributeConfig{
			Enabled: true,
		},
		HostID: ResourceAttributeConfig{
			Enabled: true,
		},
		HostName: ResourceAttributeConfig{
			Enabled: true,
		},
		HostType: ResourceAttributeConfig{
			Enabled: true,
		},
		K8sClusterName: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}
