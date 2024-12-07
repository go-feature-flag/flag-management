import { getSnippetReplace, Snippet } from "./snippet.ts";

export const JavaSnippet = new (class extends Snippet {
  getSnippetInstall() {
    return `// Graddle
implementation group: 'dev.openfeature', name: 'javasdk', version: '1.+'
implementation group: 'dev.openfeature.contrib.providers', name: 'go-feature-flag', version: '0.+'`;
  }
  getSnippetInit() {
    return `import dev.openfeature.contrib.providers.gofeatureflag.GoFeatureFlagProvider;
import dev.openfeature.contrib.providers.gofeatureflag.GoFeatureFlagProviderOptions;
import dev.openfeature.sdk.EvaluationContext;
import dev.openfeature.sdk.MutableContext;
import dev.openfeature.sdk.OpenFeatureAPI;

// ...

GoFeatureFlagProviderOptions options = GoFeatureFlagProviderOptions.builder().endpoint("https://relay.proxy.gofeatureflag.org/").build();
GoFeatureFlagProvider provider = new GoFeatureFlagProvider(options);

OpenFeatureAPI.getInstance().setProvider(provider);
OpenFeatureAPI api = OpenFeatureAPI.getInstance();
Client featureFlagClient = api.getClient();`;
  }
  getSnippetEvaluate(flagName?: string, type?: string) {
    const { func, defaultValue, resultType } = getSnippetReplace(type, "java");
    return `EvaluationContext userContext = new MutableContext("1d1b9238-2591-4a47-94cf-d2bc080892f1")
  .add("firstname", "john")
  .add("lastname", "doe")
  .add("email","john.doe@gofeatureflag.org");

${resultType} myFlag = featureFlagClient.${func}("${flagName}", ${defaultValue}, userContext);`;
  }
})();
