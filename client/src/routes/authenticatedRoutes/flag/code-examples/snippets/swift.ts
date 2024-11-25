import { getSnippetReplace, Snippet } from "./snippet.ts";

export const SwiftSnippet = new (class extends Snippet {
  getSnippetInstall() {
    return `// Swift Package Manager: In the dependencies section of Package.swift add
.package(url: "https://github.com/go-feature-flag/openfeature-swift-provider.git", from: "0.1.0")

// and in the target dependencies section add
.product(name: "GOFeatureFlag", package: "openfeature-swift-provider"),`;
  }
  getSnippetInit() {
    return `import GOFeatureFlag
import OpenFeature

// ...

let options = GoFeatureFlagProviderOptions(endpoint: "https://relay.proxy.gofeatureflag.org")
let provider = GoFeatureFlagProvider(options: options)

let evaluationContext = MutableContext(
  targetingKey: "myTargetingKey", 
  structure: MutableStructure()
)
OpenFeatureAPI.shared.setProvider(provider: provider, initialContext: evaluationContext)`;
  }
  getSnippetEvaluate(flagName?: string, type?: string) {
    const { func, defaultValue } = getSnippetReplace(type, "swift");
    return `let client = OpenFeatureAPI.shared.getClient()
let result = client.${func}(key: "${flagName}", defaultValue: ${defaultValue})`;
  }
})();
