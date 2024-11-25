import { getSnippetReplace, Snippet } from "./snippet.ts";

export const KotlinSnippet = new (class extends Snippet {
  getSnippetInstall() {
    return `// Graddle
api("dev.openfeature:android-sdk:0.3.0")
api("org.gofeatureflag.openfeature:gofeatureflag-kotlin-provider0.1.0")`;
  }
  getSnippetInit() {
    return `import org.gofeatureflag.openfeature.bean.GoFeatureFlagOptions
import org.gofeatureflag.openfeature.GoFeatureFlagProvider
// ...

val evaluationContext: EvaluationContext = ImmutableContext(
        targetingKey = "37b7a5b9-6bc4-4c22-9f1c-a24e96d34137",
        attributes = mapOf(
          "age" to Value.Integer(22), 
          "email" to Value.String("contact@gofeatureflag.org")
        )
    )

OpenFeatureAPI.setProvider(
    GoFeatureFlagProvider(
        options = GoFeatureFlagOptions( endpoint = "https://relay.proxy.gofeatureflag.org")
    ),
    evaluationContext
)
`;
  }
  getSnippetEvaluate(flagName?: string, type?: string) {
    const { func, defaultValue } = getSnippetReplace(type, "kt-android");
    return `val client = OpenFeatureAPI.getClient()
val result = client.${func}("${flagName}", ${defaultValue})`;
  }
})();
