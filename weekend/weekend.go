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
			mts[i].Data_ = (5 - int(time.Now().Weekday()) + 7) % 7
		case "isitweekend":
			mts[i].Data_ = time.Now().Weekday().String() == "Friday"
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
