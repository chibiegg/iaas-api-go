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

package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/fake"
	"github.com/sacloud/iaas-api-go/helper/defaults"
	"github.com/sacloud/iaas-api-go/trace"
	"github.com/sacloud/iaas-api-go/trace/otel"
	"github.com/sacloud/sacloud-go/client"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewCaller() (iaas.APICaller, error) {
	clientOpts, err := client.DefaultOption()
	if err != nil {
		return nil, err
	}
	return NewCallerWithOptions(&CallerOptions{Options: clientOpts}), nil
}

// NewCallerWithOptions 指定のオプションでiaas.APICallerを構築して返す
func NewCallerWithOptions(opts *CallerOptions) iaas.APICaller {
	return newCaller(opts)
}

func newCaller(opts *CallerOptions) iaas.APICaller {
	if opts.UserAgent == "" {
		opts.UserAgent = iaas.DefaultUserAgent
	}

	caller := iaas.NewClientWithOptions(opts.Options)

	iaas.DefaultStatePollingTimeout = 72 * time.Hour

	if opts.TraceAPI {
		// note: exact once
		trace.AddClientFactoryHooks()
	}

	if opts.OpenTelemetry {
		otel.Initialize(opts.OpenTelemetryOptions...)
		httpClient := http.DefaultClient
		if opts.HttpClient != nil {
			httpClient = opts.HttpClient
		}
		transport := httpClient.Transport
		if transport == nil {
			transport = http.DefaultTransport
		}
		httpClient.Transport = otelhttp.NewTransport(transport)
	}

	if opts.FakeMode {
		if opts.FakeStorePath != "" {
			fake.DataStore = fake.NewJSONFileStore(opts.FakeStorePath)
		}
		// note: exact once
		fake.SwitchFactoryFuncToFake()

		SetupFakeDefaults()
	}

	if opts.DefaultZone != "" {
		iaas.APIDefaultZone = opts.DefaultZone
	}

	if len(opts.Zones) > 0 {
		iaas.SakuraCloudZones = opts.Zones
	}

	if opts.APIRootURL != "" {
		if strings.HasSuffix(opts.APIRootURL, "/") {
			opts.APIRootURL = strings.TrimRight(opts.APIRootURL, "/")
		}
		iaas.SakuraCloudAPIRoot = opts.APIRootURL
	}

	return caller
}

func SetupFakeDefaults() {
	defaultInterval := 10 * time.Millisecond

	// update default polling intervals: libsacloud/sacloud
	iaas.DefaultStatePollingInterval = defaultInterval
	iaas.DefaultDBStatusPollingInterval = defaultInterval
	// update default polling intervals: libsacloud/helper/setup
	defaults.DefaultDeleteWaitInterval = defaultInterval
	defaults.DefaultProvisioningWaitInterval = defaultInterval
	defaults.DefaultPollingInterval = defaultInterval
	// update default polling intervals: libsacloud/helper/builder
	defaults.DefaultNICUpdateWaitDuration = defaultInterval
	// update default timeouts and span: libsacloud/helper/power
	defaults.DefaultPowerHelperBootRetrySpan = defaultInterval
	defaults.DefaultPowerHelperShutdownRetrySpan = defaultInterval
	defaults.DefaultPowerHelperInitialRequestRetrySpan = defaultInterval
	defaults.DefaultPowerHelperInitialRequestTimeout = defaultInterval * 100

	fake.PowerOnDuration = time.Millisecond
	fake.PowerOffDuration = time.Millisecond
	fake.DiskCopyDuration = time.Millisecond
}