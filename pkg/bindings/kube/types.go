package kube

import (
	"net"
)

// PlayOptions are optional options for replaying kube YAML files
//
//go:generate go run ../generator/generator.go PlayOptions
type PlayOptions struct {
	// Annotations - Annotations to add to Pods
	Annotations map[string]string
	// Authfile - path to an authentication file.
	Authfile *string
	// CertDir - to a directory containing TLS certifications and keys.
	CertDir *string
	// Username for authenticating against the registry.
	Username *string
	// Password for authenticating against the registry.
	Password *string
	// Network - name of the networks to connect to.
	Network *[]string
	// NoHosts - do not generate /etc/hosts file in pod's containers
	NoHosts *bool
	// Quiet - suppress output when pulling images.
	Quiet *bool
	// SignaturePolicy - path to a signature-policy file.
	SignaturePolicy *string
	// SkipTLSVerify - skip https and certificate validation when
	// contacting container registries.
	SkipTLSVerify *bool `schema:"-"`
	// SeccompProfileRoot - path to a directory containing seccomp
	// profiles.
	SeccompProfileRoot *string
	// StaticIPs - Static IP address used by the pod(s).
	StaticIPs *[]net.IP
	// StaticMACs - Static MAC address used by the pod(s).
	StaticMACs *[]net.HardwareAddr
	// ConfigMaps - slice of pathnames to kubernetes configmap YAMLs.
	ConfigMaps *[]string
	// LogDriver for the container. For example: journald
	LogDriver *string
	// LogOptions for the container. For example: journald
	LogOptions *[]string
	// Start - don't start the pod if false
	Start *bool
	// Userns - define the user namespace to use.
	Userns *string
	// Force - remove volumes on --down
	Force *bool
}

// DownOptions are optional options for tearing down kube YAML files to a k8s cluster
//
//go:generate go run ../generator/generator.go DownOptions
type DownOptions struct {
	// Force - remove volumes on --down
	Force *bool
}
