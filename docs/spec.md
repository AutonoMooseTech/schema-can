# Specification

## Concepts

The idea is that a system of communication between two or more devices on a CAN Bus can be defined as a heirachy of objects which conform to a certain schema. This information can then be consumed by applications either on the bus or as part of external monitoring to allow them to understand the content of messages which would otherwise be nothing more than streams of bytes.

At the lowest level there are signals, the actual pieces of data which are packed together to make up the data field of a message. These messages have attributes like length and identifier(s) which gives tooling all the information needed to identify a CAN frame and either decode or encode it's contents. Applications are the next level up and are a collection of commands (received by the controller) or broadcast (sent by the controller) messages which define the interface for that application.

It's important to make the distinction that SchemaCAN is not a method of sending data over CAN Bus or otherwise. It meerly a way to ensure messages are encoded and decoded in a way that information is not lost. It can however be used as a tool to help bridge CAN Bus (including SocketCAN) to another communication protocol.


## Common Fields

These fields are present in every object and should be placed at the top of each object definition

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

Signals should be seen at the atomic pieces of data, that is they cannot be subdevided any further and SchemaCAN tries to enforce this as much as possible by trying to facilitate all of the possible types of data even if they do not conform to best practices.

There are two types of signals: primitive and SLOTs. Primitive signals do not provide any information about how the data is to be interpreted in the real world. They are either purely analogous to their type or it is not possible to represent them using SLOTs. SLOTs on the other hand is a way of letting both system and humans understand the meaning of the signal in real world terms by allowing predictive encoding and decoding of information.

### Primitive

Data types tell the associated tooling two things, the ammount of space occupied by the

| Type | C Equivalent | Description                                    |
|:-----|:-------------|:-----------------------------------------------|
| u8   | uint8_t      | 8-bit unsigned integer                         |
| i8   | int8_t       | 8-bit signed integer                           |
| u16  | uint16_t     | 16-bit unsigned integer                        |
| i16  | int16_t      | 16-bit signed integer                          |
| u32  | uint32_t     | 32-bit unsigned integer                        |
| i32  | int32_t      | 32-bit signed integer                          |
| u64  | uint64_t     | 64-bit unsigned integer                        |
| i64  | int64_t      | 64-bit signed integer                          |
| uN   |              | N-bit unsigned integer                         |
| iN   |              | N-bit signed integer                           |
| f16  | _Float16     | 16-bit floating point (IEEE-754-2008 binary16) |
| f32  | float        | 32-bit floating point (IEEE-754-2008 binary32) |
| f64  | double       | 64-bit floating point (IEEE-754-2008 binary64) |

::: warning Avoid the use of floats where possible
Support for floating point numbers is for reasons of compatibility with existing systems. Generally it is regarded as best practice to avoid the use of floating point numbers as they are not space efficient.
:::

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



