# Specification

## Common Fields

These fields are present in every object and should be placed at the top of each object definition

- `version` Object schema version
- `kind` Object type
- `metadata`
	- `name` - 
	- `namespace` - 
	- `labels` - user-defined tags that can be used to add extra unstructured information to the object

## Object Types

### Message

::: details Example
<CodeGroup>
  <CodeGroupItem title="YAML">

```yaml
version: v1
kind: message
metadata:
  name: controller-status
  namespace: my-battery
spec:
  id:
    extended: 0x555
  data:
    - name: enabled
      description: set when the battery is turned on
      size: bool
    - name: voltage-ok
      description: voltage is within pre-defined limits
      size: bool
    - padding: 6
    - name: voltage
      description: total battery voltage
      slot: sae-ev-06
```

  </CodeGroupItem>
  <CodeGroupItem title="JSON">

```json
{
	"version": "v1",
	"kind": "message",
	"metadata": {
		"name": "controller-status",
		"namespace": "my-battery"
	},
	"spec": {
		"id": {
			"extended": "0x555"
		},
		"data": [
			{
				"name": "enabled",
				"description": "set when the battery is turned on",
				"size": "bool"
			},
			{
				"name": "voltage-ok",
				"description": "voltage is within pre-defined limits",
				"size": "bool"
			},
			{
				"padding": 6
			},
			{
				"name": "voltage",
				"description": "total battery voltage",
				"slot": "sae-ev-06"
			}
		]
	}
}
```

  </CodeGroupItem>
</CodeGroup>
:::

### SLOT

The scaling, limit, offset, transfer function is a concept that is brought forward by the SAE J1939 standard and it largely solves the problem of encoding and decoding between floating point and fixed point numbers. As well as being able to use custom SLOT definitions, SchemaCAN comes pre-packaged with the SLOTS defined in SAE J1939.

::: details Example
<CodeGroup>
  <CodeGroupItem title="YAML">

```yaml
version: v1
kind: slot
metadata:
  name: sae-ev-06
  namespace: j1939
spec:
  limits:
  	min: 0
  	max: 64.255 # in units
  size: 16 # in bits
  unit: V
```

  </CodeGroupItem>
  <CodeGroupItem title="JSON">

```json
{
	"version": "v1",
	"kind": "slot",
	"metadata": {
		"name": "sae-ev-06",
		"namespace": "my-battery"
	},
	"spec": {
		"scale": 0.001,
		"limits": {
			"min": 0,
			"max": 64.255
		},
		"size": 16,
		"unit": "V"
	}
}
```

  </CodeGroupItem>
</CodeGroup>
:::
