package main

type CloudflareDNSResponse struct {
	Result   DNSRecord `json:"result"`
	Success  bool      `json:"success"`
	Errors   []string  `json:"errors"`
	Messages []string  `json:"messages"`
}

type DNSRecord struct {
	ID         string                 `json:"id"`
	ZoneID     string                 `json:"zone_id"`
	ZoneName   string                 `json:"zone_name"`
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	Content    string                 `json:"content"`
	Proxiable  bool                   `json:"proxiable"`
	Proxied    bool                   `json:"proxied"`
	TTL        int                    `json:"ttl"`
	Settings   map[string]interface{} `json:"settings"`
	Meta       Meta                   `json:"meta"`
	Comment    *string                `json:"comment"`
	Tags       []string               `json:"tags"`
	CreatedOn  string                 `json:"created_on"`
	ModifiedOn string                 `json:"modified_on"`
}

type Meta struct {
	AutoAdded           bool `json:"auto_added"`
	ManagedByApps       bool `json:"managed_by_apps"`
	ManagedByArgoTunnel bool `json:"managed_by_argo_tunnel"`
}
