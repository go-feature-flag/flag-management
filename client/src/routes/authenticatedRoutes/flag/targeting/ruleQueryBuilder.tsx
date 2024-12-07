import { Badge, FloatingLabel, Select } from "flowbite-react";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { FaTrashAlt } from "react-icons/fa";
import type {
  ActionWithRulesProps,
  OperatorSelectorProps,
} from "react-querybuilder";
import { QueryBuilder } from "react-querybuilder";
import { config } from "../../../../config.ts";
import "./ruleQueryBuilder.css";

const translationBaseKey = "page.flag.targeting.queryBuilder";
export const RuleQueryBuilder = () => {
  const [query, setQuery] = useState({
    combinator: "or",
    rules: [{ field: "", operator: "eq", value: "" }],
  });
  const { t } = useTranslation();
  return (
    <div className={"flex items-start"}>
      <Badge size="sm" className={"mr-3 mt-1"} color={"green"}>
        IF
      </Badge>
      <div className={"w-full"}>
        <QueryBuilder
          getRuleClassname={() => "flex gap-4"}
          controlClassnames={{
            header: "flex w-full gap-4 mb-4 place-items-center",
            addGroup:
              "border-dotted border h-10 items-center gap-1 font-semibold bg-gray-100 text-gray-800 text-sm font-medium me-2 px-2.5 py-0.5 rounded dark:bg-gray-800 dark:text-gray-300 px-2 py-0.5",
            addRule:
              "border-dotted border h-10 items-center gap-1 font-semibold bg-green-100 text-goff-700 text-sm font-medium me-2 px-2.5 py-0.5 rounded dark:text-white-400 dark:bg-goff-300 px-2 py-0.5",
            ruleGroup:
              "border-gray-500 dark:border-green-300 border-dotted border rounded-lg p-3 mb-1 dark:bg-gray-800",
            combinators:
              "me-2 ml-1 w-fit place-items-center rounded border-dotted bg-blue-100 text-sm font-medium text-blue-800 dark:border-goff-300 dark:bg-gray-900 dark:text-goff-300",
          }}
          getDefaultOperator={() => "eq"}
          fields={[]}
          controlElements={{
            fieldSelector: FieldSelector,
            valueEditor: ValueSelector,
            operatorSelector: OperatorSelector,
            removeRuleAction: RemoveRuleAction,
            removeGroupAction: RemoveRuleAction,
          }}
          translations={{
            addRule: { label: t(`${translationBaseKey}.addRule`) },
            addGroup: { label: t(`${translationBaseKey}.addGroup`) },
          }}
          parseNumbers={true}
          resetOnFieldChange={false}
          resetOnOperatorChange={false}
          addRuleToNewGroups={true}
          operators={config.ruleOperators.map((operator) => {
            return { name: operator.name, label: t(operator.translationKey) };
          })}
          query={query}
          onQueryChange={(query) => setQuery(query)}
        />
      </div>
    </div>
  );
};

const ValueSelector = () => {
  return (
    <div className={"w-5/12"}>
      <FloatingLabel
        placeholder="Value"
        required
        sizing={"sm"}
        label={"Value"}
        variant={"outlined"}
      />
    </div>
  );
};

const FieldSelector = () => {
  return (
    <div className={"w-4/12"}>
      <FloatingLabel
        placeholder="Evaluation Context Field"
        required
        sizing={"sm"}
        label={"Evaluation Context Field"}
        variant={"outlined"}
      />
    </div>
  );
};

const OperatorSelector = ({ options }: OperatorSelectorProps) => {
  const opt = (options as { label: string; name: string }[]) ?? [];
  return (
    <div className={"w-2/12"}>
      <Select sizing={"md"}>
        {opt.map(({ name, label }) => (
          <option key={name} value={name}>
            {label}
          </option>
        ))}
      </Select>
    </div>
  );
};

const RemoveRuleAction = ({ handleOnClick }: ActionWithRulesProps) => {
  return (
    <div className={"flex w-1/12"}>
      <button
        type="button"
        className={"text-gray-800 dark:text-gray-50"}
        onClick={handleOnClick}
      >
        <FaTrashAlt className="h-4 w-4" />
      </button>
    </div>
  );
};
