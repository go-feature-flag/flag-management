import { Select, useThemeMode } from "flowbite-react";
import { useState } from "react";
import { useWatch } from "react-hook-form";
import { PrismAsyncLight as SyntaxHighlighter } from "react-syntax-highlighter";
import {
  oneDark,
  oneLight,
} from "react-syntax-highlighter/dist/esm/styles/prism";
import { CopyToClipboard } from "../../../../components/copyToClipboard/copyToClipboard.tsx";
import { Section } from "../../../../components/section/section.tsx";
import type { FeatureFlagFormData } from "../../../../models/featureFlagFormData.ts";
import { DotnetSnippet } from "./snippets/dotnet.ts";
import { GOSnippet } from "./snippets/go.ts";
import { JavaSnippet } from "./snippets/java.ts";
import { KotlinSnippet } from "./snippets/kotlin.ts";
import { NodeSnippet } from "./snippets/node.ts";
import { PythonSnippet } from "./snippets/python.ts";
import { ReactSnippet } from "./snippets/react.ts";
import { RubySnippet } from "./snippets/ruby.ts";
import type { Snippet } from "./snippets/snippet.ts";
import { SwiftSnippet } from "./snippets/swift.ts";
import { WebSnippet } from "./snippets/web.ts";

interface CodeExamplesProps {
  name: string;
  language: string;
  type: "server" | "client";
  snippets: Snippet;
}

const examples: CodeExamplesProps[] = [
  {
    name: "Web",
    language: "typescript",
    type: "client",
    snippets: WebSnippet,
  },
  {
    name: "React",
    language: "typescript",
    type: "client",
    snippets: ReactSnippet,
  },
  {
    name: "Swift (iOS / macOS)",
    language: "swift",
    type: "client",
    snippets: SwiftSnippet,
  },
  {
    name: "Kotlin (Android)",
    language: "kotlin",
    type: "client",
    snippets: KotlinSnippet,
  },
  {
    name: "Java",
    language: "java",
    type: "server",
    snippets: JavaSnippet,
  },
  {
    name: "GO",
    language: "go",
    type: "server",
    snippets: GOSnippet,
  },
  {
    name: "NodeJS",
    language: "typescript",
    type: "server",
    snippets: NodeSnippet,
  },
  {
    name: "Python",
    language: "python",
    type: "server",
    snippets: PythonSnippet,
  },
  {
    name: "Ruby",
    language: "ruby",
    type: "server",
    snippets: RubySnippet,
  },
  {
    name: ".NET",
    language: "dotnet",
    type: "server",
    snippets: DotnetSnippet,
  },
];

export const CodeExamples = () => {
  const { mode } = useThemeMode();
  const syntaxHighlighterTheme = mode === "dark" ? oneDark : oneLight;
  const [language, setLanguage] = useState("go");
  const currentLanguage =
    examples.find((example) => example.name === language) ?? examples[0];
  const { name, type } = useWatch<FeatureFlagFormData>();
  return (
    <>
      <h1 className={"my-2"}>Code Examples</h1>
      <Section
        innerTitle={"Installing the SDK"}
        info={
          "You first need to put as a dependency of your project the OpenFeature SDK and the GO Feature Flag Provider"
        }
      >
        <ChooseLanguage language={language} setLanguage={setLanguage} />
        <div className="relative w-full">
          <div className={"absolute right-2 top-2 "}>
            <CopyToClipboard
              size={"sm"}
              textToCopy={currentLanguage?.snippets.getSnippetInstall()}
            />
          </div>
          <SyntaxHighlighter
            language={currentLanguage.language}
            style={syntaxHighlighterTheme}
            wrapLongLines={true}
          >
            {currentLanguage?.snippets.getSnippetInstall() ?? ""}
          </SyntaxHighlighter>
        </div>
      </Section>
      <Section
        innerTitle={"Initialize your project"}
        info={
          "Initialize your OpenFeature SDK and use the GO Feature Flag provider"
        }
      >
        <ChooseLanguage language={language} setLanguage={setLanguage} />
        <div className="relative w-full">
          <div className={"absolute right-2 top-2 "}>
            <CopyToClipboard
              size={"sm"}
              textToCopy={currentLanguage?.snippets.getSnippetInit()}
            />
          </div>
          <SyntaxHighlighter
            language={currentLanguage.language}
            style={syntaxHighlighterTheme}
            wrapLongLines={true}
          >
            {currentLanguage?.snippets.getSnippetInit() ?? ""}
          </SyntaxHighlighter>
        </div>
      </Section>
      <Section
        innerTitle={"Evaluate your flag"}
        info={
          "It's time to use your flag, wrap your new feature with the feature flag"
        }
      >
        <ChooseLanguage language={language} setLanguage={setLanguage} />
        <div className="relative w-full">
          <div className={"absolute right-2 top-2 "}>
            <CopyToClipboard
              size={"sm"}
              textToCopy={currentLanguage?.snippets.getSnippetEvaluate()}
            />
          </div>
          <SyntaxHighlighter
            language={currentLanguage.language}
            style={syntaxHighlighterTheme}
            wrapLongLines={true}
          >
            {currentLanguage?.snippets.getSnippetEvaluate(name, type)}
          </SyntaxHighlighter>
        </div>
      </Section>
    </>
  );
};

const ChooseLanguage = ({
  language,
  setLanguage,
}: {
  language: string;
  setLanguage: (language: string) => void;
}) => {
  return (
    <Select
      value={language}
      onChange={(e) => {
        setLanguage(e.target.value);
      }}
    >
      <option disabled>Client</option>
      {examples
        .filter((example) => example.type === "client")
        .map((example) => (
          <option key={example.name}>{example.name}</option>
        ))}
      <option disabled>Server</option>
      {examples
        .filter((example) => example.type === "server")
        .map((example) => (
          <option key={example.name}>{example.name}</option>
        ))}
    </Select>
  );
};
