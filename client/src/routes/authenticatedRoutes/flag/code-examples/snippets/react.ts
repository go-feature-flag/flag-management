import { getSnippetReplace, Snippet } from "./snippet.ts";

export const ReactSnippet = new (class extends Snippet {
  getSnippetInstall() {
    return `npm install @openfeature/go-feature-flag-web-provider @openfeature/web-sdk @openfeature/react-sdk @openfeature/core`;
  }

  getSnippetInit() {
    return `import { EvaluationContext, OpenFeature, OpenFeatureProvider, useFlag } from "@openfeature/react-sdk";
import { GoFeatureFlagWebProvider } from "@openfeature/go-feature-flag-web-provider";

const goFeatureFlagWebProvider = new GoFeatureFlagWebProvider({
  endpoint: "http://localhost:1031"
});

// Set the initial context for your evaluations
OpenFeature.setContext({
  targetingKey: "user-1",
  admin: false
});

// Instantiate and set our provider (be sure this only happens once)!
// Note: there's no need to await its initialization, the React SDK handles re-rendering and suspense for you!
OpenFeature.setProvider(goFeatureFlagWebProvider);

// Enclose your content in the configured provider
function App() {
  return (
    <OpenFeatureProvider>
      <Page />
    </OpenFeatureProvider>
  );
}`;
  }

  getSnippetEvaluate(flagName?: string, type?: string) {
    const { defaultValue } = getSnippetReplace(type, "node");
    return `function Page() {
  // Use the "query-style" flag evaluation hook, specifying a flag-key and a default value.
  const { value: myFlag } = useFlag('${flagName}', ${defaultValue});
  return (
    <div className="App">
      <header className="App-header">
        {myFlag}
      </header>
    </div>
  );
}`;
  }
})();
