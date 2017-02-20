/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package weekend

import (
	"log"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
)

const (
	// Name of plugin
	Name = "weekend"
	// Version of plugin
	Version = 1
	// Type of plugin
	Type = plugin.CollectorPluginType
)

// Mock collector implementation used for testing
type Weekend struct {
}

// CollectMetrics collects metrics for testing
func (f *Weekend) CollectMetrics(mts []plugin.MetricType) ([]plugin.MetricType, error) {
	for _, p := range mts {
		log.Printf("collecting %+v\n", p)
	}

	metrics := []plugin.MetricType{}
	for i := range mts {
		switch mts[i].Namespace()[2].Value {
		case "daystillfriday":
			mts[i].Data_ = 1
		case "isitweekend":
			mts[i].Data_ = "false"
		}

		mts[i].Timestamp_ = time.Now()
		metrics = append(metrics, mts[i])
	}

	return metrics, nil
}

// GetMetricTypes returns metric types for testing
func (f *Weekend) GetMetricTypes(cfg plugin.ConfigType) ([]plugin.MetricType, error) {
	mts := []plugin.MetricType{}

	mts = append(mts, plugin.MetricType{
		Namespace_:   core.NewNamespace("jjlakis", "weekend", "daystillfriday"),
		Description_: "Days till friday",
		Unit_:        "Days",
	})
	mts = append(mts, plugin.MetricType{
		Namespace_:   core.NewNamespace("jjlakis", "weekend", "isitweekend"),
		Description_: "Is it weekend (Friday, Saturday, Sunday)?",
		Unit_:        "True/false",
	})
	return mts, nil

}

// GetConfigPolicy returns a ConfigPolicy for testing
func (f *Weekend) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	return cpolicy.New(), nil
}

// Meta returns meta data for testing
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(
		Name,
		Version,
		Type,
		[]string{plugin.SnapGOBContentType},
		[]string{plugin.SnapGOBContentType},
		plugin.CacheTTL(100*time.Millisecond),
		plugin.RoutingStrategy(plugin.StickyRouting),
	)
}
