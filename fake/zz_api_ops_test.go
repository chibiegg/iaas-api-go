// Copyright 2022 The sacloud/iaas-api-go Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// generated by 'github.com/sacloud/iaas-api-go/internal/tools/gen-api-fake-op'; DO NOT EDIT

package fake

import (
	"testing"

	"github.com/sacloud/iaas-api-go"
)

func TestResourceOps(t *testing.T) {

	if op, ok := NewArchiveOp().(iaas.ArchiveAPI); !ok {
		t.Fatalf("%s is not iaas.Archive", op)
	}

	if op, ok := NewAuthStatusOp().(iaas.AuthStatusAPI); !ok {
		t.Fatalf("%s is not iaas.AuthStatus", op)
	}

	if op, ok := NewAutoBackupOp().(iaas.AutoBackupAPI); !ok {
		t.Fatalf("%s is not iaas.AutoBackup", op)
	}

	if op, ok := NewBillOp().(iaas.BillAPI); !ok {
		t.Fatalf("%s is not iaas.Bill", op)
	}

	if op, ok := NewBridgeOp().(iaas.BridgeAPI); !ok {
		t.Fatalf("%s is not iaas.Bridge", op)
	}

	if op, ok := NewCDROMOp().(iaas.CDROMAPI); !ok {
		t.Fatalf("%s is not iaas.CDROM", op)
	}

	if op, ok := NewCertificateAuthorityOp().(iaas.CertificateAuthorityAPI); !ok {
		t.Fatalf("%s is not iaas.CertificateAuthority", op)
	}

	if op, ok := NewContainerRegistryOp().(iaas.ContainerRegistryAPI); !ok {
		t.Fatalf("%s is not iaas.ContainerRegistry", op)
	}

	if op, ok := NewCouponOp().(iaas.CouponAPI); !ok {
		t.Fatalf("%s is not iaas.Coupon", op)
	}

	if op, ok := NewDatabaseOp().(iaas.DatabaseAPI); !ok {
		t.Fatalf("%s is not iaas.Database", op)
	}

	if op, ok := NewDiskOp().(iaas.DiskAPI); !ok {
		t.Fatalf("%s is not iaas.Disk", op)
	}

	if op, ok := NewDiskPlanOp().(iaas.DiskPlanAPI); !ok {
		t.Fatalf("%s is not iaas.DiskPlan", op)
	}

	if op, ok := NewDNSOp().(iaas.DNSAPI); !ok {
		t.Fatalf("%s is not iaas.DNS", op)
	}

	if op, ok := NewEnhancedDBOp().(iaas.EnhancedDBAPI); !ok {
		t.Fatalf("%s is not iaas.EnhancedDB", op)
	}

	if op, ok := NewESMEOp().(iaas.ESMEAPI); !ok {
		t.Fatalf("%s is not iaas.ESME", op)
	}

	if op, ok := NewGSLBOp().(iaas.GSLBAPI); !ok {
		t.Fatalf("%s is not iaas.GSLB", op)
	}

	if op, ok := NewIconOp().(iaas.IconAPI); !ok {
		t.Fatalf("%s is not iaas.Icon", op)
	}

	if op, ok := NewInterfaceOp().(iaas.InterfaceAPI); !ok {
		t.Fatalf("%s is not iaas.Interface", op)
	}

	if op, ok := NewInternetOp().(iaas.InternetAPI); !ok {
		t.Fatalf("%s is not iaas.Internet", op)
	}

	if op, ok := NewInternetPlanOp().(iaas.InternetPlanAPI); !ok {
		t.Fatalf("%s is not iaas.InternetPlan", op)
	}

	if op, ok := NewIPAddressOp().(iaas.IPAddressAPI); !ok {
		t.Fatalf("%s is not iaas.IPAddress", op)
	}

	if op, ok := NewIPv6NetOp().(iaas.IPv6NetAPI); !ok {
		t.Fatalf("%s is not iaas.IPv6Net", op)
	}

	if op, ok := NewIPv6AddrOp().(iaas.IPv6AddrAPI); !ok {
		t.Fatalf("%s is not iaas.IPv6Addr", op)
	}

	if op, ok := NewLicenseOp().(iaas.LicenseAPI); !ok {
		t.Fatalf("%s is not iaas.License", op)
	}

	if op, ok := NewLicenseInfoOp().(iaas.LicenseInfoAPI); !ok {
		t.Fatalf("%s is not iaas.LicenseInfo", op)
	}

	if op, ok := NewLoadBalancerOp().(iaas.LoadBalancerAPI); !ok {
		t.Fatalf("%s is not iaas.LoadBalancer", op)
	}

	if op, ok := NewLocalRouterOp().(iaas.LocalRouterAPI); !ok {
		t.Fatalf("%s is not iaas.LocalRouter", op)
	}

	if op, ok := NewMobileGatewayOp().(iaas.MobileGatewayAPI); !ok {
		t.Fatalf("%s is not iaas.MobileGateway", op)
	}

	if op, ok := NewNFSOp().(iaas.NFSAPI); !ok {
		t.Fatalf("%s is not iaas.NFS", op)
	}

	if op, ok := NewNoteOp().(iaas.NoteAPI); !ok {
		t.Fatalf("%s is not iaas.Note", op)
	}

	if op, ok := NewPacketFilterOp().(iaas.PacketFilterAPI); !ok {
		t.Fatalf("%s is not iaas.PacketFilter", op)
	}

	if op, ok := NewPrivateHostOp().(iaas.PrivateHostAPI); !ok {
		t.Fatalf("%s is not iaas.PrivateHost", op)
	}

	if op, ok := NewPrivateHostPlanOp().(iaas.PrivateHostPlanAPI); !ok {
		t.Fatalf("%s is not iaas.PrivateHostPlan", op)
	}

	if op, ok := NewProxyLBOp().(iaas.ProxyLBAPI); !ok {
		t.Fatalf("%s is not iaas.ProxyLB", op)
	}

	if op, ok := NewRegionOp().(iaas.RegionAPI); !ok {
		t.Fatalf("%s is not iaas.Region", op)
	}

	if op, ok := NewServerOp().(iaas.ServerAPI); !ok {
		t.Fatalf("%s is not iaas.Server", op)
	}

	if op, ok := NewServerPlanOp().(iaas.ServerPlanAPI); !ok {
		t.Fatalf("%s is not iaas.ServerPlan", op)
	}

	if op, ok := NewServiceClassOp().(iaas.ServiceClassAPI); !ok {
		t.Fatalf("%s is not iaas.ServiceClass", op)
	}

	if op, ok := NewSIMOp().(iaas.SIMAPI); !ok {
		t.Fatalf("%s is not iaas.SIM", op)
	}

	if op, ok := NewSimpleMonitorOp().(iaas.SimpleMonitorAPI); !ok {
		t.Fatalf("%s is not iaas.SimpleMonitor", op)
	}

	if op, ok := NewSSHKeyOp().(iaas.SSHKeyAPI); !ok {
		t.Fatalf("%s is not iaas.SSHKey", op)
	}

	if op, ok := NewSubnetOp().(iaas.SubnetAPI); !ok {
		t.Fatalf("%s is not iaas.Subnet", op)
	}

	if op, ok := NewSwitchOp().(iaas.SwitchAPI); !ok {
		t.Fatalf("%s is not iaas.Switch", op)
	}

	if op, ok := NewVPCRouterOp().(iaas.VPCRouterAPI); !ok {
		t.Fatalf("%s is not iaas.VPCRouter", op)
	}

	if op, ok := NewWebAccelOp().(iaas.WebAccelAPI); !ok {
		t.Fatalf("%s is not iaas.WebAccel", op)
	}

	if op, ok := NewZoneOp().(iaas.ZoneAPI); !ok {
		t.Fatalf("%s is not iaas.Zone", op)
	}

}