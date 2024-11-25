import { getSnippetReplace, Snippet } from "./snippet.ts";

export const DotnetSnippet = new (class extends Snippet {
  getSnippetInstall() {
    return `dotnet add package OpenFeature.Contrib.GOFeatureFlag`;
  }
  getSnippetInit() {
    return `using OpenFeature;
using OpenFeature.Contrib.GOFeatureFlag;

// ...

var goFeatureFlagProvider = new GoFeatureFlagProvider(new GoFeatureFlagProviderOptions
{
    Endpoint = "https://relay.proxy.gofeatureflag.org/",
    Timeout = new TimeSpan(300000 * TimeSpan.TicksPerMillisecond)
});
Api.Instance.SetProvider(goFeatureFlagProvider);
var client = Api.Instance.GetClient("my-app");`;
  }
  getSnippetEvaluate(flagName?: string, type?: string) {
    const { func, defaultValue } = getSnippetReplace(type, "dotnet");
    return `var userContext = EvaluationContext.Builder()
    .Set("targetingKey", "1d1b9238-2591-4a47-94cf-d2bc080892f1") // user unique identifier (mandatory)
    .Set("firstname", "john")
    .Set("lastname", "doe")
    .Set("email", "john.doe@gofeatureflag.org")
    .Build();

var myFlag = await client.${func}("${flagName}", ${defaultValue}, userContext);`;
  }
})();
