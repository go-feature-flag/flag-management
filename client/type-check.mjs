#!/usr/bin/env node
// this code comes from https://github.com/microsoft/TypeScript/issues/47387#issuecomment-1169286648
// and is used to ignore node_modules in tsconfig.json

import path from "path";
import ts from "typescript";

//
// parse tsconfig.json
//

const configFilePath = ts.findConfigFile(
  "./",
  ts.sys.fileExists,
  "tsconfig.json"
);
const projectPath = path.resolve(path.dirname(configFilePath));
const configFile = ts.readConfigFile(configFilePath, ts.sys.readFile);

const compilerOptions = ts.parseJsonConfigFileContent(
  configFile.config,
  ts.sys,
  "./",
  {
    incremental: true,
    tsBuildInfoFile: path.join(projectPath, "type-check.tsbuildinfo"),

    // required to force emission of tsbuildinfo file(?)
    outDir: "tmp",
    noEmit: true,
  }
);

//
// setup program
//

const program = ts.createIncrementalProgram({
  rootNames: compilerOptions.fileNames,
  options: compilerOptions.options,
});

//
// run diagnostics and emit a tsbuildinfo for incremental type-checks
//

const preEmitDiagnostics = ts.getPreEmitDiagnostics(program);
const emitResult = program.emit();
const buildInfoEmitResult = program.emitBuildInfo();

const diagnostics = [
  ...preEmitDiagnostics,
  ...emitResult.diagnostics,
  ...buildInfoEmitResult.diagnostics,
].filter(
  (it) =>
    it.file.fileName.startsWith(projectPath + path.sep) &&
    !it.file.fileName.includes("node_modules")
);

//
// setup reporter utilities
//

const errorCount = ts.getErrorCountForSummary(diagnostics);
const filesInError = ts.getFilesInErrorForSummary(diagnostics);
const diagnosticsReporter = ts.createDiagnosticReporter(ts.sys, true);

const reportDiagnostics = (diagnostics) => {
  for (const diagnostic of diagnostics) {
    diagnosticsReporter(diagnostic);
  }
};

const reportSummary = (errorCount, filesInError) => {
  console.log(
    ts.getErrorSummaryText(errorCount, filesInError, ts.sys.newLine, ts.sys)
  );
};

//
// flush output to console
//

reportDiagnostics(diagnostics);
reportSummary(errorCount, filesInError);

process.exit(errorCount === 0 ? 0 : 1);
