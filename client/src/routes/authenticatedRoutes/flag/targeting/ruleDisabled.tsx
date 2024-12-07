import { Badge } from "flowbite-react";
import { useFormContext } from "react-hook-form";

export const RuleDisabled = ({ fieldName }: { fieldName?: string }) => {
  const { watch } = useFormContext();
  return (
    <>
      <div className={"flex items-start"}>
        <Badge size="sm" className={"mr-3 mt-1"} color={"green"}>
          IF
        </Badge>
        <div className={"font-mono"}>{watch(`${fieldName}.query`)}</div>
      </div>
    </>
  );
};
