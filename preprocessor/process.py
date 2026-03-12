import json as j
import sys
import copy

def schema_from_ref(ref):
    return ref.split('/')[-1]

def get_enum_value_from_connection_spec(connection_spec):
    schemas = connection_spec['allOf']
    for schema in schemas:
        if '$ref' in schema:
            continue  # this is a reference, skip
        if 'properties' in schema:
            return schema['properties']['type']['enum'][0]
    return ''

def updateConnectionTypesRef(schema, oldReference, newReference):
    newSchema = {}
    for key in schema:
        if isinstance(schema[key], dict):
            newSchema[key] = updateConnectionTypesRef(schema[key], oldReference, newReference)
        elif key == '$ref' and schema[key] == oldReference:
             newSchema[key] = newReference
        else:
            newSchema[key] = schema[key]

    return newSchema

def rename_schema_keys(schemas, old_prefix):
    # Remove prefix and convert first character to lowercase
    renamed_schemas = {}
    for key, value in schemas.items():
        if key.startswith(old_prefix):
            new_key = key[len(old_prefix):]
            if new_key:
                new_key = new_key[0].lower() + new_key[1:]
            renamed_schemas[new_key] = value
        else:
            renamed_schemas[key] = value
    return renamed_schemas

def update_all_refs(obj, old_prefix):
    # Remove prefix and convert to lowercase for $refs
    if isinstance(obj, dict):
        new_obj = {}
        for key, value in obj.items():
            if key == '$ref' and isinstance(value, str):
                # Update schema references that use the old prefix
                if f'schemas/{old_prefix}' in value:
                    parts = value.split(f'schemas/{old_prefix}')
                    if len(parts) == 2:
                        schema_name = parts[1]
                        if schema_name:
                            schema_name = schema_name[0].lower() + schema_name[1:]
                        new_obj[key] = parts[0] + 'schemas/' + schema_name
                    else:
                        new_obj[key] = value
                else:
                    new_obj[key] = value
            else:
                new_obj[key] = update_all_refs(value, old_prefix)
        return new_obj
    elif isinstance(obj, list):
        return [update_all_refs(item, old_prefix) for item in obj]
    else:
        return obj

def fix_discriminator_mappings(schemas, schema_name):
    if schema_name not in schemas:
        return
    schema = schemas[schema_name]
    if 'discriminator' not in schema:
        return
    discriminator = schema['discriminator']
    if 'mapping' not in discriminator:
        return
    # Determine the types schema name
    if 'source' in schema_name.lower():
        types_schema_name = 'source_types'
    elif 'connection' in schema_name.lower():
        types_schema_name = 'connection_types'
    else:
        return
    if types_schema_name not in schemas:
        return
    types_schema = schemas[types_schema_name]
    if 'oneOf' not in types_schema:
        return

    # Build a map of enum values to their correct schema references
    enum_to_schema = {}
    for item in types_schema['oneOf']:
        if '$ref' not in item:
            continue
        ref = item['$ref']
        if '#/components/schemas/' in ref:
            schema_ref = ref.split('#/components/schemas/')[-1]
            if schema_ref in schemas:
                # Get the enum value from the schema
                enum_value = get_enum_value_from_connection_spec(schemas[schema_ref])
                if enum_value:
                    enum_to_schema[enum_value] = f"#/components/schemas/{schema_ref}"

    # Update the discriminator mapping with correct references
    for enum_key, ref_value in discriminator['mapping'].items():
        if enum_key in enum_to_schema:
            discriminator['mapping'][enum_key] = enum_to_schema[enum_key]

if len(sys.argv) < 3:
    print("Usage python process.py input-schema.json output-schema.json")
    sys.exit(1)

inputSchemaFile = sys.argv[1]
outputSchemaFile = sys.argv[2]

with open(inputSchemaFile, 'r+') as f:
    api_spec = j.load(f)

# Removes prefix for all schema keys of connectors spec file
OLD_PREFIX = 'Etleap-api-connectors.v2_'
api_spec['components']['schemas'] = rename_schema_keys(
    api_spec['components']['schemas'],
    OLD_PREFIX
)

# Update all $ref values throughout the entire spec
api_spec = update_all_refs(api_spec, OLD_PREFIX)

connection_types = api_spec['components']['schemas']['connection_types']['oneOf']
connection_types_update = api_spec['components']['schemas']['connection_update_types']['oneOf']

paths = api_spec['paths']
schemas = api_spec['components']['schemas']

CONNECTIONS_PATH = '/connections'
CONNECTIONS_PATH_WITH_FRAGMENT = CONNECTIONS_PATH + '#'
CONNECTIONS_ID_PATH = '/connections/{id}'
CONNECTIONS_ID_PATH_WITH_FRAGMENT = CONNECTIONS_ID_PATH + '#'

POST = 'post'
GET = 'get'
PATCH = 'patch'
DELETE = 'delete'

for connection_type_update in connection_types_update:
    update_connection_schema = schema_from_ref(connection_type_update['$ref'])
    enum_value = get_enum_value_from_connection_spec(schemas[update_connection_schema])

    paths[CONNECTIONS_ID_PATH_WITH_FRAGMENT + enum_value] = {}
    paths[CONNECTIONS_ID_PATH_WITH_FRAGMENT + enum_value][PATCH] = updateConnectionTypesRef(paths[CONNECTIONS_ID_PATH][PATCH], '#/components/schemas/connection_update_types', connection_type_update['$ref'])
    paths[CONNECTIONS_ID_PATH_WITH_FRAGMENT + enum_value][PATCH]['operationId'] = 'update-' + enum_value + '-connection'
    paths[CONNECTIONS_ID_PATH_WITH_FRAGMENT + enum_value][PATCH]['x-speakeasy-entity-operation'] = 'Connection' + enum_value + '#update'

    schemas[update_connection_schema]['x-speakeasy-entity'] = 'Connection' + enum_value
    schemas[update_connection_schema]['x-speakeasy-param-suppress-computed-diff'] = 'true'

for connection_type in connection_types:
    connections_schema = schema_from_ref(connection_type['$ref'])

    enum_value = get_enum_value_from_connection_spec(schemas[connections_schema])

    paths[CONNECTIONS_PATH_WITH_FRAGMENT + enum_value] = {}
    paths[CONNECTIONS_PATH_WITH_FRAGMENT + enum_value][POST] = updateConnectionTypesRef(paths[CONNECTIONS_PATH][POST], '#/components/schemas/connection_types', connection_type['$ref'])
    paths[CONNECTIONS_PATH_WITH_FRAGMENT + enum_value][POST]['x-speakeasy-entity-operation'] = 'Connection' + enum_value + '#create'
    paths[CONNECTIONS_PATH_WITH_FRAGMENT + enum_value][POST]['x-speakeasy-entity-group'] = 'Connection'
    paths[CONNECTIONS_PATH_WITH_FRAGMENT + enum_value][POST]['operationId'] = 'create-' + enum_value + '-connection'

    paths[CONNECTIONS_ID_PATH_WITH_FRAGMENT + enum_value]['parameters'] = paths[CONNECTIONS_ID_PATH]['parameters']
    paths[CONNECTIONS_ID_PATH_WITH_FRAGMENT + enum_value][GET] = updateConnectionTypesRef(paths[CONNECTIONS_ID_PATH][GET], '#/components/schemas/connection_types', connection_type['$ref'])
    paths[CONNECTIONS_ID_PATH_WITH_FRAGMENT + enum_value][GET]['operationId'] = 'get-' + enum_value + '-connection'

    paths[CONNECTIONS_ID_PATH_WITH_FRAGMENT + enum_value][GET]['x-speakeasy-entity-operation'] = 'Connection' + enum_value + '#get'

    paths[CONNECTIONS_ID_PATH_WITH_FRAGMENT + enum_value][PATCH] = updateConnectionTypesRef(paths[CONNECTIONS_ID_PATH_WITH_FRAGMENT + enum_value][PATCH], '#/components/schemas/connection_types', connection_type['$ref'])

    # Copy DELETES over, and annotate
    paths[CONNECTIONS_ID_PATH_WITH_FRAGMENT + enum_value][DELETE] = updateConnectionTypesRef(paths[CONNECTIONS_ID_PATH][DELETE], '', '')
    paths[CONNECTIONS_ID_PATH_WITH_FRAGMENT + enum_value][DELETE]['operationId'] = 'delete-' + enum_value + '-connection'
    paths[CONNECTIONS_ID_PATH_WITH_FRAGMENT + enum_value][DELETE]['x-speakeasy-entity-operation'] = 'Connection' + enum_value + '#delete'

    schemas[connections_schema]['x-speakeasy-entity'] = 'Connection' + enum_value
    schemas[connections_schema]['x-speakeasy-param-suppress-computed-diff'] = 'true'

def update_ref_in_all_of(all_of, new_reference):
    for obj in all_of:
        if isinstance(obj, dict) and '$ref' in obj:
            obj['$ref'] = new_reference

def remove_properties_in_all_of(all_of, keep_properties):
    """
    Removes all properties not in the keep_properties array

    Args:
        all_of (list): The list of allOf schema blocks
        properties (list): List of property names to keep
    """
    for obj in all_of:
        if isinstance(obj, dict) and not '$ref' in obj and 'properties' in obj:
            new_properties = dict()
            for prop in obj['properties']:
                if prop in keep_properties:
                    new_properties[prop] = obj['properties'][prop]
            obj['properties'] = new_properties

def remove_required_properties(all_of, properties):
    """
    For each object block in the given allOf:
    - Removes specified properties from 'required' property
    - Deletes 'required' property if it becomes empty
    
    Args:
        all_of (list): The list of allOf schema blocks
        properties (list): List of property names to remove from 'required' and to update
    """
    for block in all_of:
        if isinstance(block, dict) and block.get("type") == "object":
            # Remove items from required list
            required = block.get("required", [])
            if isinstance(required, list):
                block["required"] = [r for r in required if r not in properties]
                if not block["required"]:
                    del block["required"]

# Create source_x_update schemas for each source type
schemas['source_types_update'] = copy.deepcopy(schemas['source_types'])
schemas['source_types_update']['title'] += ' Update'
source_types_update = schemas['source_types_update']

# Create source type enum to re-use
if 'source_type_enum' not in schemas:
    schemas['source_type_enum'] = {
        'type': 'string',
        'enum': []
    }

# set up each source_x_update schemas with the correct properties and type references
for source_type in source_types_update['oneOf']:
    source_type_schema = schema_from_ref(source_type['$ref'])
    enum_value = get_enum_value_from_connection_spec(schemas[source_type_schema])
    source_type['$ref'] += '_update'

    schemas[source_type_schema + '_update'] = copy.deepcopy(schemas[source_type_schema])
    schemas[source_type_schema + '_update']['title'] += ' Update'

    update_ref_in_all_of(schemas[source_type_schema + '_update']['allOf'], "#/components/schemas/source_update")
    remove_properties_in_all_of(schemas[source_type_schema + '_update']['allOf'], ['type'])
    remove_required_properties(schemas[source_type_schema + '_update']['allOf'], ['type'])

    schemas[source_type_schema + '_update']["x-empty-superclass"] = True

    # Add enum value if not already added
    if enum_value not in schemas['source_type_enum']['enum']:
        schemas['source_type_enum']['enum'].append(enum_value)


# Add _update to every reference in the discriminator mapping too for source_types_update
source_types_update['discriminator']['mapping'] = {
    k: v + '_update' for k, v in source_types_update['discriminator']['mapping'].items()
}

# Update references to use the newly generated schemas
schemas['pipeline_update']['properties']['source']['$ref'] = "#/components/schemas/source_types_update"
schemas['source_update']['properties']['type'] = {
    '$ref': '#/components/schemas/source_type_enum'
}
schemas['source']['properties']['type'] = {
    '$ref': '#/components/schemas/source_type_enum'
}
del schemas['source_update']['x-speakeasy-name-override']

# Fix discriminator mappings to point to the correct schema references
fix_discriminator_mappings(schemas, 'source')
fix_discriminator_mappings(schemas, 'connection')

with open(outputSchemaFile, 'w+') as f:
    j.dump(api_spec, f, indent=4)
