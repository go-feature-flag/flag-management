export abstract class Snippet {
  abstract getSnippetInstall(): string;

  abstract getSnippetInit(): string;

  abstract getSnippetEvaluate(flagName?: string, type?: string): string;
}

export type ExampleProps = Record<
  string,
  Record<string, { func: string; defaultValue: string; resultType?: string }>
>;

export const snippetReplacement: ExampleProps = {
  python: {
    boolean: {
      func: "get_boolean_value",
      defaultValue: "false",
    },
    string: {
      func: "get_string_value",
      defaultValue: '"default"',
    },
    integer: {
      func: "get_integer_value",
      defaultValue: "0",
    },
    double: {
      func: "get_float_value",
      defaultValue: "0.9",
    },
    json: {
      func: "get_object_value",
      defaultValue: "my_object_instance",
    },
  },
  go: {
    boolean: {
      func: "BooleanValue",
      defaultValue: "false",
    },
    string: {
      func: "StringValue",
      defaultValue: '"default"',
    },
    integer: {
      func: "IntValue",
      defaultValue: "0",
    },
    double: {
      func: "FloatValue",
      defaultValue: "0.9",
    },
    json: {
      func: "ObjectValue",
      defaultValue: "my_object_instance",
    },
  },
  node: {
    boolean: {
      func: "getBooleanValue",
      defaultValue: "false",
    },
    string: {
      func: "getStringValue",
      defaultValue: "'default'",
    },
    integer: {
      func: "getNumberValue",
      defaultValue: "0",
    },
    double: {
      func: "getNumberValue",
      defaultValue: "0.9",
    },
    json: {
      func: "getObjectValue<MyObject>",
      defaultValue: "my_object_instance",
    },
  },
  ruby: {
    boolean: {
      func: "fetch_boolean_value",
      defaultValue: "false",
    },
    string: {
      func: "fetch_string_value",
      defaultValue: '"default"',
    },
    integer: {
      func: "fetch_number_value",
      defaultValue: "0",
    },
    double: {
      func: "fetch_number_value",
      defaultValue: "0.9",
    },
    json: {
      func: "fetch_object_value",
      defaultValue: '{"default" => true}',
    },
  },
  dotnet: {
    boolean: {
      func: "GetBooleanValueAsync",
      defaultValue: "false",
    },
    string: {
      func: "GetStringValueAsync",
      defaultValue: '"default"',
    },
    integer: {
      func: "GetIntegerValueAsync",
      defaultValue: "0",
    },
    double: {
      func: "GetDoubleValueAsync",
      defaultValue: "0.9",
    },
    json: {
      func: "GetObjectValueAsync",
      defaultValue: "my_object_instance",
    },
  },
  java: {
    boolean: {
      func: "getBooleanValue",
      defaultValue: "false",
      resultType: "Boolean",
    },
    string: {
      func: "getStringValue",
      defaultValue: '"default"',
      resultType: "String",
    },
    integer: {
      func: "getIntegerValue",
      defaultValue: "0",
      resultType: "Integer",
    },
    double: {
      func: "getDoubleValue",
      defaultValue: "0.9",
      resultType: "Double",
    },
    json: {
      func: "getObjectValue",
      defaultValue: "my_object_instance",
      resultType: "Structure",
    },
  },
  swift: {
    boolean: {
      func: "getBooleanValue",
      defaultValue: "false",
    },
    string: {
      func: "getStringValue",
      defaultValue: '"default"',
    },
    integer: {
      func: "getIntegerValue",
      defaultValue: "0",
    },
    double: {
      func: "getDoubleValue",
      defaultValue: "0.9",
    },
    json: {
      func: "getObjectValue",
      defaultValue: 'Value.structure(["key":Value.integer("1234")]',
    },
  },
  "kt-android": {
    boolean: {
      func: "getBooleanValue",
      defaultValue: "false",
    },
    string: {
      func: "getStringValue",
      defaultValue: '"default"',
    },
    integer: {
      func: "getIntegerValue",
      defaultValue: "0",
    },
    double: {
      func: "getDoubleValue",
      defaultValue: "0.9",
    },
    json: {
      func: "getObjectValue",
      defaultValue:
        'Value.structure(mapOf("email" to Value.String("contact@gofeatureflag.org")))',
    },
  },
};

export const getSnippetReplace = (type = "boolean", language = "node") => {
  return (
    snippetReplacement[language][type] ?? {
      func: "<error>",
      defaultValue: "<error>",
    }
  );
};
