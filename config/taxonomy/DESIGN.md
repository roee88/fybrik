# Taxonomy

The project defines a set of immutable structural JSON schemas, e.g., for catalog and policy manager request and response. These schemas are considered part of the project API and cannot be customized directly. However, since the taxonomy is meant to be configurable, a `taxonomy.json` file is referenced from these schemas for any definition that is customizable.

The `taxonomy.json` file is generated from a base taxonomy file and zero or more taxonomy layer files. The base taxonomy is maintained by the project and includes all of the structural definitions that are subject to customization (e.g.: tags, actions). The taxonomy layers are maintained by users and external systems that add customizations over the base taxonomy (e.g., defining specific tags, actions).

A go library and a command line utility to generate `taxonomy.json` from base and layers is available.

## Examples for base and layers

If a document is valid against a compiled taxonomy then it MUST also be valid against the base taxonomy (the other direction is not always true). To help write good base and layer taxonomy files that preserve this property, an overview of typical types and how their base and layers may look like is given below. You don't need to declare the type of a structure as part of a definition as these are just guidelines for common types that we think are useful.

### Union type

A union type defines an object that can take one of multiple structures. This can be useful for defining new enforcement actions and new connection types.

The base structure of a union is an object with just a `name` field and `additionalProperties: true`:

```yaml
definitions:
  MyUnion:
    type: object
    properties:
      name:
        type: string
    additionalProperties: true
    required: [name]
```

It can be extended by adding `oneOf` in a layer. For example:

```yaml
definitions:
  MyUnion:
    oneOf:
    - $ref: "#/definitions/MyUnionSubtype1"
    - $ref: "#/definitions/MyUnionSubtype2"
  MyUnionSubtype1:
    type: object
    properties:
      custom:
        type: string
    additionalProperties: false
  MyUnionSubtype2:
    ...
```

The `taxonomy compile` command will generate an output that looks the following:

```yaml
definitions:
  MyUnion:
    type: object
    properties:
      name:
        type: string
        enum: [MyUnionSubtype1, MyUnionSubtype2]
      MyUnionSubtype1:
        $ref: "#/definitions/MyUnionSubtype1"
      MyUnionSubtype2:
        $ref: "#/definitions/MyUnionSubtype2"
    oneOf:
    - properties:
        name:
            enum:
            - MyUnionSubtype1
        required:
        - name
        - MyUnionSubtype1
    - properties:
        name:
            enum:
            - MyUnionSubtype2
        required:
        - name
        - MyUnionSubtype2
    additionalProperties: false
  MyUnionSubtype1:
    type: object
    properties:
      custom:
        type: string
    additionalProperties: false
  MyUnionSubtype2:
    ...
```

### Map type

A map is an mapping between string keys and some typed values. The base definition is an empty object with `additionalProperties: true` which represents an arbitrary map to any value types:

```yaml
definitions:
  MyMap:
    type: object
    additionalProperties: true
```

A layer can then be added with explicit peroperties on top of a base empty structure. It can also override `additionalProperties` to forbid unknowns. For example, the following layer only allows a key `someProperty` holding string values:

```yaml
definitions:
  MyMap:
    properties:
      someProperty:
        type: string
    additionalProperties: false
```

This can be useful for example to definition of tags and runtime context attributes.

### Enum type

For enum the base structure is a primitive type, typically a `string`:

```yaml
definitions:
  MyEnum:
    type: string
```

This can then be customized in layers, for example:

```yaml
definitions:
  MyEnum:
    enum: [value1, value2, value3]
```

You SHOULD define a separate type for any base enum because it's useful for code generation. Enums are also customized automatically by the compiler from validation properties as described next.


### Immutable types

An immutable type is a type that has a fixed structure that should not be changed by layers. However, layers can add validation properties on top of the immutable structure.

For immutables the base structure is any structural definition excluding any _custom_ validation. For example:

```yaml
definitions:
  MyImmutable:
    type: object
    properties:
      propA:
        type: string
      propB:
        type: integer
        format: int32
    additionalProperties: false
```

Customization adds `oneOf` or `anyOf` validation over the base definition, but it must not change the structure  (hence "immutable"). This is useful for example to interface or geography validation.

Example base structure:

```yaml
definitions:
  Protocol: # Enum
    type: string
  Format: # Enum
    type: string
  Interface: # Immutable
    type: object
    properties:
      protocol:
        $ref: "#/definitions/Protocol"
      format:
        $ref: "#/definitions/Format"
    additionalProperties: false
```

Example validation:

```yaml
definitions:
  Interface:
    oneOf:
      - properties:
          protocol:
            enum: [https]
          format:
            enum: [json]
      - properties:
          protocol:
            enum: [grpc]
          format:
            enum: [proto]
```

The output from `taxonomy compile`:

```yaml
definitions:
  Protocol:
    type: string
    enum:
      - https
      - grpc
  Format:
    enum:
      - json
      - proto
  Interface:
    type: object
    properties:
      protocol:
        $ref: "#/definitions/Protocol"
      format:
        $ref: "#/definitions/Format"
    oneOf:
      - properties:
          protocol:
            enum:
              - https
          format:
            enum:
              - json
      - properties:
          protocol:
            enum:
              - grpc
          format:
            enum:
              - proto
    additionalProperties: false
```
