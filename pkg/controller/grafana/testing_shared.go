package grafana

import "github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1"

var Mockplugina100 = v1alpha1.GrafanaPlugin{
	Name:    "a",
	Version: "1.0.0",
}

var Mockplugina101 = v1alpha1.GrafanaPlugin{
	Name:    "a",
	Version: "1.0.1",
}

var Mockplugina102 = v1alpha1.GrafanaPlugin{
	Name:    "a",
	Version: "1.0.2",
}

var Mockpluginb100 = v1alpha1.GrafanaPlugin{
	Name:    "b",
	Version: "1.0.0",
}

var Mockpluginc100 = v1alpha1.GrafanaPlugin{
	Name:    "c",
	Version: "1.0.0",
}

var MockPluginList = v1alpha1.PluginList{Mockplugina100, Mockplugina101, Mockpluginb100}

var MockDashboard = v1alpha1.GrafanaDashboard{
	Status: v1alpha1.GrafanaDashboardStatus{
		Messages: []v1alpha1.GrafanaDashboardStatusMessage{},
	},
}

var MockGrafana = v1alpha1.Grafana{
	Status: v1alpha1.GrafanaStatus{
		Phase:            0,
		InstalledPlugins: v1alpha1.PluginList{},
	},
}
