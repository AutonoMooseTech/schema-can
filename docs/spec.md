# Specification

SchemaCAN is a protocol designed to replace the DBC format by making improvements in key areas.

## Understanding SchemaCAN Objects

SchemaCAN objects are entities which describe the the properties of CAN bus data in an open format which is easy to parse for both computers and humans.

## Object Metadata

These fields describe the object itself with some attributes like `name` being required to be parsed correctly.

- `version` Schema version to allow tooling to recognise.
- `kind` What _kind_ of object is being declared.
- `metadata`
  - `name` - Descriptive name to identify the object which must be inique withing it's type and namespace.
  - `namespace` - Optional scope of the object. If not set, the namespace value is 'default'.
  - `labels` - Optional key-value tags that can be used to add extra unstructured information to the object.

#### Example Common Field Definition

<CodeGroup>
  <CodeGroupItem title="YAML">

```yaml
version: v1
kind: signal
metadata:
  name: gear-position
  namespace: transmission
  labels:
    manual_revision: 1.3.0
```

  </CodeGroupItem>
  <CodeGroupItem title="JSON">

```json
{
  "version": "v1",
  "kind": "signal",
  "metadata": {
    "name": "gear-position",
    "namespace": "transmission",
    "labels": {
      "manual_revision": "1.3.0"
    }
  }
}
```

  </CodeGroupItem>
</CodeGroup>

## Signals

Signals object define how single pieces of data are encoded and decoded. Signals can either follow a primitive type or specified scaling, limit, offset, transfer function (SLOT).

### Primitive Types

#### Integer

Signed and unsigned integers of any size between 1 and 64 bits can be used as primitive data types using the notation `u[n]` for unsigned and `i[n]` for signed.

Here are some examples

```
u8: unsigned 8-bit integer
u3: unsigned 3-bit integer
i15: signed 15-bit integer
i64: signed 64-bit integer
```

#### Floating-Point

::: warning Avoid the use of floating point types where possible
Support for floating point types is for reasons of backwards-compatibility. Generally it is regarded as best practice to avoid floating point numbers with CAN bus.
:::

Floating point numbers must be defined using one of the following types.

| Type | C Equivalent | Description                                    |
|:-----|:-------------|:-----------------------------------------------|
| f16  | _Float16     | 16-bit floating point (IEEE-754-2008 binary16) |
| f32  | float        | 32-bit floating point (IEEE-754-2008 binary32) |
| f64  | double       | 64-bit floating point (IEEE-754-2008 binary64) |

#### Arrays and Strings

Arrays are fixed length collections of the same primitive type. They can be defined using the pattern `type[length]` and will occupy the number of bits equal to the product of the length of the type and the length of the array.

<CodeGroup>
  <CodeGroupItem title="YAML">

```yaml
- name: quadrants
  type: u6[4]
```

  </CodeGroupItem>
  <CodeGroupItem title="JSON">

```json
{
  "name": "quadrants",
  "type": "u6[4]"
}
```

  </CodeGroupItem>
</CodeGroup>

### SLOTs

SLOT stands for scaling, limit, offset and transfer function which is a concept from the [SAE J1939-71](https://www.sae.org/standards/content/j1939/71_202002/) standard to largely sove the problem of encoding and decofing between foatping point and fixed point numbers.

A slot can either be defined on its own to be used by referend by a message or it can be defined in place in the message data field if it's not going to be needed anywhere else. If possible, it's a good idea to use the pre-defined SLOTs in the [J1939 Digital Annex](https://www.sae.org/standards/content/j1939da_202201/).

#### Example SLOT
<CodeGroup>
  <CodeGroupItem title="YAML">

```yaml
version: v1
kind: slot
metadata:
  name: sae-ev-06
  namespace: j1939
spec:
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
    "min": 0,
    "max": 64.255,
    "size": 16,
    "unit": "V"
  }
}
```

  </CodeGroupItem>
</CodeGroup>

## Messages

#### Example Message
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



