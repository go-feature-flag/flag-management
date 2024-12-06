import { OFREPWebProvider } from "@openfeature/ofrep-web-provider";
import { OpenFeature } from "@openfeature/react-sdk";
import { v4 as uuidv4 } from "uuid";

export const initGoFeatureFlag = async () => {
  // get GOFF targetingKey from localStorage if exist
  let targetingKey = localStorage.getItem("goff_targetingKey");
  if (!targetingKey) {
    targetingKey = uuidv4();
    localStorage.setItem("goff_targetingKey", targetingKey);
  }
  const goFeatureFlagWebProvider = new OFREPWebProvider({
    baseUrl: "https://relay.proxy.gofeatureflag.org",
    pollInterval: 60000,
  });

  await OpenFeature.setContext({
    targetingKey,
    source: "go-feature-flag app",
  });
  await OpenFeature.setProviderAndWait(goFeatureFlagWebProvider);
};
