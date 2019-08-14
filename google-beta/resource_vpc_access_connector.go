// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceVpcAccessConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpcAccessConnectorCreate,
		Read:   resourceVpcAccessConnectorRead,
		Delete: resourceVpcAccessConnectorDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVpcAccessConnectorImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(6 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"ip_cidr_range": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"max_throughput": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(200, 1000),
				Default:      1000,
			},
			"min_throughput": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(200, 1000),
				Default:      200,
			},
			"network": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceVpcAccessConnectorCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandVpcAccessConnectorName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandVpcAccessConnectorNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	ipCidrRangeProp, err := expandVpcAccessConnectorIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); !isEmptyValue(reflect.ValueOf(ipCidrRangeProp)) && (ok || !reflect.DeepEqual(v, ipCidrRangeProp)) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}
	minThroughputProp, err := expandVpcAccessConnectorMinThroughput(d.Get("min_throughput"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("min_throughput"); !isEmptyValue(reflect.ValueOf(minThroughputProp)) && (ok || !reflect.DeepEqual(v, minThroughputProp)) {
		obj["minThroughput"] = minThroughputProp
	}
	maxThroughputProp, err := expandVpcAccessConnectorMaxThroughput(d.Get("max_throughput"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("max_throughput"); !isEmptyValue(reflect.ValueOf(maxThroughputProp)) && (ok || !reflect.DeepEqual(v, maxThroughputProp)) {
		obj["maxThroughput"] = maxThroughputProp
	}

	obj, err = resourceVpcAccessConnectorEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{VpcAccessBasePath}}projects/{{project}}/locations/{{region}}/connectors?connectorId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Connector: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Connector: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	waitErr := vpcAccessOperationWaitTime(
		config, res, project, "Creating Connector",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Connector: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating Connector %q: %#v", d.Id(), res)

	// This is useful if the resource in question doesn't have a perfectly consistent API
	// That is, the Operation for Create might return before the Get operation shows the
	// completed state of the resource.
	time.Sleep(5 * time.Second)

	return resourceVpcAccessConnectorRead(d, meta)
}

func resourceVpcAccessConnectorRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{VpcAccessBasePath}}projects/{{project}}/locations/{{region}}/connectors/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("VpcAccessConnector %q", d.Id()))
	}

	res, err = resourceVpcAccessConnectorDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing VpcAccessConnector because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Connector: %s", err)
	}

	if err := d.Set("name", flattenVpcAccessConnectorName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading Connector: %s", err)
	}
	if err := d.Set("network", flattenVpcAccessConnectorNetwork(res["network"], d)); err != nil {
		return fmt.Errorf("Error reading Connector: %s", err)
	}
	if err := d.Set("ip_cidr_range", flattenVpcAccessConnectorIpCidrRange(res["ipCidrRange"], d)); err != nil {
		return fmt.Errorf("Error reading Connector: %s", err)
	}
	if err := d.Set("state", flattenVpcAccessConnectorState(res["state"], d)); err != nil {
		return fmt.Errorf("Error reading Connector: %s", err)
	}
	if err := d.Set("min_throughput", flattenVpcAccessConnectorMinThroughput(res["minThroughput"], d)); err != nil {
		return fmt.Errorf("Error reading Connector: %s", err)
	}
	if err := d.Set("max_throughput", flattenVpcAccessConnectorMaxThroughput(res["maxThroughput"], d)); err != nil {
		return fmt.Errorf("Error reading Connector: %s", err)
	}

	return nil
}

func resourceVpcAccessConnectorDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{VpcAccessBasePath}}projects/{{project}}/locations/{{region}}/connectors/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Connector %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Connector")
	}

	err = vpcAccessOperationWaitTime(
		config, res, project, "Deleting Connector",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Connector %q: %#v", d.Id(), res)
	return nil
}

func resourceVpcAccessConnectorImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/connectors/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenVpcAccessConnectorName(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenVpcAccessConnectorNetwork(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenVpcAccessConnectorIpCidrRange(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenVpcAccessConnectorState(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenVpcAccessConnectorMinThroughput(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenVpcAccessConnectorMaxThroughput(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func expandVpcAccessConnectorName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVpcAccessConnectorNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVpcAccessConnectorIpCidrRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVpcAccessConnectorMinThroughput(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVpcAccessConnectorMaxThroughput(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceVpcAccessConnectorEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "name")
	return obj, nil
}

func resourceVpcAccessConnectorDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// Take the returned long form of the name and use it as `self_link`.
	// Then modify the name to be the user specified form.
	// We can't just ignore_read on `name` as the linter will
	// complain that the returned `res` is never used afterwards.
	// Some field needs to be actually set, and we chose `name`.
	d.Set("self_link", res["name"].(string))
	res["name"] = d.Get("name").(string)
	return res, nil
}