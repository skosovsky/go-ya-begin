package main

var composeOut = map[string]interface{}{
	"services": map[string]interface{}{
		"backend": map[string]interface{}{
			"build": map[string]interface{}{
				"context": "backend",
				"target":  "builder",
			},
		},
		"proxy": map[string]interface{}{
			"depends_on": []interface{}{
				"backend",
			},
			"image": "nginx",
			"ports": []interface{}{
				"80:80",
			},
			"volumes": []interface{}{
				map[string]interface{}{
					"read_only": true,
					"source":    "./proxy/nginx.conf",
					"target":    "/etc/nginx/conf.d/default.conf",
					"type":      "bind",
				},
			},
		},
	},
}
