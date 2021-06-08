// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hostmetricsreceiver

import (
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ikrijah2/nudge/component/componenttest"
	"github.com/ikrijah2/nudge/config"
	"github.com/ikrijah2/nudge/config/configtest"
	"github.com/ikrijah2/nudge/internal/processor/filterset"
	"github.com/ikrijah2/nudge/receiver/hostmetricsreceiver/internal"
	"github.com/ikrijah2/nudge/receiver/hostmetricsreceiver/internal/scraper/cpuscraper"
	"github.com/ikrijah2/nudge/receiver/hostmetricsreceiver/internal/scraper/diskscraper"
	"github.com/ikrijah2/nudge/receiver/hostmetricsreceiver/internal/scraper/filesystemscraper"
	"github.com/ikrijah2/nudge/receiver/hostmetricsreceiver/internal/scraper/loadscraper"
	"github.com/ikrijah2/nudge/receiver/hostmetricsreceiver/internal/scraper/memoryscraper"
	"github.com/ikrijah2/nudge/receiver/hostmetricsreceiver/internal/scraper/networkscraper"
	"github.com/ikrijah2/nudge/receiver/hostmetricsreceiver/internal/scraper/pagingscraper"
	"github.com/ikrijah2/nudge/receiver/hostmetricsreceiver/internal/scraper/processesscraper"
	"github.com/ikrijah2/nudge/receiver/hostmetricsreceiver/internal/scraper/processscraper"
	"github.com/ikrijah2/nudge/receiver/scraperhelper"
)

func TestLoadConfig(t *testing.T) {
	factories, err := componenttest.NopFactories()
	require.NoError(t, err)

	factory := NewFactory()
	factories.Receivers[typeStr] = factory
	cfg, err := configtest.LoadConfigFile(t, path.Join(".", "testdata", "config.yaml"), factories)

	require.NoError(t, err)
	require.NotNil(t, cfg)

	assert.Equal(t, len(cfg.Receivers), 2)

	r0 := cfg.Receivers["hostmetrics"]
	defaultConfigCPUScraper := factory.CreateDefaultConfig()
	defaultConfigCPUScraper.(*Config).Scrapers = map[string]internal.Config{
		cpuscraper.TypeStr: (&cpuscraper.Factory{}).CreateDefaultConfig(),
	}

	assert.Equal(t, defaultConfigCPUScraper, r0)

	r1 := cfg.Receivers["hostmetrics/customname"].(*Config)
	expectedConfig := &Config{
		ScraperControllerSettings: scraperhelper.ScraperControllerSettings{
			ReceiverSettings: config.ReceiverSettings{
				TypeVal: typeStr,
				NameVal: "hostmetrics/customname",
			},
			CollectionInterval: 30 * time.Second,
		},
		Scrapers: map[string]internal.Config{
			cpuscraper.TypeStr:        &cpuscraper.Config{},
			diskscraper.TypeStr:       &diskscraper.Config{},
			loadscraper.TypeStr:       &loadscraper.Config{},
			filesystemscraper.TypeStr: &filesystemscraper.Config{},
			memoryscraper.TypeStr:     &memoryscraper.Config{},
			networkscraper.TypeStr: &networkscraper.Config{
				Include: networkscraper.MatchConfig{
					Interfaces: []string{"test1"},
					Config:     filterset.Config{MatchType: "strict"},
				},
			},
			processesscraper.TypeStr: &processesscraper.Config{},
			pagingscraper.TypeStr:    &pagingscraper.Config{},
			processscraper.TypeStr: &processscraper.Config{
				Include: processscraper.MatchConfig{
					Names:  []string{"test2", "test3"},
					Config: filterset.Config{MatchType: "regexp"},
				},
			},
		},
	}

	assert.Equal(t, expectedConfig, r1)
}

func TestLoadInvalidConfig_NoScrapers(t *testing.T) {
	factories, err := componenttest.NopFactories()
	require.NoError(t, err)

	factory := NewFactory()
	factories.Receivers[typeStr] = factory
	_, err = configtest.LoadConfigFile(t, path.Join(".", "testdata", "config-noscrapers.yaml"), factories)

	require.EqualError(t, err, "receiver \"hostmetrics\" has invalid configuration: must specify at least one scraper when using hostmetrics receiver")
}

func TestLoadInvalidConfig_InvalidScraperKey(t *testing.T) {
	factories, err := componenttest.NopFactories()
	require.NoError(t, err)

	factory := NewFactory()
	factories.Receivers[typeStr] = factory
	_, err = configtest.LoadConfigFile(t, path.Join(".", "testdata", "config-invalidscraperkey.yaml"), factories)

	require.EqualError(t, err, "error reading receivers configuration for hostmetrics: invalid scraper key: invalidscraperkey")
}
