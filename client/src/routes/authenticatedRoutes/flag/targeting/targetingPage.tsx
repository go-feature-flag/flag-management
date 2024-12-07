import { Button, Timeline } from "flowbite-react";
import { useFormContext, useWatch } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { FaPlusCircle } from "react-icons/fa";
import { TbTargetArrow } from "react-icons/tb";
import { v4 as uuidv4 } from "uuid";
import { Section } from "../../../../components/section/section.tsx";
import type {
  FeatureFlagFormData,
  FormFlagPageTargetingRule,
} from "../../../../models/featureFlagFormData.ts";
import { Rule } from "./rule.tsx";

const translationBaseKey = "page.flag.targeting";

export const TargetingPage = () => {
  const { t } = useTranslation();
  const { setValue } = useFormContext<FeatureFlagFormData>();
  const { targetingRules } = useWatch<FeatureFlagFormData>();

  function addRule() {
    const t = (targetingRules as FormFlagPageTargetingRule[]) ?? [];

    // TODO: when building the name of the rule, we should check if the name is already in use.
    const name = `Rule ${(targetingRules?.length ?? 0) + 1}`;
    t.push({
      name,
      variation: {
        id: uuidv4(),
      },
      disabled: false,
      query: "",
    });
    setValue("targetingRules", t, {
      shouldValidate: true,
      shouldDirty: true,
    });
  }

  return (
    <Section
      outsideTitle={t(`${translationBaseKey}.title`)}
      info={t(`${translationBaseKey}.description`)}
    >
      <RuleTimeline />
      <AddRuleTopButton onClick={addRule} />
    </Section>
  );
};

const AddRuleTopButton = ({ onClick }: { onClick: () => void }) => {
  const { t } = useTranslation();
  return (
    <div className={"flex"}>
      <Button size={"xs"} onClick={onClick}>
        <FaPlusCircle className={"mr-2"} />{" "}
        {t(`${translationBaseKey}.addRuleTopButton.title`)}
      </Button>
      <div className={"flex items-center pl-5 text-sm italic text-gray-500"}>
        {t(`${translationBaseKey}.addRuleTopButton.notice`)}
      </div>
    </div>
  );
};

const RuleTimeline = () => {
  const { targetingRules } = useWatch<FeatureFlagFormData>();
  const { setValue } = useFormContext<FeatureFlagFormData>();

  const deleteRule = (ruleId?: string) => {
    const t = (targetingRules as FormFlagPageTargetingRule[]) ?? [];
    setValue(
      "targetingRules",
      t?.filter((rule) => rule.variation?.id !== ruleId),
    );
  };

  return (
    <Timeline className={"ml-1 mt-5"}>
      {targetingRules?.map((rule, index) => (
        <Timeline.Item key={rule.variation?.id}>
          <Timeline.Point icon={TbTargetArrow} />
          <Timeline.Content>
            <Rule
              rule={rule}
              handleDelete={deleteRule}
              fieldName={`targetingRules[${index}]`}
            />
          </Timeline.Content>
        </Timeline.Item>
      ))}
    </Timeline>
  );
};
