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
	"fmt"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
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
	return []plugin.MetricType{}, fmt.Errorf("Not implemented")
}

// GetMetricTypes returns metric types for testing
func (f *Weekend) GetMetricTypes(cfg plugin.ConfigType) ([]plugin.MetricType, error) {
	return []plugin.MetricType{}, fmt.Errorf("Not implemented")

}

// GetConfigPolicy returns a ConfigPolicy for testing
func (f *Weekend) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	return nil, fmt.Errorf("Not implemented")
}

// Meta returns meta data for testing
func Meta() *plugin.PluginMeta {
	return nil
}
