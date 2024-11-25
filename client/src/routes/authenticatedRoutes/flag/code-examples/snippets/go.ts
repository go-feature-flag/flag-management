import { getSnippetReplace, Snippet } from "./snippet.ts";

export const GOSnippet = new (class extends Snippet {
  getSnippetInstall() {
    return `go get github.com/open-feature/go-sdk
go get github.com/open-feature/go-sdk-contrib/providers/go-feature-flag`;
  }
  getSnippetInit() {
    return `import (
  // ...
  gofeatureflag "github.com/open-feature/go-sdk-contrib/providers/go-feature-flag/pkg"
  of "github.com/open-feature/go-sdk/pkg/openfeature"
)

// ...

options := gofeatureflag.ProviderOptions{
    Endpoint: "https://relay.proxy.gofeatureflag.org",
}
provider, err := gofeatureflag.NewProvider(options)
of.SetProvider(provider)
client := of.NewClient("my-app")`;
  }
  getSnippetEvaluate(flagName?: string, type?: string) {
    const { func, defaultValue } = getSnippetReplace(type, "go");
    return `evaluationCtx := of.NewEvaluationContext(
    "1d1b9238-2591-4a47-94cf-d2bc080892f1",
    map[string]interface{}{
      "firstname", "john",
      "lastname", "doe",
      "email", "john.doe@gofeatureflag.org",
    })
myFlag, _ := client.${func}(context.TODO(), "${flagName}", ${defaultValue}, evaluationCtx)
`;
  }
})();
