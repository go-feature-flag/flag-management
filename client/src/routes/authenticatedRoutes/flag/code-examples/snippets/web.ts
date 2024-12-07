import { getSnippetReplace, Snippet } from "./snippet.ts";

export const WebSnippet = new (class extends Snippet {
  getSnippetInstall() {
    return `npm install @openfeature/go-feature-flag-web-provider @openfeature/web-sdk`;
  }
  getSnippetInit() {
    return `const evaluationCtx: EvaluationContext = {
  targetingKey: 'user-key',
  email: 'john.doe@gofeatureflag.org',
  name: 'John Doe',
};

const goFeatureFlagWebProvider = new GoFeatureFlagWebProvider({
  endpoint: "https://relay.proxy.gofeatureflag.org",
}, logger);

await OpenFeature.setContext(evaluationCtx); // Set the static context for OpenFeature
OpenFeature.setProvider(goFeatureFlagWebProvider); // Attach the provider to OpenFeature
const client = await OpenFeature.getClient();

// You can add handlers to know what happen in the provider
client.addHandler(ProviderEvents.Ready, () => { ... });
client.addHandler(ProviderEvents.Error, () => { //... });
client.addHandler(ProviderEvents.Stale, () => { //... });
client.addHandler(ProviderEvents.ConfigurationChanged, () => { //... });
`;
  }
  getSnippetEvaluate(flagName?: string, type?: string) {
    const { func, defaultValue } = getSnippetReplace(type, "node");
    return `const value = client.${func}('${flagName}', ${defaultValue});`;
  }
})();
