import { getSnippetReplace, Snippet } from "./snippet.ts";

export const NodeSnippet = new (class extends Snippet {
  getSnippetInstall() {
    return `npm i @openfeature/server-sdk @openfeature/go-feature-flag-provider`;
  }
  getSnippetInit() {
    return `import {EvaluationContext, OpenFeature} from "@openfeature/server-sdk";
import {GoFeatureFlagProvider} from  "@openfeature/go-feature-flag-provider";


// init Open Feature SDK with GO Feature Flag provider
const goFeatureFlagProvider: GoFeatureFlagProvider = new GoFeatureFlagProvider({
endpoint: 'https://relay.proxy.gofeatureflag.org',
});
OpenFeature.setProvider(goFeatureFlagProvider);
const featureFlagClient = OpenFeature.getClient('my-app');`;
  }
  getSnippetEvaluate(flagName?: string, type?: string) {
    const { func, defaultValue } = getSnippetReplace(type, "node");
    return `const userContext: EvaluationContext = {
  targetingKey: '1d1b9238-2591-4a47-94cf-d2bc080892f1', // user unique identifier
  firstname: 'john',
  lastname: 'doe',
  email: 'john.doe@gofeatureflag.org',
  // ...
};

const myFlag = await featureFlagClient.${func}('${flagName}', ${defaultValue}, userContext);`;
  }
})();
