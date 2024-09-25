package cfg

import "github.com/flevin58/fin/tools"

func AddApp(app string) {
	Apps = append(Apps, app)
}

func AddApps(apps ...string) {
	Apps = append(Apps, apps...)
}

func RemoveApp(app string) {
	i, found := tools.FindIndexOf(Apps, app)
	if found {
		Apps = tools.RemoveAtIndex(Apps, i)
	}
}

func RemoveApps(apps ...string) {
	for _, app := range apps {
		RemoveApp(app)
	}
}
