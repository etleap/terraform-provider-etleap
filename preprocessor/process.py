import json as j
import sys

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

if len(sys.argv) < 3:
    print("Usage python process.py input-schema.json output-schema.json")
    sys.exit(1)

inputSchemaFile = sys.argv[1]
outputSchemaFile = sys.argv[2]

with open(inputSchemaFile, 'r+') as f:
    api_spec = j.load(f)

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

with open(outputSchemaFile, 'w+') as f:
    j.dump(api_spec, f, indent=4)
