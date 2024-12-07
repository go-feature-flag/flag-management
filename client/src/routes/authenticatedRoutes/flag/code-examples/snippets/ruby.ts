import { getSnippetReplace, Snippet } from "./snippet.ts";

export const RubySnippet = new (class extends Snippet {
  getSnippetInstall() {
    return `# Add this line to your application's Gemfile:
gem 'openfeature-go-feature-flag-provider'
# And then execute:
bundle install

## Or install it yourself as:
gem install openfeature-go-feature-flag-provider`;
  }
  getSnippetInit() {
    return `options = OpenFeature::GoFeatureFlag::Options.new(endpoint: "https://relay.proxy.gofeatureflag.org")
provider = OpenFeature::GoFeatureFlag::Provider.new(options:)

OpenFeature::SDK.configure do |config|
   config.set_provider(provider)
end
client = OpenFeature::SDK.build_client`;
  }
  getSnippetEvaluate(flagName?: string, type?: string) {
    const { func, defaultValue } = getSnippetReplace(type, "ruby");
    return `evaluation_context = OpenFeature::SDK::EvaluationContext.new(targeting_key: "9b9450f8-ab5c-4dcf-872f-feda3f6ccb16")
    
my_flag = client.${func}(
  flag_key: "${flagName}",
  default_value: ${defaultValue},
  evaluation_context:
)`;
  }
})();
