import { getSnippetReplace, Snippet } from "./snippet.ts";

export const PythonSnippet = new (class extends Snippet {
  getSnippetInstall() {
    return `pip install gofeatureflag-python-provider`;
  }
  getSnippetInit() {
    return `from gofeatureflag_python_provider.provider import GoFeatureFlagProvider
from gofeatureflag_python_provider.options import GoFeatureFlagOptions
from openfeature import api
from openfeature.evaluation_context import EvaluationContext

// ...

goff_provider = GoFeatureFlagProvider(
    options=GoFeatureFlagOptions(endpoint="https://gofeatureflag.org/")
)
api.set_provider(goff_provider)
client = api.get_client(name="test-client")`;
  }
  getSnippetEvaluate(flagName?: string, type?: string) {
    const { func, defaultValue } = getSnippetReplace(type, "python");
    return `evaluation_ctx = EvaluationContext(
    targeting_key="d45e303a-38c2-11ed-a261-0242ac120002",
    attributes={
        "email": "john.doe@gofeatureflag.org",
        "firstname": "john",
        "lastname": "doe",
    },
)

my_flag = client.${func}(
          flag_key="${flagName}",
          default_value=${defaultValue},
          evaluation_context=ctx,
      )`;
  }
})();
